package storage

// this is the typical interface to describe the structure
// of a storage implementation
type Storage interface {

	// Receive a list of Buckets
	ListBuckets() ([]string, error)

	// delete a specific key-value from a bucket
	DeleteBucket(string) error

	// List the keys from a bucket, without the value
	ListKeys(string) ([]string, error)

	// receive a specific key-value from a bucket
	GetKey(string, string) (string, error)

	// write a defined key-value into a given bucket
	Write(string, string, string) error

	// delete a specific key-value from a bucket
	DeleteKey(string, string) error

	// close the connection to the storage
	Close() error
}
