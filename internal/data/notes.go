package data

import (
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
)

//go:generate mockery --case underscore --name KycFlowsQ
type NotesQ interface {
	New() NotesQ
	Insert(notes ...Note) ([]string, error)
	Select(selector NoteSelector) ([]Note, error)
	Get() ([]Note, error)
	Delete() error
	UpdateContent(id int64, newContent string) error
}

type NoteSelector struct {
	PageParams   pgdb.OffsetPageParams
	ID           []int64
	Content      []string
	CreatedFrom  *time.Time
	CreatedUntil *time.Time
}

type Note struct {
	ID        int64     `db:"id" structs:"-"`
	Content   string    `db:"content" structs:"content"`
	CreatedAt time.Time `db:"created_at" structs:"created_at"`
}
