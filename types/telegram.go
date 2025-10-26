package types

import (
	"encoding/json"

	tg "gopkg.in/telebot.v4"
)

type TelegramAPIResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code,omitempty"`
	Description string          `json:"description,omitempty"`
}

type StickerSet struct {
	Name     string       `json:"name"`
	Title    string       `json:"title"`
	Stickers []tg.Sticker `json:"stickers"`
}

type EmojiSet struct {
	Name     string       `json:"name"`
	Title    string       `json:"title"`
	Stickers []tg.Sticker `json:"stickers"`
}

type FileResponse struct {
	FilePath     string `json:"file_path"`
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
}
