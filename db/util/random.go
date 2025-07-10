package util

import (
	"math/rand"
	"time"
)

// generate random number
// generate random string of size=6
// generate random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "JPY", "GBP", "AUD"}
	return currencies[rand.Intn(len(currencies))]
}

func RandomAmount() int64 {
	return (int64(rand.Intn(10000)))
}

func RandomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomAccount() int64 {
	var ret int64 = 0
	for i := 0; i < 8; i++ {
		ret = ret*10 + int64(rand.Intn(10))
	}

	return ret
}
