package v1

import (
	"net/http"
	/*
		"encoding/json"

		"github.com/gorilla/mux"
		jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
		"github.com/jnnkrdb/vaultrdb/crud/config"
		"k8s.io/apimachinery/pkg/api/errors"
		v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
		"k8s.io/apimachinery/pkg/types"
		"sigs.k8s.io/controller-runtime/pkg/client"
	*/)

func CREATE(w http.ResponseWriter, r *http.Request) {

	// ------------------------------------------------
	// currently this function is deactivated
	http.Error(w, "not implemented", http.StatusNotImplemented)
	// ------------------------------------------------

	/*

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

		// check if the object already exists
		var alreadyExists bool = true
		if err := config.KClient.Get(r.Context(), types.NamespacedName{
			Namespace: vars["namespace"],
			Name:      vars["name"],
		}, reqObj, &client.GetOptions{}); err != nil {

			if errors.IsNotFound(err) {
				alreadyExists = false
			} else {
				config.CrudLog.Error(err, "error receiving object",
					"object.kind", reqObj.GetObjectKind().GroupVersionKind().Kind,
					"object.group", reqObj.GetObjectKind().GroupVersionKind().Group,
					"object.version", reqObj.GetObjectKind().GroupVersionKind().Version,
				)
				http.Error(w, "error receiving object", http.StatusInternalServerError)
				return
			}
		}

		// if the object does not exist, then validate the object,
		// that was send in the request body
		if !alreadyExists {
			if err := json.NewDecoder(r.Body).Decode(reqObj); err != nil {
				config.CrudLog.Error(err, "error decoding object",
					"object.kind", reqObj.GetObjectKind().GroupVersionKind().Kind,
					"object.group", reqObj.GetObjectKind().GroupVersionKind().Group,
					"object.version", reqObj.GetObjectKind().GroupVersionKind().Version,
					"r.Body", reqObj,
				)
				http.Error(w, "error decoding object", http.StatusInternalServerError)
				return
			}

			// use the new object to apply a new object to the cluster
			var new client.Object
			switch vars["kind"] {
			case "vrdbconfigs":
				if incoming, ok := reqObj.(*jnnkrdbdev1.VRDBConfig); ok {
					new = &jnnkrdbdev1.VRDBConfig{
						ObjectMeta: v1.ObjectMeta{
							Name:        vars["name"],
							Namespace:   vars["namespace"],
							Annotations: incoming.Annotations,
							Labels:      incoming.Labels,
						},
						NamespaceSelector: incoming.NamespaceSelector,
						Data:              incoming.Data,
					}
				}
			case "vrdbsecrets":
				if incoming, ok := reqObj.(*jnnkrdbdev1.VRDBSecret); ok {
					new = &jnnkrdbdev1.VRDBSecret{
						ObjectMeta: v1.ObjectMeta{
							Name:        vars["name"],
							Namespace:   vars["namespace"],
							Annotations: incoming.Annotations,
							Labels:      incoming.Labels,
						},
						NamespaceSelector: incoming.NamespaceSelector,
						Data:              incoming.Data,
						StringData:        incoming.StringData,
						Type:              incoming.Type,
					}
				}
			case "vrdbrequests":
				if incoming, ok := reqObj.(*jnnkrdbdev1.VRDBRequest); ok {
					new = &jnnkrdbdev1.VRDBConfig{
						ObjectMeta: v1.ObjectMeta{
							Name:        vars["name"],
							Namespace:   vars["namespace"],
							Annotations: incoming.Annotations,
							Labels:      incoming.Labels,
						},
						NamespaceSelector: incoming.NamespaceSelector,
						Data:              incoming.Data,
					}
				}
			}

			// create the cached object and send the updated version as a response
			if err := config.KClient.Create(r.Context(), new, &client.CreateOptions{}); err != nil {
				config.CrudLog.Error(err, "error creating object")
				http.Error(w, "error creating object", http.StatusInternalServerError)
				return
			}

			// update the cached object and send as a response
			if err := config.KClient.Get(r.Context(), types.NamespacedName{
				Name:      vars["name"],
				Namespace: vars["namespace"],
			}, reqObj, &client.GetOptions{}); err != nil {
				if errors.IsNotFound(err) {
					http.Error(w, "error object not found", http.StatusNotFound)
				} else {
					config.CrudLog.Error(err, "error creating object")
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
	*/
}
