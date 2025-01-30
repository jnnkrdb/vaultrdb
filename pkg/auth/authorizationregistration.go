package auth

import (
	"encoding/json"
	"strings"
)

const (
	ACCESSLEVEL_LIST   string = "list"
	ACCESSLEVEL_FIND   string = "find"
	ACCESSLEVEL_WRITE  string = "write"
	ACCESSLEVEL_DELETE string = "delete"
)

// this object is mandatory, to determine the objects a sideqar has to receive
// by the vault.
// the vault inherits the authorization registration and therefore all assets and
// access rights to these assets
type AuthorizationRegistration struct {
	// this is the id the sideqar will have, along with the auth passphrase
	// so the ruleset for the sideqar can be correlated
	RegistrationID string `json:"registrationid"`
	Passphrase     string `json:"passphrase"`

	// the accesses describe which objects the holder of the authorizationregistration
	// can access and with which permissions
	Accesses []AccessConfig `json:"accesses"`
}

// this object defines the kind of access to an specifc
type AccessConfig struct {
	Path   string   `json:"path"`
	Keys   []string `json:"keys"`
	Levels []string `json:"level"`
}

/*
	EXAMPLE:
		{
			"registrationid": "why-spoilage-cartwheel-monastery",
			"passphrase": "Bw9$Ga&7GrXm2K7cSH@DrT23YKe6tyin",
			"accesses": [
				{
					"path": "@bucket1.bucket2.bucket3",
					"keys": [ "key_1", "key_a", "key_b" ],
					"level": [ "list", "find", "write" ]
				}
			]
		}

		With these access rights, the corresponding registrator can list the
		three mentioned keys from the mentioned buckets, find them (receive their values)
		and write to them
*/

func (a AuthorizationRegistration) ToBytes() ([]byte, error) {
	return json.Marshal(a)
}

func (a *AuthorizationRegistration) FromJson(data []byte) error {
	return json.Unmarshal(data, a)
}

func (a AuthorizationRegistration) IstAuthenticated(passphrase string) bool {
	return (a.Passphrase == passphrase)
}

func (a AuthorizationRegistration) HasAccess(path, key, accesslevel string) bool {

	for _, aAccess := range a.Accesses {
		if path != aAccess.Path {
			continue
		}

		for _, aKey := range aAccess.Keys {
			if aKey != key {
				continue
			}

			if strings.Contains(strings.Join(aAccess.Levels, "%"), accesslevel) {
				return true
			}
		}
	}

	return false
}
