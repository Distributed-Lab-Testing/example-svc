/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Note struct {
	Key
	Attributes NoteAttributes `json:"attributes"`
}
type NoteResponse struct {
	Data     Note     `json:"data"`
	Included Included `json:"included"`
}

type NoteListResponse struct {
	Data     []Note   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustNote - returns Log from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustNote(key Key) *Note {
	var note Note
	if c.tryFindEntry(key, &note) {
		return &note
	}
	return nil
}
