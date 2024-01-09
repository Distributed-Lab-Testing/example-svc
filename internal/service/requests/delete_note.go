package requests

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type DeleteNoteRequest struct {
	ID int64 `json:"id"`
}

func NewDeleteNoteRequest(r *http.Request) (*DeleteNoteRequest, error) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return nil, errors.New("invalid note ID")
	}

	return &DeleteNoteRequest{
		ID: id,
	}, nil
}
