package uti

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijkmnqprstxz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random Interger between min - max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generate a random String of lenght n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generate a random Owner name
func RandomOwner() string {
	return RandomString(10)
}

func RandomMoney() int64 {
	return RandomInt(100, 1000)
}
