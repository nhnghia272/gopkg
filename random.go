package gopkg

import (
	"math/rand"
	"time"
)

const (
	charsetString  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetNumber  = "0123456789"
	charsetSpecial = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	charset        = charsetString + charsetNumber + charsetSpecial
)

func Random(length uint) string {
	bytes := make([]byte, length)
	for i := range bytes {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		bytes[i] = charset[r.Intn(len(charset))]
	}
	return string(bytes)
}

func RandomString(length uint) string {
	bytes := make([]byte, length)
	for i := range bytes {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		bytes[i] = charsetString[r.Intn(len(charsetString))]
	}
	return string(bytes)
}

func RandomNumber(length uint) string {
	bytes := make([]byte, length)
	for i := range bytes {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		bytes[i] = charsetNumber[r.Intn(len(charsetNumber))]
	}
	return string(bytes)
}

func RandomSpecial(length uint) string {
	bytes := make([]byte, length)
	for i := range bytes {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		bytes[i] = charsetSpecial[r.Intn(len(charsetSpecial))]
	}
	return string(bytes)
}
