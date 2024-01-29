package utils

import (
	"math/rand"
	"time"
)

func GenerateToken() string {
	return String(40)
}

func GenerateRandomString(len int) string {
	return String(len)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const simpleCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const complexCharset = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func ComplexString(length int) string {
	return StringWithCharset(length, complexCharset)
}

func SimpleString(length int) string {
	return StringWithCharset(length, simpleCharset)
}
