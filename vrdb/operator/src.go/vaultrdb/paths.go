package vaultrdb

import (
	"fmt"
	"os"
)

// returns an absolute path, build on the env variable VRDB_DIRECTORY_ROOT
//
// result example:
//
// [VRDB_DIRECTORY_ROOT] + "/web/swagger/..."
//
// /opt/vaultrdb/web/swagger/...
func RootDir(relalivePath string) string {

	return fmt.Sprintf("%s%s", os.Getenv("VRDB_DIRECTORY_ROOT"), relalivePath)
}
