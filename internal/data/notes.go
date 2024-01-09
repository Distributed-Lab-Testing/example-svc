package data

import (
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
)

type NotesQ interface {
	New() NotesQ
	Insert(notes ...Note) ([]int64, error)
	Select(selector NoteSelector) ([]Note, error)
	Get(id int64) (*Note, error)
	Delete(id int64) error
	UpdateContent(id int64, newContent string) error
}

type NoteSelector struct {
	PageParams pgdb.OffsetPageParams
	ID         []int64
	Content    []string
	CreatedAt  []string
}

type Note struct {
	ID        int64     `db:"id" structs:"-"`
	Content   string    `db:"content" structs:"content"`
	CreatedAt time.Time `db:"created_at" structs:"created_at"`
}
