package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Distributed-Lab-Testing/example-svc/internal/service/ctx"
	"github.com/go-chi/chi"
)

type NoteUpdateRequest struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	noteIDParam := chi.URLParam(r, "id")
	noteID, err := strconv.ParseInt(noteIDParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	var updateReq NoteUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = ctx.DB(r).Notes().UpdateContent(noteID, updateReq.Content)
	if err != nil {
		http.Error(w, "Failed to update note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
