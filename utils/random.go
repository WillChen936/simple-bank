package utils

import (
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(uint64(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(length int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < length; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func CreateRandomOwner() string {
	firstname := RandomString(int(RandomInt(4, 8)))
	lastname := RandomString(int(RandomInt(2, 6)))
	return firstname + " " + lastname
}

func CreateRandomMoney() int64 {
	return RandomInt(0, 1000)
}

func CreateRandomCurrency() string {
	currencies := Currencies
	size := len(currencies)

	return currencies[rand.Intn(size)]
}
