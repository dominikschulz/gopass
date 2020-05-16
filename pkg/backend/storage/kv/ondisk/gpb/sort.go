package gpb

import "sort"

type ByRevision []*Revision

func (r ByRevision) Len() int           { return len(r) }
func (r ByRevision) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRevision) Less(i, j int) bool { return r[i].Created.Seconds > r[j].Created.Seconds }

func (e *Entry) Latest() *Revision {
	sort.Sort(ByRevision(e.Revisions))
	return e.Revisions[len(e.Revisions)-1]
}
