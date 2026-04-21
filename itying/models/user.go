package models

import "time"

type User struct {
	Username string    `json:"username"`
	Id       int       `json:"id"`
	AddTime  time.Time `json:"add_time"`
}

func (User) TableName() string {
	return "users"
}
