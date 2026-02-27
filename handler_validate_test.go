package main

import "testing"

func TestProfoundBlocker(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no profanity",
			input:    "I love coding in Go",
			expected: "I love coding in Go",
		},
		{
			name:     "single bad word",
			input:    "That was a real kerfuffle",
			expected: "That was a real ****",
		},
		{
			name:     "case sensitivity",
			input:    "FORNAX is a constellation",
			expected: "**** is a constellation",
		},
		{
			name:     "with punctuation",
			input:    "What a sharbert! it was.",
			expected: "What a sharbert! it was.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := profoundBlocker(tt.input)
			if actual != tt.expected {
				t.Errorf("profoundBlocker() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
