package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-logr/logr"

	_ "github.com/lib/pq"
)

var (

	// should the operator use a connection to postgres or not
	USEPOSTGRES bool = false

	PSQL *sql.DB
)

// handle the postgres connection tests
func PSQLConnected() error {

	if !USEPOSTGRES {
		return nil
	}

	if err := PSQL.Ping(); err != nil {
		return err
	}

	return nil
}

// create the postgres connection
func Connect(_log logr.Logger) error {

	// check if postgres url is configured
	var psql_host, psql_port, psql_user, psql_pass, psql_db string
	var usePsql bool = true
	for _, c := range []struct {
		ENV string
		Var *string
	}{
		{"POSTGRES_HOST", &psql_host},
		{"POSTGRES_PORT", &psql_port},
		{"POSTGRES_USER", &psql_user},
		{"POSTGRES_PASSWORD", &psql_pass},
		{"POSTGRES_DB", &psql_db},
	} {
		if res, set := os.LookupEnv(c.ENV); set {
			c.Var = &res
			continue
		}
		_log.Info("environment variable not set", "env", c.ENV)
		usePsql = false
	}

	if !usePsql {
		_log.Info("receiving postgres configs", "usePostgres", usePsql)
		return nil
	}

	if r, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		psql_host, psql_port, psql_user, psql_pass, psql_db,
	)); err != nil {
		_log.Error(err, "error connecting to postgres server")
		return err
	} else {
		_log.Info("postgres connection upserted")
		PSQL = r
		USEPOSTGRES = true
	}

	if err := PSQLConnected(); err != nil {
		_log.Error(err, "error pinging postgres server")
		USEPOSTGRES = false
		return err
	}

	if _, err := PSQL.Exec("CREATE TABLE IF NOT EXISTS public.vault (psqlid text NOT NULL, data text NOT NULL, PRIMARY KEY (psqlid)); ALTER TABLE IF EXISTS public.vault OWNER to %s;", psql_user); err != nil {
		_log.Error(err, "error executing preflight statement")
		return err
	}

	return nil
}

// disconnect from postgres
func Disconnect() error {
	return PSQL.Close()
}
