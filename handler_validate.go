package main

import (
	"errors"
	"strings"
)

func validateChirp(msg string) (string, error) {
	const maxChirpLength = 140
	if len(msg) > maxChirpLength {
		return "", errors.New("Chirp is too long")
	}
	cleaned := profoundBlocker(msg)
	return cleaned, nil
}

func profoundBlocker(msg string) string {
	forbiddenWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	words := strings.Split(msg, " ")
	for i, word := range words {
		_, ok := forbiddenWords[strings.ToLower(word)]
		if ok {
			words[i] = "****"
		}
	}
	cleanedMsg := strings.Join(words, " ")
	return cleanedMsg
}
