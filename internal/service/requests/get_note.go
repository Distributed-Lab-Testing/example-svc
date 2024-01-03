package requests

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type GetNotesRequest struct {
	ID int64
}

func NewGetNotesRequest(r *http.Request) (*GetNotesRequest, error) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return nil, errors.New("no note ID provided")
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, errors.New("invalid note ID")
	}
	return &GetNotesRequest{ID: id}, nil
}
