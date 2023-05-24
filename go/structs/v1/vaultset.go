package v1

import (
	"database/sql"
	"fmt"

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
