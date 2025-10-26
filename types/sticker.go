package types

import tg "gopkg.in/telebot.v4"

type StickerType string

const (
	StickerTypeRegular StickerType = "regular"
	StickerTypeEmoji   StickerType = "custom_emoji"
)

type DownloadedSticker struct {
	Path    string
	Sticker tg.Sticker
}

type DownloadedEmoji struct {
	Path  string
	Emoji tg.Sticker
}
