package data

type DB interface {
	New() DB
	Notes() NotesQ
}
