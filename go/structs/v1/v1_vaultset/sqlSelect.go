package v1_vaultset

import (
	"fmt"

	"github.com/jnnkrdb/vaultrdb/settings"
)

// select all existing vaultsets in the database
func Select() ([]VaultSet, error) {

	// executing the select statement, with all possible columns
	if rows, err := settings.PSQL.Query("SELECT id, description, data FROM public.vaultsets;"); err != nil {

		return []VaultSet{}, fmt.Errorf("error receiving vaultsets from db: %v", err)

	} else {

		// close the result, when finishing the parsing
		defer rows.Close()

		// caching variables
		var result []VaultSet
		var _id, _desc, _data string

		// parse all rows into the reserved variables
		for rows.Next() {

			// scan the response and parse them into an item
			if err = rows.Scan(&_id, &_desc, &_data); err != nil {

				return []VaultSet{}, fmt.Errorf("error parsing the item into the response struct: %v", err)
			}

			// append the results into the list
			result = append(result, VaultSet{_id, _desc, _data})
		}

		return result, nil
	}
}

// select a specific vaultset via uuid
func SelectByID(uuid string) (VaultSet, error) {

	var vs VaultSet

	// execute the select by id query and parse the data into the struct
	if err := settings.PSQL.QueryRow("SELECT id, description, data FROM public.vaultsets WHERE id=$1;", uuid).
		Scan(&vs.ID, &vs.Description, &vs.Data); err != nil {

		return VaultSet{}, fmt.Errorf("error selecting vaultset with id[%s]: %#v", uuid, err)
	}

	return vs, nil
}
