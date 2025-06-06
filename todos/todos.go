package todos

import (
	"time"
)

type Entry struct {
	CreatedAt time.Time
	Text      string
	ID        int
	Due       *time.Time
}

type Todos struct {
	Entries *[]Entry
	Counter int
}

func New() *Todos {
	var es []Entry
	return &Todos{
		Counter: 0,
		Entries: &es,
	}
}

func (t *Todos) Add(txt string, due *time.Time) {
	id := t.Counter + 1
	t.Counter = id

	entries := *t.Entries

	entries = append(entries, Entry{
		ID:        id,
		Text:      txt,
		CreatedAt: time.Now().In(time.Local),
		Due:       due,
	})
	t.Entries = &entries
}

func (t *Todos) List() []Entry {
	return *t.Entries
}

func (t *Todos) MarkDone(id int) bool {
	var entries []Entry
	removed := false
	for _, e := range *t.Entries {
		if e.ID != id {
			entries = append(entries, e)
		} else {
			removed = true
		}
	}

	t.Entries = &entries
	return removed
}

func (t *Todos) Nuke() {
	var entries []Entry
	t.Entries = &entries
}

func (t *Todos) ResetCounter() {
	var entries []Entry
	curr := *t.Entries
	for i, e := range curr[:] {
		e.ID = i + 1
		entries = append(entries, e)
	}
	l := len(*t.Entries)
	t.Counter = l
	t.Entries = &entries
}
