package v1_vaultset

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jnnkrdb/vaultrdb/settings"
)

// insert a new vaultset
//
// if id is empty, a new uuidv4 will be created
// else, the uuidv4 in the id field will be validated and used,
// if the format is correct
func Insert(vs *VaultSet) error {

	// set a max tries counter loop to validate the id of the vaultset
	for counter := 1; counter < 6; counter++ {

		// if the id is unset, then set a new uuidv4
		if len(vs.ID) == 0 {
			vs.ID = uuid.New().String()
		}

		// validate the format of the uuidv4
		if _, err := uuid.Parse(vs.ID); err != nil {
			return fmt.Errorf("error parsing delivered uuidv4[%s]: %#v", vs.ID, err)
		}

		// check the existance of a vaultset with the given id
		if exists, err := vaultsetExists(vs.ID); err != nil {

			return err

		} else {

			// if the vaultset with the given id, does not exist, then create it
			if !exists {

				// encrypt the value of vs.Data
				//				if data, err := crypt.EncryptWithDefault(vs.Data); err != nil {
				//
				//					return fmt.Errorf("error encrypting data: %v", err)
				//
				//				} else {
				//
				//					// update the vaultset data, to the encrypted version
				//					vs.Data = data
				//				}

				// insert the vaultset into the database
				if _, err := settings.PSQL.Exec("INSERT INTO public.vaultsets (id, description, data) VALUES ($1, $2, $3);", vs.ID, vs.Description, vs.Data); err != nil {

					return fmt.Errorf("error isnerting the vaultset: %#v", err)
				}

				return nil
			}
		}
	}

	// exit this function with a max tries are exceeded error
	return fmt.Errorf("error: max tries eexceeded")
}
