package models

import (
	"github.com/Distributed-Lab-Testing/example-svc/internal/data"
	"github.com/Distributed-Lab-Testing/example-svc/resources"
)

func NoteAsResource(note data.Note) resources.CreateNoteRequest {
	return resources.CreateNoteRequest{
		Data: NoteAsAttribute(note),
	}
}

func NoteAsAttribute(note data.Note) resources.CreateNote {
	return resources.CreateNote{
		Key: resources.NewKeyInt64(note.ID, resources.NOTES),
		Attributes: resources.CreateNoteAttributes{
			Content:   note.Content,
			CreatedAt: note.CreatedAt,
		},
	}
}

func NotesListAsResource(note []data.Note) resources.CreateNoteListRequest {
	resourcesNotes := make([]resources.CreateNote, len(note))
	for i, dataNote := range note {
		resourcesNotes[i] = NoteAsAttribute(dataNote)
	}

	return resources.CreateNoteListRequest{
		Data: resourcesNotes,
	}
}
