package auth

var tests = []struct {
	Access      AuthorizationRegistration
	AccessTests []struct {
		AuthReq AuthRequest
		Level   string
		Path    string
		Keys    []string
	}
}{
	{Access: AuthorizationRegistration{
		RegistrationID: "random-1",
		Passphrase: "SUPER_SAFE_PASSPHRASE_!",
		Accesses: []AccessConfig{
			AccessConfig{
				Path: "",
			}
		},
	}}
}
