package util

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	RAND_MIN = 100000
	RAND_MAX = 999999
)

func RandomVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(RAND_MAX - RAND_MIN) + RAND_MIN
	
	return strconv.Itoa(n)
}
