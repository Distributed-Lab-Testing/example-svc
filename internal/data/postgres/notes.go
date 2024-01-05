package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Distributed-Lab-Testing/example-svc/internal/data"
	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	notesTable = "notes"

	notesId        = "id"
	notesContent   = "content"
	notesCreatedAt = "created_at"
)

var notesColumns = []string{
	notesContent,
	notesCreatedAt,
}

func NewNotesQ(db *pgdb.DB) data.NotesQ {
	return &notesQ{
		db:  db.Clone(),
		sql: squirrel.Select(fmt.Sprintf("%s.*", notesTable)).From(notesTable),
	}
}

type notesQ struct {
	db  *pgdb.DB
	sql squirrel.SelectBuilder
}

func (q *notesQ) New() data.NotesQ {
	return NewNotesQ(q.db)
}

func (q *notesQ) Select(selector data.NoteSelector) ([]data.Note, error) {
	var result []data.Note
	err := q.db.Select(&result, applyNotesSelector(q.sql, selector))
	return result, err
}

func (q *notesQ) SelectWithPageParams(pageParams pgdb.OffsetPageParams) ([]data.Note, error) {
	var result []data.Note
	err := q.db.Select(&result, pageParams.ApplyTo(q.sql))
	return result, err
}

func (q *notesQ) Get(id int64) (*data.Note, error) {
	var result data.Note
	err := q.db.Get(&result, q.sql.Where(squirrel.Eq{notesId: id}))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &result, err
}

func (q *notesQ) Insert(notes ...data.Note) ([]int64, error) {
	if len(notes) == 0 {
		return nil, nil
	}

	ids := make([]int64, 0)

	query := squirrel.Insert(notesTable).Columns(notesColumns...)
	for _, note := range notes {
		query = query.Values(note.Content, note.CreatedAt)
	}

	query = query.Suffix("returning id")
	return ids, q.db.Select(&ids, query)
}

func (q *notesQ) Delete(id int64) error {
	statement := squirrel.Delete(notesTable).Where(squirrel.Eq{notesId: id})
	return q.db.Exec(statement)
}

func (q *notesQ) UpdateContent(id int64, newContent string) error {
	statement := squirrel.Update(notesTable).Set(notesContent, newContent).Where(squirrel.Eq{notesId: id})

	err := q.db.Exec(&statement)
	if err != nil {
		return err
	}

	return q.db.Exec(&statement)
}

func applyNotesSelector(sql squirrel.SelectBuilder, selector data.NoteSelector) squirrel.SelectBuilder {
	sql = selector.PageParams.ApplyTo(sql, notesId)
	if len(selector.ID) > 0 {
		sql = sql.Where(squirrel.Eq{notesId: selector.ID})
	}
	if len(selector.Content) > 0 {
		sql = sql.Where(squirrel.Eq{notesContent: selector.Content})
	}

	return sql
}
