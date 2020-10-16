package user

import (
	"time"
)

type User struct {
	ID           int
	Name         string
	Occupation   string
	Email        string
	PasswordHash string
	Avatar       string
	Role         string
	CreatedAt    time.Time
	UpatedAt     time.Time
}
