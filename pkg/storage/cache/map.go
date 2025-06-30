package cache

import "fmt"

type Cache map[string]map[string]string

func NewCacheMap() Cache {
	return make(Cache)
}

// -------------------------------------------------------------- required functions

func (cs Cache) ListBuckets() ([]string, error) {
	var bucketList = []string{}
	for bucketKey := range cs {
		bucketList = append(bucketList, bucketKey)
	}
	return bucketList, nil
}

func (cs Cache) WriteBucket(bucket string) error {
	cs[bucket] = make(map[string]string)
	return nil
}

func (cs Cache) DeleteBucket(bucket string) error {
	delete(cs, bucket)
	return nil
}

func (cs Cache) ListKeys(bucket string) ([]string, error) {
	var keyList = []string{}
	for keyKey := range cs[bucket] {
		keyList = append(keyList, keyKey)
	}
	return keyList, nil
}

func (cs Cache) GetKey(bucket, key string) (string, error) {
	b, ok := cs[bucket]
	if !ok {
		return "", fmt.Errorf("bucket %s does not exist or couldn't be found", bucket)
	}
	v, ok := b[key]
	if !ok {
		return "", fmt.Errorf("key [%s/%s] does not exist or couldn't be found", bucket, key)
	}
	return v, nil
}

func (cs Cache) WriteKey(bucket, key, value string) error {
	if _, ok := cs[bucket]; !ok {
		cs.WriteBucket(bucket)
	}
	// recheck if the bucket exists now
	if _, ok := cs[bucket]; !ok {
		return fmt.Errorf("no bucketMap found")
	}
	cs[bucket][key] = value
	return nil
}

func (cs Cache) DeleteKey(bucket, key string) error {
	delete(cs[bucket], key)
	return nil
}

// WARNING! This method kills the complete content of the cachemap
func (cs Cache) Close() error {
	clear(cs)
	return nil
}
