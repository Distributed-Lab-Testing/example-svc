package handlers

import (
	"net/http"

	"github.com/Distributed-Lab-Testing/example-svc/internal/service/ctx"
	"github.com/Distributed-Lab-Testing/example-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetNotesRequest(r)
	if err != nil || request == nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = ctx.NotesQ(r).Delete()
	if err != nil {
		ctx.Log(r).WithError(err).WithField("note_id", nil).Error("failed to delete note")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, http.StatusNoContent)
}
