package db

import "time"

type PackType string

const (
	PackTypeSticker PackType = "sticker"
	PackTypeEmoji   PackType = "emoji"
)

type Pack struct {
	ID           int64     `db:"id"`
	UserID       int64     `db:"user_id"`
	PackName     string    `db:"pack_name"`
	PackTitle    string    `db:"pack_title"`
	PackType     PackType  `db:"pack_type"`
	PackLink     string    `db:"pack_link"`
	StickerCount int       `db:"sticker_count"`
	CreatedAt    time.Time `db:"created_at"`
}

