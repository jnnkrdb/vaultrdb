package v1

type VaultResponse struct {
	ID         int    // id of the connection
	VaultSetID string // uuidv4 of the requested data
	Type       string // configmap/secret
	Namespace  string // namespace of the object
	Name       string // name of the object
}
