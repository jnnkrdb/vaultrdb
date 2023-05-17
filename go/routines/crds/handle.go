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

		for _, vr := range vrList.Items { // handle each vaultrequest in the cluster

			log.Printf("[%s/%s]\n", vr.Namespace, vr.Name)

			for _, vrdata := range vr.Data { // handle each data request in the vaultrequest

				log.Printf("\t\t%v\n", vrdata)
			}
		}
	}
}
