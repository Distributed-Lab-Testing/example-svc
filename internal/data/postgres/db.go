package postgres

import (
	"github.com/Distributed-Lab-Testing/example-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type db struct {
	raw *pgdb.DB
}

func NewDB(rawDB *pgdb.DB) data.DB {
	return &db{
		raw: rawDB,
	}
}

func (db *db) New() data.DB {
	return NewDB(db.raw.Clone())
}

func (db *db) Notes() data.NotesQ {
	return NewNotesQ(db.raw)
}
