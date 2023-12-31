package handlers

import (
	"net/http"

	"github.com/Distributed-Lab-Testing/example-svc/internal/service/ctx"
	"github.com/Distributed-Lab-Testing/example-svc/internal/service/models"
	"github.com/Distributed-Lab-Testing/example-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetNotesRequest(r)
	if err != nil || request == nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	entry, err := ctx.DB(r).Notes().Get(request.ID)
	if err != nil {
		ctx.Log(r).WithError(err).Error("failed to get note entry from the database")
		ape.Render(w, problems.InternalError())
		return
	}
	if entry == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, models.NoteAsResource(*entry))
}
