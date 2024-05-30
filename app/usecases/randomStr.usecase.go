package usecases

import "math/rand"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// validate if randomString is valid
func ValidateRandomString(randomString string) bool {
	for _, r := range randomString {
		if !isLetter(r) {
			return false
		}
	}
	return true
}

func isLetter(r rune) bool {
	for _, l := range letterRunes {
		if r == l {
			return true
		}
	}
	return false
}
