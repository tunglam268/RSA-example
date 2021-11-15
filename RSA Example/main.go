package rsaexample

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	//tạo mã
	privateKey, err := rsa.GenerateKey(rand.Reader, int(2048))
	if err != nil {
		fmt.Println(err)
	}

	publicKey := privateKey.PublicKey
	secretMessage := "This is super secret message!"

	//mã hóa key kèm theo thông điệp
	encryptedMessage := RSA_OAEP_Encrypt(secretMessage, publicKey)

	fmt.Println("Cipher Text:", encryptedMessage)

	//giải mã
	RSA_OAEP_Decrypt(encryptedMessage, *privateKey)
}

//mã hõa
func RSA_OAEP_Encrypt(secretMessage string, key rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(secretMessage), label)
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}

//giải mã
func RSA_OAEP_Decrypt(cipherText string, privKey rsa.PrivateKey) string {
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &privKey, ct, label)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Plaintext:", string(plaintext))
	return string(plaintext)
}
