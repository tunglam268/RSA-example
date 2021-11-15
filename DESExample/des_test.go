package desexample

import "testing"

var (
	key = []byte("2fa6c1e9")
	str = "I love this beautiful world!"
)

func BenchmarkDecrypt(b *testing.B) {
	for n := 0; n < 100000; n++ {
		Decrypt(str, key)
	}
}

func BenchmarkEncrypt(b *testing.B) {
	for n := 0; n < 100000; n++ {
		Encrypt(str, key)
	}
}
