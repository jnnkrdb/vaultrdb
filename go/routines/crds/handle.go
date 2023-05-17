package crds

import (
	"log"

	structs_v1 "github.com/jnnkrdb/vaultrdb/structs/v1"
)

func HandleCRDS() {

	// handling the crds in the cluster to create the configmaps/secrets for the specific requests
	for {
		var (
			vrList structs_v1.VaultRequestList
			err    error
		)

		if vrList, err = structs_v1.GetVaultRequestList(); err != nil {
			log.Panicf("error receiving list of vaulterquests: %#v", err)
		}

		for _, vaultrequest := range vrList.Items {

			log.Printf("%v\n", vaultrequest)
		}
	}
}
