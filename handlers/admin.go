package handlers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	tg "gopkg.in/telebot.v4"

	"tg-sticker-stiller-bot/db"
)

var adminIDs []int64

func InitAdminIDs() {
	adminIDsStr := os.Getenv("ADMIN_IDS")
	if adminIDsStr == "" {
		log.Println("Warning: ADMIN_IDS not set, broadcast feature will be disabled")
		return
	}

	ids := strings.Split(adminIDsStr, ",")
	for _, idStr := range ids {
		idStr = strings.TrimSpace(idStr)
		if idStr == "" {
			continue
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Printf("Warning: Invalid admin ID '%s': %v", idStr, err)
			continue
		}
		adminIDs = append(adminIDs, id)
	}

	if len(adminIDs) > 0 {
		log.Printf("Loaded %d admin ID(s)", len(adminIDs))
	}
}

func IsAdmin(userID int64) bool {
	for _, id := range adminIDs {
		if id == userID {
			return true
		}
	}
	return false
}

func HandleBroadcast(ctx tg.Context, repo *db.Repository) error {
	if !IsAdmin(ctx.Sender().ID) {
		return nil
	}

	args := strings.TrimPrefix(ctx.Text(), "/broadcast ")
	if args == "/broadcast" || args == "" {
		return ctx.Send(
			"📢 *Broadcast Command*\n\n"+
				"Usage: `/broadcast <message>`\n\n"+
				"Send a message to all active users.",
			&tg.SendOptions{ParseMode: tg.ModeMarkdown},
		)
	}

	message := args

	users, err := repo.GetAllActiveUsers()
	if err != nil {
		log.Printf("Failed to get users: %v", err)
		return ctx.Send("❌ Failed to fetch users from database.")
	}

	if len(users) == 0 {
		return ctx.Send("⚠️ No active users found.")
	}

	err = ctx.Send(fmt.Sprintf("📤 Broadcasting to %d users...", len(users)))
	if err != nil {
		return err
	}

	successCount := 0
	failCount := 0
	blockedCount := 0

	for i, user := range users {
		recipient := &tg.User{ID: user.UserID}
		_, sendErr := ctx.Bot().Send(recipient, message)
		if sendErr != nil {
			errStr := sendErr.Error()
			if strings.Contains(errStr, "blocked") || strings.Contains(errStr, "user is deactivated") {
				blockedCount++
			} else {
				log.Printf("Failed to send to user %d (%s): %v", user.UserID, user.Username, sendErr)
				failCount++
			}
		} else {
			successCount++
		}

		if (i+1)%50 == 0 {
			time.Sleep(50 * time.Millisecond)
		}
	}

	result := fmt.Sprintf(
		"✅ *Broadcast Complete*\n\n"+
			"📬 Sent: `%d`\n"+
			"🚫 Blocked: `%d`\n"+
			"❌ Failed: `%d`\n"+
			"📊 Total users: `%d`",
		successCount, blockedCount, failCount, len(users),
	)

	return ctx.Send(result, &tg.SendOptions{ParseMode: tg.ModeMarkdown})
}

func HandleAdminStats(ctx tg.Context, repo *db.Repository) error {
	if !IsAdmin(ctx.Sender().ID) {
		return nil
	}

	userCount, err := repo.GetUserCount()
	if err != nil {
		log.Printf("Failed to get user count: %v", err)
		return ctx.Send("❌ Failed to fetch statistics.")
	}

	stats := fmt.Sprintf(
		"📊 *Bot Statistics*\n\n"+
			"👥 Active users: `%d`\n"+
			"🕒 Server time: `%s`",
		userCount,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	return ctx.Send(stats, &tg.SendOptions{ParseMode: tg.ModeMarkdown})
}
