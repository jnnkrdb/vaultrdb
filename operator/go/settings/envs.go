package settings

import (
	"fmt"
	"os"
)

var (
	PSQL_CONNECTION = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", PSQL_HOST, PSQL_PORT, PSQL_USER, PSQL_PASSWORD, PSQL_DATABASE)

	PSQL_USER     = os.Getenv("PSQL_USER")
	PSQL_PASSWORD = os.Getenv("PSQL_PASSWORD")
	PSQL_HOST     = os.Getenv("PSQL_HOST")
	PSQL_PORT     = os.Getenv("PSQL_PORT")
	PSQL_DATABASE = os.Getenv("PSQL_PORT")

	CRYPTKEY = os.Getenv("CRYPTKEY")
)
