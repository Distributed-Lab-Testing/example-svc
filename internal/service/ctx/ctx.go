package ctx

import (
	"context"
	"errors"
	"net/http"

	"github.com/Distributed-Lab-Testing/example-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

type Note struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

const (
	logCtxKey ctxKey = iota
	logsCtxKey
	dbCtxKey
	notesCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func NotesQ(r *http.Request) (data.NotesQ, error) {
	v := r.Context().Value(notesCtxKey)
	if v == nil {
		return nil, errors.New("NotesQ not found in request context")
	}

	notesQ, ok := v.(data.NotesQ)
	if !ok {
		return nil, errors.New("context value is not of type data.NotesQ")
	}

	return notesQ, nil
}

func CtxDB(entry data.DB) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dbCtxKey, entry)
	}
}

func DB(r *http.Request) data.DB {
	return r.Context().Value(dbCtxKey).(data.DB)
}
