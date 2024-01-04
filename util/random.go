package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomInt returns a random number between max and min
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(charset)
	for i := 0; i < n; i++ {
		c := charset[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner returns a random string of 6 characters
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney returns a random number between 0 and 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency return a random currency code
func RandomCurrency() string {
	currencies := []string{
		"USD",
		"CAD",
		"EUR",
	}
	return currencies[rand.Intn(len(currencies))]
}
