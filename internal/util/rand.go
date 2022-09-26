package util

import (
	"math/rand"
	"strconv"
	"time"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

const (
	RAND_MIN = 100000
	RAND_MAX = 999999
)

func RandomVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(RAND_MAX-RAND_MIN) + RAND_MIN

	return strconv.Itoa(n)
}

func RandomString(length int) string {
	str := make([]rune, length)

	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}

	return string(str)
}
