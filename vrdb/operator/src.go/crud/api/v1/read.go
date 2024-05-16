package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
	"github.com/jnnkrdb/vaultrdb/crud/config"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func READ(w http.ResponseWriter, r *http.Request) {

	// register the rest variables
	vars := mux.Vars(r)
	config.CrudLog.Info("got vars", "vars", vars)

	// declare a new client.Object for the apiserver request
	var reqObj client.Object

	switch vars["kind"] {
	case "vrdbconfigs":
		reqObj = &jnnkrdbdev1.VRDBConfig{}
	case "vrdbsecrets":
		reqObj = &jnnkrdbdev1.VRDBSecret{}
	case "vrdbrequests":
		http.Error(w, "type not implemented", http.StatusNotImplemented)
		return
		//reqObj = &jnnkrdbdev1.VRDBRequest{}
		// case "namespaces":
		// 	reqObj = &metav1.Namespace{}
	default:
		http.Error(w, "type not implemented", http.StatusNotImplemented)
		return
	}

	if err := config.KClient.Get(r.Context(), types.NamespacedName{
		Namespace: vars["namespace"],
		Name:      vars["name"],
	}, reqObj, &client.GetOptions{}); err != nil {

		config.CrudLog.Error(err, "error receiving object",
			"object.kind", reqObj.GetObjectKind().GroupVersionKind().Kind,
			"object.group", reqObj.GetObjectKind().GroupVersionKind().Group,
			"object.version", reqObj.GetObjectKind().GroupVersionKind().Version,
		)

		if errors.IsNotFound(err) {
			http.Error(w, "error object not found", http.StatusNotFound)
		} else {
			http.Error(w, "error receiving object", http.StatusInternalServerError)
		}

		return
	}

	// encode to json and serve
	if err := json.NewEncoder(w).Encode(reqObj); err != nil {
		config.CrudLog.Error(err, "error parsing object to json")
		http.Error(w, "error parsing object to json", http.StatusInternalServerError)
	}
}
