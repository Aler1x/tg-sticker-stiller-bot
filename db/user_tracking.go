package db

import "time"

type User struct {
	UserID       int64     `db:"user_id"`
	Username     string    `db:"username"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	LanguageCode string    `db:"language_code"`
	IsActive     bool      `db:"is_active"`
	CreatedAt    time.Time `db:"created_at"`
	LastSeenAt   time.Time `db:"last_seen_at"`
}

