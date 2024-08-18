package gopkg

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CopyArray[E comparable](src []E) []E {
	dst := make([]E, len(src))
	copy(dst, src)
	return dst
}

func CopyArray2D[E comparable](s [][]E) [][]E {
	c := make([][]E, len(s))
	for i, v := range s {
		c[i] = make([]E, len(v))
		copy(c[i], v)
	}
	return c
}

func UniqueFunc[E1, E2 comparable](s []E1, f func(E1) E2) []E1 {
	s2 := make([]E1, 0)
	m2 := make(map[E2]E1)
	for _, v := range s {
		k := f(v)
		if _, ok := m2[k]; !ok {
			m2[k] = v
			s2 = append(s2, v)
		}
	}
	return s2
}

func MapFunc[E1, E2 comparable](s []E1, f func(E1) E2) []E2 {
	s2 := make([]E2, len(s))
	for i, v := range s {
		s2[i] = f(v)
	}
	return s2
}

func FilterFunc[E comparable](s []E, f func(E) bool) []E {
	s2 := make([]E, 0)
	for _, v := range s {
		if f(v) {
			s2 = append(s2, v)
		}
	}
	return s2
}

func ReduceFunc[E1, E2 comparable](s []E1, a E2, f func(a E2, e E1, i int) E2) E2 {
	for i, v := range s {
		a = f(a, v, i)
	}
	return a
}

func SomeFunc[E comparable](s []E, f func(E) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

func EveryFunc[E comparable](s []E, f func(E) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

func HashPassword(password string, cost int) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes)
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateSignature(data, secret string) string {
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

func VerifySignature(data, secret, signature string) bool {
	payload := CreateSignature(data, secret)
	return subtle.ConstantTimeCompare([]byte(payload), []byte(signature)) == 1
}

func ObjectID(id string) primitive.ObjectID {
	oid, _ := primitive.ObjectIDFromHex(id)
	return oid
}

func ObjectIDs(ids []string) []primitive.ObjectID {
	oids := []primitive.ObjectID{}
	for _, id := range ids {
		if oid := ObjectID(id); !oid.IsZero() {
			oids = append(oids, oid)
		}
	}
	return oids
}
