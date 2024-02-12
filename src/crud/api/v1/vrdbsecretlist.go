package v1

import (
	"net/http"
)

type VRDBSecretListCRUD struct{}

func (rw *VRDBSecretListCRUD) Create(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "creating not implemented", http.StatusNotImplemented)
}

func (rw *VRDBSecretListCRUD) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "deletion not implemented", http.StatusNotImplemented)
}

func (rw *VRDBSecretListCRUD) Update(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "updating not implemented", http.StatusNotImplemented)
}

func (rw *VRDBSecretListCRUD) Read(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "reading not implemented", http.StatusNotImplemented)
}
