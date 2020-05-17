package gpb

import (
	"sort"
)

// ByRevision sorts to latest revision to the top, i.e. [0]
type ByRevision []*Revision

func (r ByRevision) Len() int           { return len(r) }
func (r ByRevision) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRevision) Less(i, j int) bool { return r[i].Created.Seconds > r[j].Created.Seconds }

func (e *Entry) Latest() *Revision {
	sort.Sort(ByRevision(e.Revisions))
	return e.Revisions[0]
}

func (e *Entry) IsDeleted() bool {
	return e.Latest().GetTombstone()
}

func (e *Entry) Delete(msg string) bool {
	if e.IsDeleted() {
		return false
	}
	e.Revisions = append(e.Revisions, &Revision{
		Message:   msg,
		Tombstone: true,
	})
	return true
}
