package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// uses the encryption with the given passphrase, returns the encrypted version
// of "text" or an error
//
// Parameters:
//   - `passphrase` : string > contains the string, which is used to encrypt the value of `text`
//   - `text` : string > contains the text to encrypt
func Encrypt(passphrase, text string) (string, error) {

	plaintext := []byte(text)

	block, err := aes.NewCipher([]byte(passphrase))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}
