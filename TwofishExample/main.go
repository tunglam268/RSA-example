/*
Package secretbox is used to authenticate and secure small messages. It
provides an interface similar to NaCL, but uses Twofish-256 with a Keccak
MAC construction.

Messages should be secured using the Seal function, and recovered using
the Open function. A box (an authenticated and encrypted message) will be
Overhead bytes longer than the message it came from; this package will not
obscure the length of the message. Keys, if they are not generated using
the GenerateKey function, should be KeySize bytes long.

The boxes used in this package are suitable for 50-year security,
assuming proper key management.
*/
package secretbox

import (
	"crypto/cipher"
	"crypto/rand"
	"crypto/subtle"
	"errors"
	"io"

	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/twofish"
)

// The keys used by secretbox are 32 bytes for the encryption key,
// and 48 bytes for the MAC key. BoxVersion indicates the current
// version of the box format.
const (
	BoxVersion = 1
	KeySize    = 80
)

// The cipher lists specify the ciphers in use in the box. This
// primarily exists to facilitate a future upgrade to Serpent and
// Skein.
const (
	CipherTwofishKeccak = iota + 1
)

// PRNG is the io.Reader from which random data is read.
var PRNG = rand.Reader

const twofishKeyLen = 32
const tagLength = 48 // Size of a Keccak-384-MAC tag.

// Overhead is the number of additional bytes a secure box will
// contain in addition to the message length. A secured message contains
// an additional Twofish block (containing the IV), the message tag,
// which is currently a Keccak-384 tag of 48 bytes, and a byte each
// for the box format version in use and the cipher list in use.
var Overhead = twofish.BlockSize + tagLength + 2

// GenerateKey returns a suitable new key.
func GenerateKey() (*[KeySize]byte, bool) {
	key := new([KeySize]byte)
	_, err := io.ReadFull(PRNG, key[:])
	if err != nil {
		return nil, false
	}
	return key, true
}

func keccakMAC(key, in []byte) []byte {
	h := sha3.NewLegacyKeccak256()
	h.Write(key)
	h.Write(in)
	return h.Sum(nil)
}

func generateIV() ([]byte, error) {
	iv := make([]byte, twofish.BlockSize)
	_, err := io.ReadFull(PRNG, iv)
	return iv, err
}

// Zero wipes the byte slice passed in.
func Zero(in []byte) {
	inLen := len(in)
	for i := 0; i < inLen; i++ {
		in[i] ^= in[i]
	}
}

func encrypt(key, in []byte) ([]byte, error) {
	if len(key) != twofishKeyLen {
		return nil, errors.New("secretbox: invalid key size")
	}
	c, err := twofish.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ctext := make([]byte, len(in)+twofish.BlockSize)
	iv, err := generateIV()
	if err != nil {
		return nil, err
	}
	copy(ctext, iv)

	ctr := cipher.NewCTR(c, iv)
	ctr.XORKeyStream(ctext[twofish.BlockSize:], in)
	return ctext, nil
}

func decrypt(key, ctext []byte) ([]byte, error) {
	if len(key) != twofishKeyLen {
		return nil, errors.New("secretbox: invalid key size")
	}
	if len(ctext) < twofish.BlockSize {
		return nil, errors.New("secretbox: invalid ciphertext")
	}

	iv := ctext[:twofish.BlockSize]
	ct := ctext[twofish.BlockSize:]

	c, err := twofish.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ptext := make([]byte, len(ct))
	ctr := cipher.NewCTR(c, iv)
	ctr.XORKeyStream(ptext, ct)
	return ptext, nil
}

// Seal returns a box that stores an encrypted and authenticated message
// that will be Overhead bytes longer than the original input.
func Seal(in []byte, key *[KeySize]byte) (box []byte, ok bool) {
	ctext, err := encrypt(key[:twofishKeyLen], in)
	if err != nil {
		return
	}

	// The box needs to contain the ciphertext, the message
	// tag, and two bytes for the version and cipher in use.
	box = make([]byte, len(ctext)+2+tagLength)
	box[0] = BoxVersion
	box[1] = CipherTwofishKeccak
	copy(box[2:], ctext)

	tag := keccakMAC(key[twofishKeyLen:], box[:2+len(ctext)])
	copy(box[2+len(ctext):], tag)
	ok = true

	return
}

// Open authenticates and decrypts a box secured using the Seal
// function. The returned message will be Overhead bytes shorter than
// the box.
func Open(box []byte, key *[KeySize]byte) (msg []byte, ok bool) {
	if len(box) < Overhead {
		return
	} else if box[0] != BoxVersion {
		// Note: upgrades to the format require a conversion
		// tool
		return
	} else if box[1] != CipherTwofishKeccak {
		// Future cipher suites will require handling here.
		return
	}

	tagStart := len(box) - tagLength
	tag := keccakMAC(key[twofishKeyLen:], box[:tagStart])
	if subtle.ConstantTimeCompare(tag, box[tagStart:]) != 1 {
		return
	}

	out, err := decrypt(key[:twofishKeyLen], box[2:tagStart])
	if err != nil {
		return
	}
	msg = out
	ok = true

	return
}
