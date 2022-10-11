package common

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var upperRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var intRunes = []rune("0123456789")

func RandUpperString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = upperRunes[rand.Intn(len(upperRunes))]
	}
	return string(b)
}
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func RandIntRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = intRunes[rand.Intn(len(intRunes))]
	}
	return string(b)
}
func Md5String(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	return fmt.Sprintf("%x", w.Sum(nil))
}
