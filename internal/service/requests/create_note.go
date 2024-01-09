package requests

import (
	"encoding/json"
	"net/http"

	"github.com/Distributed-Lab-Testing/example-svc/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type CreateNoteRequest resources.CreateNoteRequest

func NewCreateNoteRequest(r *http.Request) (*CreateNoteRequest, error) {
	var request CreateNoteRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, validation.Errors{
			"body": errors.Wrap(err, "failed to decode"),
		}
	}

	return &request, request.validate(r)
}

func (r *CreateNoteRequest) validate(httpReq *http.Request) error {
	return validation.Errors{
		"content": validation.Validate(
			&r.Data.Attributes.Content,
			validation.Required,
		),
		"created_at": validation.Validate(
			&r.Data.Attributes.CreatedAt,
			validation.Required,
		),
	}.Filter()
}
