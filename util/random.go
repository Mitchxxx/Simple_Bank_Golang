package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Random Int generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

// Random String generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Random Owner generates a random Owner name
func RandomOwner() string {
	return RandomString(6)
}

// Random Money generates a random amount of Money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Random Currency selects a currency at random
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}