/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateNote struct {
	Key
	Attributes CreateNoteAttributes `json:"attributes"`
}
type CreateNoteRequest struct {
	Data     CreateNote `json:"data"`
	Included Included   `json:"included"`
}

type CreateNoteListRequest struct {
	Data     []CreateNote `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustCreateNote - returns CreateNote from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateNote(key Key) *CreateNote {
	var createNote CreateNote
	if c.tryFindEntry(key, &createNote) {
		return &createNote
	}
	return nil
}
