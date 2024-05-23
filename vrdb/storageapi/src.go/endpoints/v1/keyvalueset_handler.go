package v1

import (
	"encoding/json"
	"net/http"
	"vrdb-storage/objects"
	"vrdb-storage/server"

	"github.com/gorilla/mux"
	"vrdb.go/logging"
)

// list all results of the kvs table
//
// parameters will be handled in the future
func KVS_List(w http.ResponseWriter, r *http.Request) {

	var kvs_list []objects.KeyValueSet
	if result := server.Database.Find(&kvs_list); result.Error != nil {

		logging.Log.Info("error receiving list of kvs", "response-code", http.StatusInternalServerError, "kvs_list", kvs_list, "result.Error", result.Error)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	} else {

		if err := json.NewEncoder(w).Encode(kvs_list); err != nil {

			logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "err", err)

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

// find the first object by a given key
func KVS_Find(w http.ResponseWriter, r *http.Request) {

	// get the uid from the query params
	if key, ok := mux.Vars(r)["key"]; !ok {

		logging.Log.Info("key in query is missing", "response-code", http.StatusBadRequest, "query", mux.Vars(r))

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

	} else {

		var kvs objects.KeyValueSet

		if result := server.Database.First(&kvs, "key = ?", key); result.Error != nil {

			logging.Log.Info("error receiving item of kvs", "response-code", http.StatusInternalServerError, "result.Error", result.Error)

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		} else {

			if err := json.NewEncoder(w).Encode(kvs); err != nil {

				logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "kvs", kvs, "err", err)

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
	}
}

// insert a new kvs into the database
func KVS_Insert(w http.ResponseWriter, r *http.Request) {

	var (
		obj struct {
			Key         string `json:"key"`
			Value       string `json:"value"`
			Description string `json:"description"`
		}
	)

	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {

		logging.Log.Info("error parsing body into struct", "response-code", http.StatusBadRequest, "obj", obj, "err", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	// create object in database
	if result := server.Database.Create(&objects.KeyValueSet{
		Key:         obj.Key,
		Value:       obj.Value,
		Description: obj.Description}); result.Error != nil {

		logging.Log.Info("error creating object in database", "response-code", http.StatusInternalServerError, "obj", obj, "result.Error", result.Error)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	// get the object id
	var kvs = objects.KeyValueSet{}
	if result := server.Database.First(&kvs, "key = ?", obj.Key); result.Error != nil {

		logging.Log.Info("error finding kvs", "response-code", http.StatusInternalServerError, "kvs", kvs, "result.Error", result.Error)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	// send result
	if err := json.NewEncoder(w).Encode(kvs); err != nil {

		logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "kvs", kvs, "err", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func KVS_Update(w http.ResponseWriter, r *http.Request) {

	var (
		obj struct {
			Key         string `json:"key"`
			Value       string `json:"value"`
			Description string `json:"description"`
		}
	)

	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {

		logging.Log.Info("error parsing body into struct", "response-code", http.StatusBadRequest, "obj", obj, "err", err)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	// get the object id
	var kvs = objects.KeyValueSet{}
	if result := server.Database.First(&kvs, "key = ?", obj.Key); result.Error != nil {

		logging.Log.Info("error finding kvs", "response-code", http.StatusInternalServerError, "kvs", kvs, "result.Error", result.Error)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	kvs.Value = obj.Value
	kvs.Description = obj.Description

	// create object in database
	if result := server.Database.Save(&kvs); result.Error != nil {

		logging.Log.Info("error creating object in database", "response-code", http.StatusInternalServerError, "obj", obj, "result.Error", result.Error)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	// send result
	if err := json.NewEncoder(w).Encode(kvs); err != nil {

		logging.Log.Info("error parsing result into json", "response-code", http.StatusInternalServerError, "kvs", kvs, "err", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func KVS_Remove(w http.ResponseWriter, r *http.Request) {

	// get the uid from the query params
	if key, ok := mux.Vars(r)["key"]; !ok {

		logging.Log.Info("key in query is missing", "response-code", http.StatusBadRequest, "query", mux.Vars(r))

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

	} else {

		// get the object id
		var kvs = objects.KeyValueSet{}
		if result := server.Database.First(&kvs, "key = ?", key); result.Error != nil {

			logging.Log.Info("error finding kvs", "response-code", http.StatusInternalServerError, "kvs", kvs, "result.Error", result.Error)

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}

		// remove the object
		if result := server.Database.Delete(&kvs); result.Error != nil {

			logging.Log.Info("error removing kvs", "response-code", http.StatusInternalServerError, "kvs", kvs, "result.Error", result.Error)

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}
	}

	w.Write([]byte("Deleted"))
}
