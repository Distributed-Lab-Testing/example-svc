package handlers

import (
	"net/http"
	"strconv"

	"github.com/Distributed-Lab-Testing/example-svc/internal/service/ctx"
	"github.com/Distributed-Lab-Testing/example-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	updateReq, err := requests.NewUpdateNote(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	noteID, err := strconv.ParseInt(updateReq.ID, 10, 64)
	if err != nil {
		ctx.Log(r).WithError(err).WithField("note_id", updateReq.ID).Error("invalid note ID")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	err = ctx.DB(r).Notes().UpdateContent(noteID, updateReq.Data.Attributes.Content)
	if err != nil {
		ctx.Log(r).WithError(err).WithField("note_id", noteID).Error("failed to update note")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
