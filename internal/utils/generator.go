package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateUsername(firstName, lastName string, rollNo int) string {
	f := strings.ToLower(strings.TrimSpace(firstName))
	l := strings.ToLower(strings.TrimSpace(lastName))
	// remove spaces
	f = strings.ReplaceAll(f, " ", "")
	l = strings.ReplaceAll(l, " ", "")

	return fmt.Sprintf("%s.%s.%d", f, l, rollNo)
}

func GeneratePassword() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 12)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
