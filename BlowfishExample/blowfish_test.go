package blowfishexample

import (
	"testing"
)

var (
	plaintext = []byte("this is the plaintext string")         // plaintext
	secretkey = []byte("1234567890abcdefghijklmnopqrstuvwxyz") // dat key
)

func BenchmarkDecrypt(b *testing.B) {
	paddedplaintext := blowfishChecksizeAndPad(plaintext)
	for n := 0; n < 100000; n++ {
		blowfishEncrypt([]byte(paddedplaintext), secretkey)
	}
}

func BenchmarkEncrypt(b *testing.B) {
	paddedplaintext := blowfishChecksizeAndPad(plaintext)
	for n := 0; n < 100000; n++ {
		blowfishDecrypt([]byte(paddedplaintext), plaintext)
	}
}
