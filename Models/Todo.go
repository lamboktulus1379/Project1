package Models

import "time"

type Todo struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UserID      uint      `json:"userId"`
}

func (b *Todo) TableName() string {
	return "todo"
}
