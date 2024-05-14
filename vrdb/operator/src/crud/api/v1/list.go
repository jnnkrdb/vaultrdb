package v1

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
	"github.com/jnnkrdb/vaultrdb/crud/config"
	metav1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func LIST_ALL(w http.ResponseWriter, r *http.Request) {
	list(r.Context(), false, w, r)
}

func LIST_NAMESPACE(w http.ResponseWriter, r *http.Request) {
	list(r.Context(), true, w, r)
}

func list(ctx context.Context, namespace bool, w http.ResponseWriter, r *http.Request) {

	// register the rest variables
	vars := mux.Vars(r)
	config.CrudLog.Info("got vars", "vars", vars)

	// declare a new client.Object for the apiserver request
	var reqObj client.ObjectList

	switch vars["kind"] {
	case "vrdbconfigs":
		reqObj = &jnnkrdbdev1.VRDBConfigList{}
	case "vrdbsecrets":
		reqObj = &jnnkrdbdev1.VRDBSecretList{}
	case "vrdbrequests":
		http.Error(w, "type not implemented", http.StatusNotImplemented)
		return
		//reqObj = &jnnkrdbdev1.VRDBRequestList{}
	case "namespaces":
		reqObj = &metav1.NamespaceList{}
	default:
		http.Error(w, "type not implemented", http.StatusNotImplemented)
		return
	}

	var clientListOptions = &client.ListOptions{}

	// if namespace is requested, then add to clientListOptions
	if namespace {
		clientListOptions.Namespace = vars["namespace"]
	}

	// requesting the list of objects
	if err := config.KClient.List(ctx, reqObj, clientListOptions); err != nil {
		config.CrudLog.Error(err, "error receiving list of objects",
			"object.kind", reqObj.GetObjectKind().GroupVersionKind().Kind,
			"object.group", reqObj.GetObjectKind().GroupVersionKind().Group,
			"object.version", reqObj.GetObjectKind().GroupVersionKind().Version,
		)
		http.Error(w, "error receiving list of objects", http.StatusInternalServerError)
		return
	}

	// encode to json and serve
	if err := json.NewEncoder(w).Encode(reqObj); err != nil {
		config.CrudLog.Error(err, "error parsing list to json")
		http.Error(w, "error receiving list of objects", http.StatusInternalServerError)
	}
}
