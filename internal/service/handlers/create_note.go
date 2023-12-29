package handlers

import (
	"net/http"

	"github.com/Distributed-Lab-Testing/example-svc/internal/data"
	"github.com/Distributed-Lab-Testing/example-svc/internal/service/ctx"
	"github.com/Distributed-Lab-Testing/example-svc/internal/service/requests"
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

	_, err = ctx.DB(r).Notes().Insert(data.Note{
		Content:   request.Data.Attributes.Content,
		CreatedAt: request.Data.Attributes.CreatedAt,
	})
	if err != nil {
		ctx.Log(r).WithError(err).Error("failed to insert new log")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
