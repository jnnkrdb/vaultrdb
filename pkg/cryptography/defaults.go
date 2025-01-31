package cryptography

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"os"
	"time"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

// the default encryption passphrase will be created from the ca-certificate.
// the certificate whould be located under /opt/vaultrdb/config/certs/ca-cert.sha
//
// can be changed
const _passphraseSourceFile string = "/opt/vaultrdb/config/certs/ca.crt"

// returns the created hash from the ca.crt file
// if an error occurs, a random string will be returned
func GetPassphraseFromCACertHASH() (result string) {

	// create a default random string from the current time and encode it into b64
	result = base64.RawStdEncoding.EncodeToString([]byte(time.Now().Format(time.RFC3339Nano)))

	// reading the ca.crt content from the requested file
	var cacrt_content []byte
	var err error
	if cacrt_content, err = os.ReadFile(_passphraseSourceFile); err != nil {

		logging.SLog.Info("error creating passphrase from ca.crt, using random created passphrase", "err", err.Error())

		return
	}

	// create a hash from the content of the ca.crt
	var hashGen = sha1.New()
	hashGen.Write(cacrt_content)
	result = hex.EncodeToString(hashGen.Sum(nil))

	logging.SLog.Info("created hash from ca.crt file", "source", _passphraseSourceFile)

	return
}
