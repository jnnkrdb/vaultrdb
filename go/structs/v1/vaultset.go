package v1

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jnnkrdb/corerdb/crypt"
	"github.com/jnnkrdb/vaultrdb/settings"
)

type VaultSet struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

func Select(aditionalSQLSuffix string) (vsArr []VaultSet, err error) {
	var rows *sql.Rows
	if rows, err = settings.PSQL.Query(fmt.Sprintf("SELECT id, description, data FROM vaultsets%s;", aditionalSQLSuffix)); err != nil {
		err = fmt.Errorf("error receiving vaultsets from db: %#v", err)
	} else {
		for rows.Next() {
			var new VaultSet
			if err = rows.Scan(&new.ID, &new.Description, &new.Data); err != nil {
				err = fmt.Errorf("error parsing vaultset from datasets: %#v", err)
			} else {
				vsArr = append(vsArr, new)
			}
		}
	}
	return
}

func SelectByID(id string) (vs VaultSet, err error) {
	if err = settings.PSQL.QueryRow("SELECT id, description, data FROM vaultsets WHERE id=$1;", id).Scan(&vs.ID, &vs.Description, &vs.Data); err != nil {
		err = fmt.Errorf("error selecting vaultset with id[%s]: %#v", id, err)
	}
	return
}

// insert a new vaultset, if id is empty, a new uuidv4 will be created, else, the uuidv4 in the id field will be validated and used, if the format is correct
func (vs *VaultSet) Create() error {
	var counter = 1
	for {

		if counter > 5 {
			return fmt.Errorf("error: max tries eexceeded")
		}

		// set the first uuid
		if vs.ID == "" {
			vs.ID = uuid.New().String()
		} else {
			if _, err := uuid.Parse(vs.ID); err != nil {
				return fmt.Errorf("error parsing delivered uuidv4: %#v", err)
			}
		}

		if exists, err := vaultsetExists(vs.ID); err != nil {
			return err
		} else if !exists {
			break
		}
		counter++
	}

	// insert the data into the database
	if d, err := crypt.EncryptWithDefault(vs.Data); err != nil {
		return fmt.Errorf("error encrypting data: %v", err)
	} else {
		vs.Data = d
	}

	if _, err := settings.PSQL.Exec("INSERT INTO vaultsets (id, description, data) VALUES ($1, $2, $3);", vs.ID, vs.Description, vs.Data); err != nil {
		return err
	}
	return nil
}

// update a specific vaultset with new values
func (vs *VaultSet) Update() error {
	if _, err := settings.PSQL.Exec("UPDATE vaultsets SET description=$1, data=$2 WHERE id=$3;", vs.Description, vs.Data, vs.ID); err != nil {
		return fmt.Errorf("error updating the vaultset: %#v", err)
	}
	return nil
}

// remove a vaultset via uuidv4
func Delete(id string) error {
	if exists, err := vaultsetExists(id); err != nil {
		return err
	} else {
		if exists {
			if _, err := settings.PSQL.Exec("DELETE FROM vaultsets WHERE id=$1", id); err != nil {
				return fmt.Errorf("error removing the vaultset with id[%s]: %#v", id, err)
			}
		}
	}
	return nil
}

func vaultsetExists(id string) (bool, error) {
	var exists bool
	if err := settings.PSQL.QueryRow("SELECT exists(SELECT 1 FROM vaultsets WHERE id=$1)", id).Scan(&exists); err != nil {
		return false, fmt.Errorf("error checking existance of a vaultsets via uuidv4[%s]: %#v", id, err)
	}
	return exists, nil
}
