package todos

import (
	"time"
)

type Entry struct {
	CreatedAt time.Time
	Text      string
	ID        int
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

func (t *Todos) Add(txt string) {
	id := t.Counter + 1
	t.Counter = id

	entries := *t.Entries

	entries = append(entries, Entry{
		ID:        id,
		Text:      txt,
		CreatedAt: time.Now(),
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
