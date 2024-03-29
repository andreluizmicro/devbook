package models

import (
	"time"
)

type User struct {
	Id         uint64    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	NickName   string    `json:"nickname,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
}
