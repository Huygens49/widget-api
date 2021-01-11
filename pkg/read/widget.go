package read

import "time"

type Widget struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	Owner       string    `json:"owner"`
	Value       int       `json:"value"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
