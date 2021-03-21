package random

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

// RandomGenerator utility class to generate ranodm instances to use on tests
type RandomGenerator struct {
	timestamp int64
}

// NewRandomGenerator creates a new random generator instance
func NewRandomGenerator() *RandomGenerator {
	randomizer := &RandomGenerator{
		timestamp: time.Now().UnixNano(),
	}
	randomizer.initialize()
	return randomizer
}

func (randomizer *RandomGenerator) initialize() {
	rand.Seed(randomizer.timestamp)
}

// RandomInt generates a random integer between min and max
func (randomizer *RandomGenerator) RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomFloat generates a random float between min and max
func (randomizer *RandomGenerator) RandomFloat(min, max float64) float64 {
	result := min + (rand.Float64() * (max - min))
	return math.Floor(result*100) / 100
}

// RandomString generates a random string of length n
func (randomizer *RandomGenerator) RandomString(n int) string {
	return randomizer.RandomFromSource(n, alphabet)
}

// RandomNumericString generates a random string of length n
func (randomizer *RandomGenerator) RandomNumericString(n int) string {
	return randomizer.RandomFromSource(n, digits)
}

// RandomAlphaNumericString generates a random string of length n
func (randomizer *RandomGenerator) RandomAlphaNumericString(n int) string {
	return randomizer.RandomFromSource(n, characters)
}

// RandomFromSource generates a random string with the characters on source
func (randomizer *RandomGenerator) RandomFromSource(n int, source string) string {
	var sb strings.Builder
	k := len(source)

	for i := 0; i < n; i++ {
		c := source[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomUsername generates a random username
func (randomizer *RandomGenerator) RandomUsername() string {
	return randomizer.RandomFromSource(8, alphabet)
}

// RandomName generates a random name
func (randomizer *RandomGenerator) RandomName() string {
	return randomizer.RandomFromSource(30, alphabetWithSpace)
}

// RandomEmail generates a random email
func (randomizer *RandomGenerator) RandomEmail() string {
	return randomizer.RandomUsername() + mailSuffix
}

// RandomZipcode generates a random zipcode
func (randomizer *RandomGenerator) RandomZipcode() string {
	return randomizer.RandomNumericString(5)
}
