package tripledesexample

import "testing"

var (
	triplekey = "12345678" + "12345678" + "12345678"
	plaintext = []byte("Secret message")
)

func BenchmarkEncrypt(b *testing.B) {
	for n := 0; n < 100000; n++ {
		TripleDesEncrypt(plaintext, []byte(triplekey))
	}
}

func BenchmarkDecrypt(b *testing.B) {
	crypted, err := TripleDesEncrypt(plaintext, []byte(triplekey))
	if err != nil {
		panic(err)
	}
	for n := 0; n < 100000; n++ {
		TripleDesDecrypt(crypted, []byte(triplekey))
	}
}
