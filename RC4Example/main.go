package main

import (
	"crypto/rc4"
	"fmt"
	"log"
)

func main() {
	// ENCRYPT
	c, err := rc4.NewCipher([]byte("dsadsad"))
	if err != nil {
		log.Fatalln(err)
	}
	src := []byte("asdsad")
	fmt.Println("Plaintext: ", src)

	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)
	fmt.Println("Ciphertext: ", dst)

	// DECRYPT
	c2, err := rc4.NewCipher([]byte("dsadsad"))
	if err != nil {
		log.Fatalln(err)
	}
	src2 := make([]byte, len(dst))
	c2.XORKeyStream(src2, dst)
	fmt.Println("Plaintext': ", src2)
}
