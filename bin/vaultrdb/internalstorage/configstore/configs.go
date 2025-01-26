package configstore

import (
	"strings"

	"github.com/jnnkrdb/vaultrdb/pkg/logging"
)

type Config string

func SetConfig(c Config, value string) error {

	split := strings.Split(string(c), ":")

	return DB.WriteKey(split[0], split[1], value)
}

func GetConfig(c Config) string {

	split := strings.Split(string(c), ":")

	res, err := DB.ReadKey(split[0], split[1])
	if err != nil {
		logging.SLog.Error("error reading requested key from configstore",
			"config", string(c),
			"error", err.Error(),
		)
		return ""
	}

	return res
}

// TODO: implement auto bucket creation for desired buckets
var buckets = []string{
	"webhooks",
}

const (
	EncryptionPassphrase Config = "security/encryption:passphrase"
)
