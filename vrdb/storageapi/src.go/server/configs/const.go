package configs

// full list of created buckets
// by the service
var Buckets = []string{
	BucketVault,
	BucketWebhooks,
}

// list of buckets, used by the storage api

const (
	BucketVault    string = "vault"
	BucketWebhooks string = "webhooks"
)
