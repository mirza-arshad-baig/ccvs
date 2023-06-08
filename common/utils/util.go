package utils

import (
	"fmt"
	"regexp"
)

func ExtractBin(cardNumber string) (string, error) {
	// Remove any non-digit characters from the card number
	reg := regexp.MustCompile("[^0-9]+")
	digitsOnly := reg.ReplaceAllString(cardNumber, "")

	if len(digitsOnly) < 6 {
		return "", fmt.Errorf("invalid credit card number")
	}

	return digitsOnly[:6], nil
}
