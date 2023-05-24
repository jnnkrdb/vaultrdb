package crds

import (
	"log"

	"github.com/jnnkrdb/vaultrdb/structs/v1/v1_vaultrequest"
)

func HandleCRDS() {

	// handling the crds in the cluster to create the configmaps/secrets for the specific requests
	for {
		var (
			vrList v1_vaultrequest.VaultRequestList
			err    error
		)

		if vrList, err = v1_vaultrequest.GetVaultRequestList(); err != nil {
			log.Panicf("error receiving list of vaulterquests: %#v", err)
		}

		for _, vr := range vrList.Items { // handle each vaultrequest in the cluster
			for _, vrdata := range vr.Data { // handle each data request in the vaultrequest

				if e := vrdata.Validate(); e != nil {
					log.Printf("[%s/%s:%s] error validating data: %v\n", vr.Namespace, vr.Name, vrdata.VaultSetID, e)
				}
			}
		}
	}
}
