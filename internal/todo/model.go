package todo

import "time"

type Todo struct {
	ID        string
	Text      string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
	OwnerID   string
}
