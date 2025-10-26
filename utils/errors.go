package utils

import (
	"log"
)

// BotError represents a custom error with i18n support
type BotError struct {
	Message   string
	I18nKey   string
	ErrorCode string
}

func (e *BotError) Error() string {
	return e.Message
}

// NewBotError creates a new BotError and logs it
func NewBotError(message, i18nKey, errorCode string) *BotError {
	log.Printf("[%s] %s", errorCode, message)
	return &BotError{
		Message:   message,
		I18nKey:   i18nKey,
		ErrorCode: errorCode,
	}
}

// FailFast panics if error is not nil
func FailFast(err error) {
	log.Printf("Error: %v", err)
	if err != nil {
		panic(err)
	}
}
