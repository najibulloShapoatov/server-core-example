package service

import (
	"net/http"

	"github.com/najibulloShapoatov/server-core/server"
)

type RequestBody struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Response struct {
	ID          int          `json:"id"`
	RequestBody *RequestBody `json:"requestBody"`
}

//GET /request
func (s *Service) GetRequest(ctx *server.Context) (interface{}, int, error) {

	return map[string]interface{}{"status": "OK"}, http.StatusOK, nil
}

//GET /request1/{path1}
func (s *Service) GetRequest1(ctx *server.Context, path1 string) (interface{}, int, error) {

	return map[string]interface{}{"status": "OK", "path1": path1}, http.StatusOK, nil
}

//POST /request  body: RequestBody
func (s *Service) CreateRequest(ctx *server.Context, rb *RequestBody) (*Response, int, error) {

	if rb.ID < 0 {
		status, err := ctx.ErrorBadRequest("id is negative")
		return nil, status, err
	}

	resp := &Response{
		ID:          0,
		RequestBody: rb,
	}
	return resp, http.StatusOK, nil
}
