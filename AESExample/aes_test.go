package aesexample

import (
	"log"
	"testing"
)

var (
	CIPHER_KEY = []byte("0123456789012345")
	msg        = "Secret message !!!"
)

func BenchmarkEncrypt(b *testing.B) {
	for n := 0; n < 100000; n++ {
		Encrypt(CIPHER_KEY, msg)
	}
}

func BenchmarkDecrypt(b *testing.B) {
	encrypted, err := Encrypt(CIPHER_KEY, msg)
	if err != nil {
		log.Println(err)
	}
	for n := 0; n < 100000; n++ {
		Decrypt(CIPHER_KEY, encrypted)
	}
}
