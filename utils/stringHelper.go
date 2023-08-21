package utils

import (
	"math/rand"
	"strings"
)

var passwordCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GeneratePassword(passwordLength int) string {
	var password strings.Builder

	for i := 0; i < passwordLength; i++ {
		random := rand.Intn(len(passwordCharSet))
		password.WriteString(string(passwordCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func Paginate(x []string, pageNum int, pageSize int) []string {
	start := pageNum * pageSize
	sliceLength := len(x)

	if start > sliceLength {
		start = sliceLength
	}

	end := start + pageSize
	if end > sliceLength {
		end = sliceLength
	}

	return x[start:end]
}
