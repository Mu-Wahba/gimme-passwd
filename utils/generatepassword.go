package utils

import (
	"math/rand"
)

const charset = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ%^&*()$!"

func GeneratePasswd(length int) string {
	r := rand.New(rand.NewSource(rand.Int63()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}
