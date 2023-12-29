/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateNote struct {
	Key
	Attributes UpdateNoteAttributes `json:"attributes"`
}
type UpdateNoteResponse struct {
	Data     UpdateNote `json:"data"`
	Included Included   `json:"included"`
}

type UpdateNoteListResponse struct {
	Data     []UpdateNote `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustUpdateNote - returns UpdateNote from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateNote(key Key) *UpdateNote {
	var updateNote UpdateNote
	if c.tryFindEntry(key, &updateNote) {
		return &updateNote
	}
	return nil
}
