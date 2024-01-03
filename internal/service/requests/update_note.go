package requests

import (
	"encoding/json"
	"net/http"

	"github.com/Distributed-Lab-Testing/example-svc/resources"
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UpdateNoteRequest struct {
	ID   string         //`json:"-"`
	Data resources.Note `json:"data"`
}

// func NewUpdateNote(r *http.Request) (UpdateNoteRequest, error) {
// 	noteIDParam := chi.URLParam(r, "id")
// 	noteID, err := strconv.ParseInt(noteIDParam, 10, 64)

// 	if err != nil {
// 		return UpdateNoteRequest{}, err
// 	}

// 	request := UpdateNoteRequest{
// 		ID: noteID,
// 	}

// 	if err := json.NewDecoder(r.Body).Decode(&request.Data); err != nil {
// 		return request, err
// 	}

// 	return request, request.validate()
// }

// func (r UpdateNoteRequest) validate() error {
// 	return validation.ValidateStruct(&r,
// 		validation.Field(&r.ID, validation.Required),
// 		validation.Field(&r.Data.Attributes.Content, validation.Required),
// 	)
// }

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
	return validation.Errors{
		"user_id": validation.Validate(
			&r.ID,
			validation.Required,
			// is.UUID,
		),
	}.Filter()
}
