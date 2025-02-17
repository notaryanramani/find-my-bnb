package utils

import (
	"regexp"
)

func ValidatePassword(password string) bool {
	if password == "" {
		return false
	}

	if len(password) < 8 {
		return false
	}

	upperCaseRegex := regexp.MustCompile(`[A-Z]`)
	lowerCaseRegex := regexp.MustCompile(`[a-z]`)
	numberRegex := regexp.MustCompile(`[0-9]`)

	if !upperCaseRegex.MatchString(password) {
		return false
	}

	if !lowerCaseRegex.MatchString(password) {
		return false
	}

	if !numberRegex.MatchString(password) {
		return false
	}
	return true
}

func ValidateUsername(username string) bool {
	return len(username) >= 5
}
