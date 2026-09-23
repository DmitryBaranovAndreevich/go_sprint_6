package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ParseData(data string) string {
	input := strings.TrimSpace(data)

	if input == "" {
		return ""
	}

	parts := strings.Split(input, "")

	isText := false

	for _, part := range parts {
		if strings.ToLower(part) != strings.ToUpper(part) {
			isText = true
			break
		}
	}

	if isText {
		return morse.ToMorse(input)
	}

	return morse.ToText(string(input))
}
