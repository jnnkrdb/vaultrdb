package auth

import "encoding/json"

// this object is mainly used by the sideqar/operator
// to request access to specific objects from the vault
type AuthRequest struct {
	RegistrationID string `json:"registrationid"`
	Passphrase     string `json:"passphrase"`
}

func (a AuthRequest) ToBytes() ([]byte, error) {
	return json.Marshal(a)
}

func (a *AuthRequest) FromJson(data []byte) error {
	return json.Unmarshal(data, a)
}
