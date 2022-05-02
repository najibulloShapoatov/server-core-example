package auth

import (
	"time"

	"github.com/najibulloShapoatov/server-core/cache/redis"
	"github.com/najibulloShapoatov/server-core/server"
	"github.com/najibulloShapoatov/server-core/server/session"
)

type UserInfo struct {
	Token  string `json:"token"`
	Email  string `json:"email"`
	UserID string `json:"userId"`
	Role   string `json:"roles"`
}

func Middleware(next server.HandlerFunc) server.HandlerFunc {
	return func(ctx *server.Context) error {
		// add session to the context
		if ctx.Session == nil {
			var sessionID string
			// search for a session id on the session cookie
			cookie, _ := ctx.Request.Cookie(ctx.Server.Config.Session.CookieName)
			// if session cookie is not present try on the session header
			if cookie == nil {
				sessionID = ctx.Request.Header.Get("X-Session-Id")
			} else {
				sessionID = cookie.Value
			}
			var details = &UserInfo{}
			var rds = redis.New(nil)
			if rds != nil {
				if err := rds.Get(sessionID, &details); err == nil {
					ctx.Session = &session.Session{
						IP:      ctx.RemoteAddr(),
						Created: time.Now(),
						Name:    &details.Email,
					}
					ctx.Set("accountDetails", details)
				}
			}
		}
		return next(ctx)
	}
}
