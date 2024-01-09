package requests

import (
	"encoding/json"
	"net/http"

	"github.com/Distributed-Lab-Testing/example-svc/resources"
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UpdateNoteRequest struct {
	ID   string
	Data resources.Note `json:"data"`
}

func NewUpdateNote(r *http.Request) (UpdateNoteRequest, error) {
	request := UpdateNoteRequest{
		ID: chi.URLParam(r, "id"),
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, validation.Errors{"request": err}
	}

	return request, nil
}

func (r UpdateNoteRequest) validate() error {
	return validation.Errors{}.Filter()
}
