package util

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet          = "abcdefghijklmnopqrstuvwxyz"
	alphabetWithSpace = alphabet + " "
	digits            = "1234567890"
	characters        = alphabetWithSpace + digits
	mailSuffix        = "@gmail.com"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomFloat generates a random float between min and max
func RandomFloat(min, max float64) float64 {
	result := min + (rand.Float64() * (max - min))
	return math.Floor(result*100) / 100
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	return RandomFromSource(n, alphabet)
}

// RandomNumericString generates a random string of length n
func RandomNumericString(n int) string {
	return RandomFromSource(n, digits)
}

// RandomAlphaNumericString generates a random string of length n
func RandomAlphaNumericString(n int) string {
	return RandomFromSource(n, characters)
}

// RandomFromSource generates a random string with the characters on source
func RandomFromSource(n int, source string) string {
	var sb strings.Builder
	k := len(source)

	for i := 0; i < n; i++ {
		c := source[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomUsername generates a random username
func RandomUsername() string {
	return RandomFromSource(8, alphabet)
}

// RandomName generates a random name
func RandomName() string {
	return RandomFromSource(30, alphabetWithSpace)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return RandomUsername() + mailSuffix
}

// RandomZipcode generates a random zipcode
func RandomZipcode() string {
	return RandomNumericString(5)
}
