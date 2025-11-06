package domain

import (
	"time"
)

type User struct {
	ID           string                 `db:"id"`
	CreatedAt    time.Time              `db:"created_at"`
	UpdatedAt    time.Time              `db:"updated_at"`
	Email        string                 `db:"email"`
	Role         string                 `db:"role"`
	Status       string                 `db:"status"`
	PasswordHash string                 `db:"password_hash"`
	Attrs        map[string]interface{} `db:"attrs"`
}
