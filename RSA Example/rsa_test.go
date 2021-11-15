package rsaexample

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"testing"
)

func BenchmarkEncrypt(b *testing.B) {
	privateKey, err := rsa.GenerateKey(rand.Reader, int(2048))
	if err != nil {
		fmt.Println(err)
	}
	pubKey := privateKey.PublicKey
	secretMessage := "This is super secret message!"
	for n := 0; n < 100000; n++ {
		RSA_OAEP_Encrypt(secretMessage, pubKey)
	}
}

func BenchmarkDecrypt(b *testing.B) {
	privateKey, err := rsa.GenerateKey(rand.Reader, int(2048))
	if err != nil {
		fmt.Println(err)
	}
	pubKey := privateKey.PublicKey
	secretMessage := "This is super secret message!"
	for n := 0; n < 100000; n++ {
		encrypt := RSA_OAEP_Encrypt(secretMessage, pubKey)
		RSA_OAEP_Decrypt(encrypt, *privateKey)
	}
}
