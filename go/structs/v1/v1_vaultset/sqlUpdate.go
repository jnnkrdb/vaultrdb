package v1_vaultset

import (
	"fmt"

	"github.com/jnnkrdb/vaultrdb/settings"
)

// update a vaultset by object
func Update(vs VaultSet) (VaultSet, error) {

	// execute the update sql
	if _, err := settings.PSQL.Exec("UPDATE public.vaultsets SET description=$1, data=$2 WHERE id=$3;", vs.Description, vs.Data, vs.ID); err != nil {

		return VaultSet{}, fmt.Errorf("error updating the vaultset: %#v", err)
	}

	// check the updated version and take the result
	return SelectByID(vs.ID)
}
