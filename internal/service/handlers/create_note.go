package handlers

import (
	"net/http"
	"time"

	"github.com/Distributed-Lab-Testing/example-svc/internal/data"
	"github.com/Distributed-Lab-Testing/example-svc/internal/service/ctx"
	"github.com/Distributed-Lab-Testing/example-svc/internal/service/requests"
	"github.com/Distributed-Lab-Testing/example-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateNoteRequest(r)
	if err != nil || request == nil {
		ctx.Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	ids, err := ctx.DB(r).Notes().Insert(data.Note{
		Content:   request.Data.Attributes.Content,
		CreatedAt: time.Now(),
	})
	if err != nil {
		ctx.Log(r).WithError(err).Error("failed to insert new note")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if len(ids) == 0 {
		ctx.Log(r).Error("no ID returned after note insertion")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	ape.Render(w, resources.NewIdResponse(resources.NewKeyInt64(ids[0], resources.NOTES)))
}
