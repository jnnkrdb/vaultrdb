package v1_vaultset

import (
	"fmt"
	"log"

	"github.com/jnnkrdb/vaultrdb/settings"
)

// remove a vaultset by uuid
//
// if the vaultset exists, it will be deleted. if it does not exist, there won't be an error
func Delete(uuidv4 string) error {

	// check if the vaultset exists
	if exists, err := vaultsetExists(uuidv4); err != nil {

		return err

	} else {

		// if the vaultset exists, it will be deleted
		if exists {

			// execute the sql
			if _, err := settings.PSQL.Exec("DELETE FROM public.vaultsets WHERE id=$1", uuidv4); err != nil {

				return fmt.Errorf("error removing the vaultset with id[%s]: %#v", uuidv4, err)

			} else {

				log.Printf("VaultSet [%s] removed\n", uuidv4)
			}
		}
	}

	return nil
}
