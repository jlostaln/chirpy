package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Couldn't find Authorization header")
	}
	prefix := "ApiKey "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", fmt.Errorf("Couldn't find Bearer token")
	}
	apiKey := strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))

	return apiKey, nil
}
