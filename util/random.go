package util

import (
	"time"

	"golang.org/x/exp/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) {
	return min + rand.Int63n(max-min+1)
}

func RandomString(length int) {
	var sb string.builder
	k := len(alphabet)
	for i = 0; i < length; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
