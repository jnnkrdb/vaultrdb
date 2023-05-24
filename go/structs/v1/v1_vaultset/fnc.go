package v1_vaultset

import (
	"fmt"

	"github.com/jnnkrdb/vaultrdb/settings"
)

// check if a vaultset exists in the database
func vaultsetExists(uuidv4 string) (bool, error) {
	var exists bool
	if err := settings.PSQL.QueryRow("SELECT exists(SELECT 1 FROM vaultsets WHERE id=$1)", uuidv4).Scan(&exists); err != nil {
		return false, fmt.Errorf("error checking existance of a vaultsets via uuidv4[%s]: %#v", uuidv4, err)
	}
	return exists, nil
}
