package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// uses the decryption with the given passphrase, returns the decrypted version
// of "text" or an error
//
// Parameters:
//   - `passphrase` : string > contains the string, which is used to decrypt the value of `text`
//   - `text` : string > contains the text to decrypt
func Decrypt(_passphrase, text string) (string, error) {

	if ciphertext, err := base64.URLEncoding.DecodeString(text); err != nil {

		return "", err

	} else {

		if block, err := aes.NewCipher([]byte(_passphrase)); err != nil {

			return "", err

		} else {

			if len(ciphertext) < aes.BlockSize {
				return "", errors.New("ciphertext to short")
			}

			iv := ciphertext[:aes.BlockSize]
			ciphertext = ciphertext[aes.BlockSize:]

			stream := cipher.NewCFBDecrypter(block, iv)

			stream.XORKeyStream(ciphertext, ciphertext)

			return string(ciphertext), nil
		}
	}
}
