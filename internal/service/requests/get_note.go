package requests

import (
	"net/http"
)

type GetNotesRequest struct {
}

func NewGetNotesRequest(r *http.Request) (*GetNotesRequest, error) {
	return &GetNotesRequest{}, nil
}
