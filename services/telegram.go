package services

import (
	"fmt"
	"log"
	"strings"
	"tg-sticker-stiller-bot/types"
	"tg-sticker-stiller-bot/utils"

	tg "gopkg.in/telebot.v4"
)

func FetchStickerSet(bot *tg.Bot, name string) (*types.StickerSet, error) {
	return utils.WithRetry(func() (*types.StickerSet, error) {
		stickerSet, err := bot.StickerSet(name)
		if err != nil {
			if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "404") {
				log.Printf("Sticker set not found: %s", name)
				return nil, utils.NewBotError(
					fmt.Sprintf("Sticker set not found: %s", name),
					"sticker-not-found",
					"STICKER_SET_NOT_FOUND",
				)
			}
			log.Printf("Telegram API error fetching sticker set: %v", err)
			return nil, utils.NewBotError(
				fmt.Sprintf("Telegram API error: %v", err),
				"fetch-failed",
				"TELEGRAM_API_ERROR",
			)
		}

		return &types.StickerSet{
			Name:     stickerSet.Name,
			Title:    stickerSet.Title,
			Stickers: stickerSet.Stickers,
		}, nil
	})
}

func FetchEmojiSet(bot *tg.Bot, name string) (*types.EmojiSet, error) {
	return utils.WithRetry(func() (*types.EmojiSet, error) {
		stickerSet, err := bot.StickerSet(name)
		if err != nil {
			if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "404") {
				log.Printf("Emoji set not found: %s", name)
				return nil, utils.NewBotError(
					fmt.Sprintf("Emoji set not found: %s", name),
					"emoji-not-found",
					"EMOJI_SET_NOT_FOUND",
				)
			}
			log.Printf("Telegram API error fetching emoji set: %v", err)
			return nil, utils.NewBotError(
				fmt.Sprintf("Telegram API error: %v", err),
				"fetch-emoji-failed",
				"TELEGRAM_API_ERROR",
			)
		}

		return &types.EmojiSet{
			Name:     stickerSet.Name,
			Title:    stickerSet.Title,
			Stickers: stickerSet.Stickers,
		}, nil
	})
}
