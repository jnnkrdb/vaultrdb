package main

import (
	"log"
	"os"

	"github.com/jnnkrdb/corerdb/crypt"
	"github.com/jnnkrdb/gomw/middlewares/security/cors"
	"github.com/jnnkrdb/k8s/operator"
	"github.com/jnnkrdb/vaultrdb/routines/api"
	"github.com/jnnkrdb/vaultrdb/routines/crds"
	"github.com/jnnkrdb/vaultrdb/settings"
	structs_v1 "github.com/jnnkrdb/vaultrdb/structs/v1"
)

func main() {

	// ---------------------------------------------------
	log.Println("initialize the configs")

	// integrating logs
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	log.SetOutput(os.Stdout)

	// settings the encryption default key
	crypt.SetDefaultPassphrase(settings.CRYPTKEY)

	// settings cors
	cors.SetHeaders("*")
	cors.SetMethods("GET,POST,PUT,DELETE,OPTIONS")
	cors.SetOrigin("*")

	// ---------------------------------------------------
	var err error

	log.Println("initialize the operator struct for the k8s crd requests")

	// initialize the operator struct for the k8s crd requests
	if err = operator.InitCRDOperatorRestClient(structs_v1.GroupName, structs_v1.GroupVersion, structs_v1.AddToScheme); err == nil {

		// start cluster resource handler
		go crds.HandleCRDS()

		// start api backend
		api.HandleAPI()
	}

	if err != nil {
		log.Printf("%#v\n", err)
	}

}
