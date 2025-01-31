package cryptography

import (
	"encoding/base64"
	"math/rand"
	"time"
)

// the default encryption passphrase will be created from the ca-certificate.
// the certificate whould be located under /opt/vaultrdb/config/certs/ca-cert.sha
//
// can be changed
//const _passphraseSourceFile string = "/opt/vaultrdb/config/certs/ca.crt"

const _defaultPassphrase string = "+Q3qw2K2NBR1pYMCWFYOcltFl3HxCcqAuWotezDmOS0="

// returns the created hash from the ca.crt file
// if an error occurs, a random string will be returned
func GetPassphraseFromCACertHASH() (result string) {

	randomHash := make([]byte, 32)

	// wait for 1250 ms to avoid random numbers to get generated multiple times
	time.Sleep(1250 * time.Millisecond)

	src := rand.New(rand.NewSource(time.Now().Unix()))

	if _, err := src.Read(randomHash); err != nil {

		result = _defaultPassphrase
		return
	}

	result = base64.StdEncoding.EncodeToString(randomHash)
	// create a default random string from the current time and encode it into b64
	//result = base64.RawStdEncoding.EncodeToString([]byte(time.Now().Format(time.RFC3339Nano)))

	return
}
