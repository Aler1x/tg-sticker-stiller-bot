package services

import (
	"fmt"
	"log"
	"strings"
	"tg-sticker-stiller-bot/db"
	"tg-sticker-stiller-bot/types"
	"tg-sticker-stiller-bot/utils"
	"time"

	tg "gopkg.in/telebot.v4"
)

type ProgressCallback func(current, total int)

func CreateStickerSet(bot *tg.Bot, userID int64, botname string, title string, stickers []tg.Sticker, stickerType types.StickerType, repo *db.Repository, progressCallback ProgressCallback) (string, error) {
	downloadedStickers := DownloadAllStickers(bot, stickers)
	if len(downloadedStickers) == 0 {
		log.Printf("No stickers could be downloaded for user %d", userID)
		return "", fmt.Errorf("no stickers could be downloaded")
	}

	filePaths := make([]string, len(downloadedStickers))
	for i, ds := range downloadedStickers {
		filePaths[i] = ds.Path
	}
	defer utils.CleanupFiles(filePaths)

	normalizedName := utils.NormalizePackName(title)
	setName := utils.GenerateSetName(normalizedName, botname)

	user := &tg.User{ID: userID}

	var telegramStickerType tg.StickerSetType
	var packLink string
	var dbPackType db.PackType

	if stickerType == types.StickerTypeEmoji {
		telegramStickerType = tg.StickerCustomEmoji
		packLink = fmt.Sprintf("https://t.me/addemoji/%s", setName)
		dbPackType = db.PackTypeEmoji
	} else {
		telegramStickerType = tg.StickerRegular
		packLink = fmt.Sprintf("https://t.me/addstickers/%s", setName)
		dbPackType = db.PackTypeSticker
	}

	// Create set with first sticker only
	firstSticker := downloadedStickers[0]
	emoji := firstSticker.Sticker.Emoji
	if emoji == "" {
		emoji = "😀"
	}

	firstInput := tg.InputSticker{
		File:     tg.FromDisk(firstSticker.Path),
		Format:   utils.GetStickerFormat(firstSticker.Sticker),
		Emojis:   []string{emoji},
		Keywords: []string{},
	}

	stickerSet := &tg.StickerSet{
		Type:  telegramStickerType,
		Name:  setName,
		Title: title,
		Input: []tg.InputSticker{firstInput},
	}

	err := bot.CreateStickerSet(user, stickerSet)
	if err != nil {
		if isNameTakenError(err) {
			log.Printf("Sticker set name already exists: %s for user %d", title, userID)
			return "", utils.NewBotError(
				fmt.Sprintf("Sticker set name already exists: %s", title),
				"name-taken",
				"STICKER_SET_NAME_TAKEN",
			)
		}
		log.Printf("Failed to create sticker set: %v", err)
		return "", err
	}

	// Add remaining stickers one by one
	totalStickers := len(downloadedStickers)

	// start from 0 more user friendly
	if progressCallback != nil && totalStickers > 1 {
		progressCallback(0, totalStickers)
	}

	for i := 1; i < totalStickers; i++ {
		stickerData := downloadedStickers[i]
		emoji := stickerData.Sticker.Emoji
		if emoji == "" {
			emoji = "😀"
		}

		inputSticker := tg.InputSticker{
			File:     tg.FromDisk(stickerData.Path),
			Format:   utils.GetStickerFormat(stickerData.Sticker),
			Emojis:   []string{emoji},
			Keywords: []string{},
		}

		err := bot.AddStickerToSet(user, setName, inputSticker)
		if err != nil {
			log.Printf("Failed to add sticker %d/%d to set: %v", i+1, totalStickers, err)
			// Continue adding other stickers even if one fails
		}

		// Update progress every 10 stickers or on the last one
		if progressCallback != nil {
			if (i+1)%10 == 0 || i+1 == totalStickers {
				progressCallback(i+1, totalStickers)
			}
		}

		// Delay to avoid rate limiting (1ms between 5 stickers)
		if i < totalStickers-1 && i%5 != 0 {
			time.Sleep(time.Millisecond)
		}
	}

	if repo != nil {
		pack := &db.Pack{
			UserID:       userID,
			PackName:     setName,
			PackTitle:    title,
			PackType:     dbPackType,
			PackLink:     packLink,
			StickerCount: len(downloadedStickers),
		}
		if err := repo.CreatePack(pack); err != nil {
			log.Printf("Failed to save pack to database: %v", err)
		}
	}

	return packLink, nil
}

func isNameTakenError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return strings.Contains(errStr, "name is already occupied") || strings.Contains(errStr, "409")
}
