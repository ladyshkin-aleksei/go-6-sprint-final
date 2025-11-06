package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	if input == "" {
		return "", errors.New("input line is empty")
	}

	isMorse := strings.ContainsAny(input, ".-/ ")

	if isMorse {
		text := morse.ToText(input)
		if text == "" {
			return "", errors.New("incorrect Morse code")
		}
		return text, nil
	} else {
		morseCode := morse.ToMorse(input)
		if morseCode == "" {
			return "", errors.New("unsupported characters in the text")
		}
		return morseCode, nil
	}
}
