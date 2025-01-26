package buckets

// full list of created buckets
// by the service
func DefaultBuckets() []string {
	return []string{
		Users,
		Vault,
		Webhooks,
	}
}

// list of buckets, used by the storage api

const (
	Users    string = "users"
	Vault    string = "vault"
	Webhooks string = "webhooks"
)
