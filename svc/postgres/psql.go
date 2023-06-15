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
	var usePsql, set bool = true, false

	if psql_host, set = os.LookupEnv("POSTGRES_HOST"); !set {
		_log.Info("postgres config variable missing", "POSTGRES_HOST", psql_host)
		usePsql = false
	}
	if psql_port, set = os.LookupEnv("POSTGRES_PORT"); !set {
		_log.Info("postgres config variable missing", "POSTGRES_PORT", psql_port)
		usePsql = false
	}
	if psql_user, set = os.LookupEnv("POSTGRES_USER"); !set {
		_log.Info("postgres config variable missing", "POSTGRES_USER", psql_user)
		usePsql = false
	}
	if psql_pass, set = os.LookupEnv("POSTGRES_PASSWORD"); !set {
		_log.Info("postgres config variable missing", "POSTGRES_PASSWORD", psql_pass)
		usePsql = false
	}
	if psql_db, set = os.LookupEnv("POSTGRES_DB"); !set {
		_log.Info("postgres config variable missing", "POSTGRES_DB", psql_db)
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
		_log.Info("postgres connection created")
		PSQL = r
		USEPOSTGRES = true
	}

	if err := PSQLConnected(); err != nil {
		_log.Error(err, "error pinging postgres server")
		USEPOSTGRES = false
		return err
	}

	if _, err := PSQL.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS public.vault (psqlid text NOT NULL, data text NOT NULL, PRIMARY KEY (psqlid)); ALTER TABLE IF EXISTS public.vault OWNER to %s;", psql_user)); err != nil {
		_log.Error(err, "error executing preflight statement")
		return err
	}

	return nil
}

// disconnect from postgres
func Disconnect() error {
	return PSQL.Close()
}
