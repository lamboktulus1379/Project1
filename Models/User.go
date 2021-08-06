package Models

import "time"

type User struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phoneNumber"`
	CreatedAt   time.Time `json:"createdAt"`
	Todo        []Todo    `json:"todos"`
	Order       []Order   `json:"orders"`
}

func (b *User) TableName() string {
	return "user"
}
