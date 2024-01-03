package handlers

import (
	"net/http"

	"github.com/Distributed-Lab-Testing/example-svc/internal/service/ctx"
	"github.com/Distributed-Lab-Testing/example-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteNoteRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = ctx.DB(r).Notes().Delete(request.ID)
	if err != nil {
		ctx.Log(r).WithError(err).WithField("id", request.ID).Error("failed to delete note")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, http.StatusNoContent)
}
