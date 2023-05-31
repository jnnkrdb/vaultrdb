package settings

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var PSQL *sql.DB

func init() {
	var err error
	if PSQL, err = sql.Open("postgres", PSQL_CONNECTION); err != nil {
		log.Panicf("error opening postgres connection: %#v", err)
	}

	// running preflight table checks
	for _, sql := range []string{
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS public.vaultsets (id text NOT NULL,	description text NOT NULL, data text NOT NULL, PRIMARY KEY (id));ALTER TABLE IF EXISTS public.vaultsets OWNER to %s;`, POSTGRES_USER),
	} {
		if _, err = PSQL.Exec(sql); err != nil {
			break
		}
	}

	if err != nil {
		log.Panicf("error executing preflight table checks: %#v", err)
	}
}
