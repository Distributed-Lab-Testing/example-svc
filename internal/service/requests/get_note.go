package requests

import (
	"net/http"
)

// type GetNoteStatus struct {
// 	ID int64 `url:"-"`
// }

// func NewGetNoteStatusRequest(r *http.Request) (*GetNoteStatus, error) {
// 	idParam := chi.URLParam(r, "id")

// 	id, err := strconv.ParseInt(idParam, 10, 64)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &GetNoteStatus{
// 		ID: id,
// 	}, nil
// }

type GetNotesRequest struct {
	// Здесь могут быть другие параметры, если они нужны
}

func NewGetNotesRequest(r *http.Request) (*GetNotesRequest, error) {
	// Создаем запрос без ожидания ID
	return &GetNotesRequest{
		// Здесь могут быть другие параметры, если они нужны
	}, nil
}
