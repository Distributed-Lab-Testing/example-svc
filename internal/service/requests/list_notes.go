package requests

import "net/http"

func NewListNotesRequest(r *http.Request) (*GetNotesRequest, error) {
	return &GetNotesRequest{}, nil
}
