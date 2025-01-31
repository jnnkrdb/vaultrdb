package cryptography

import (
	"fmt"
	"strings"
	"testing"
)

func Test_GetPassphraseFromCACertHASH(t *testing.T) {

	var compString string = ""

	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {

			tmp := GetPassphraseFromCACertHASH()
			t.Logf("%s", tmp)

			if strings.Contains(compString, tmp) {
				t.Fatalf("the passphrase (%s) was generated before: %s", tmp, compString)
			}

			compString = fmt.Sprintf("%s__%s", compString, tmp)
		})
	}
}

func Test_randomizedTestCryptography(t *testing.T) {
	var tests = []struct {
		name               string
		source             string
		passphrase         string
		estimateEncryptErr bool
		estimateDecryptErr bool
	}{
		{name: "crpyto-1", source: "deviant-french-unpiloted", passphrase: "%6CdG5&*D!rAf6bora&*F^ho4n7op2Hc"},
		{name: "crpyto-2", source: "chewable-awaken-bloomers", passphrase: "6$GTzD&EJwpL8r9F^D78R@vJ7f%X@KkR"},
		{name: "crpyto-3", source: "blinked-pacify-national", passphrase: "cy$$377!hxxhVgvZHrN@rZE%#^ncY86L"},
		{name: "crpyto-4", source: "uncharted-edgy-thread", passphrase: "uCsJcyRT$3!rP9dqyJ!DZLTkdGP7Jz%j"},
		{name: "crpyto-5", source: "prankster-salvaging-shadow", passphrase: "%5n8!^dy5C!4!PNa8AHUXm7Gnh!&sCXx"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var encryptedList []string
			for i := 0; i < 5; i++ {
				encrypted, err := Encrypt(tt.passphrase, tt.source)
				if err != nil {
					t.Fatalf("error encrypting (%s) with passphrase (%s): %s", tt.source, tt.passphrase, err.Error())
				}
				encryptedList = append(encryptedList, encrypted)
			}

			t.Logf("successfully encrypted (%s) with passphrase (%s) to (%v)", tt.source, tt.passphrase, encryptedList)

			var decryptedList []string
			for _, item := range encryptedList {
				decrypted, err := Decrypt(tt.passphrase, item)
				if err != nil {
					t.Fatalf("error decrypting (%s) with passphrase (%s): %s", item, tt.passphrase, err.Error())
				}

				if decrypted != tt.source {
					t.Fatalf("error: source value (%s) and decrypted (%s) value do not match", item, tt.passphrase)
				}
				encryptedList = append(encryptedList, decrypted)
			}

			t.Logf("successfully decrypted (%v) with passphrase (%s) to (%v), estimated solution for all (%s)", encryptedList, tt.passphrase, decryptedList, tt.source)
		})
	}
}
