package reading

import "time"

type Widget struct {
	ID          uint
	Description string
	Owner       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
