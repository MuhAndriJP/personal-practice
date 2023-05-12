package helper

import (
	"math/rand"
	"time"
)

const (
	EmptyEmail    = "Email Tidak Boleh Kosong"
	EmptyPassword = "Password Tidak Boleh Kosong"
)

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
