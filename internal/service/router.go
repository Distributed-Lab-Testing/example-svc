package service

import (
	"github.com/Distributed-Lab-Testing/example-svc/internal/service/ctx"
	"github.com/Distributed-Lab-Testing/example-svc/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			ctx.CtxLog(s.log),
			ctx.CtxDB(s.db),
		),
	)
	r.Route("/integrations/example-svc", func(r chi.Router) {
		r.Route("/notes", func(r chi.Router) {
			r.Post("/", handlers.CreateNoteHandler) // CRUD Create
			r.Get("/", handlers.ListNotesHandler)   // CRUD Get as a list
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetNoteHandler)       // CRUD GET
				r.Patch("/", handlers.UpdateNoteHandler)  // CRUD Update
				r.Delete("/", handlers.DeleteNoteHandler) // CRUD Delete
			})
		})
	})

	return r
}
