package v1

type DeployedSecret deployedObject

type DeployedConfigMap deployedObject

type deployedObject struct {
	Namespace       string `json:"namespace"`
	Name            string `json:"name"`
	ResourceVersion string `json:"resourceversion"`
}
