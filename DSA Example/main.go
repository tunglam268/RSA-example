package main

import (
	"crypto/dsa"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"
	"math/big"
	"os"
)

func main() {
	//// Parameters represents the domain parameters for a key. These parameters can
	// be shared across many keys. The bit length of Q must be a multiple of 8.
	params := new(dsa.Parameters)

	// see http://golang.org/pkg/crypto/dsa/#ParameterSizes
	//GenerateParameters puts a random, valid set of DSA parameters into params. This function can take many seconds, even on fast machines.
	if err := dsa.GenerateParameters(params, rand.Reader, dsa.L1024N160); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	privatekey := new(dsa.PrivateKey)
	privatekey.PublicKey.Parameters = *params

	// GenerateKey generates a public&private key pair. The Parameters of the
	// PrivateKey must already be valid (see GenerateParameters).
	dsa.GenerateKey(privatekey, rand.Reader)

	var pubkey dsa.PublicKey
	pubkey = privatekey.PublicKey

	fmt.Println("Private Key :")
	fmt.Printf("%x \n", privatekey)

	fmt.Println("Public Key :")
	fmt.Printf("%x \n", pubkey)

	// Sign
	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	io.WriteString(h, "This is the message to be signed and verified!")
	signhash := h.Sum(nil)

	r, s, err := dsa.Sign(rand.Reader, privatekey, signhash)
	if err != nil {
		fmt.Println(err)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	fmt.Printf("Signature : %x\n", signature)

	// Verify
	verifystatus := dsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(verifystatus) // should be true

	// we add additional data to change the signhash
	io.WriteString(h, "This message is NOT to be signed and verified!")
	signhash = h.Sum(nil)

	verifystatus = dsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(verifystatus) // should be false
}
