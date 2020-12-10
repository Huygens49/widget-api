package reading

import "time"

type Widget struct {
	ID          uint
	Description string
	Owner       string
	Value       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
