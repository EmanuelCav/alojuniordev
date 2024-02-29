package helper

import (
	"net/mail"
	"regexp"
)

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidateStringAndNumber(text string) bool {
	isStringNumber := regexp.MustCompile(`^[a-zA-Z0-9ñÑ]+$`).MatchString
	return isStringNumber(text)
}

func ValidateString(text string) bool {
	isString := regexp.MustCompile(`^[a-zA-ZñÑ]+$`).MatchString
	return isString(text)
}
