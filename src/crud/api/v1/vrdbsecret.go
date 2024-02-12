package v1

import (
	"net/http"
)

type VRDBSecretCRUD struct{}

func (rw *VRDBSecretCRUD) Create(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "creating not implemented", http.StatusNotImplemented)
}

func (rw *VRDBSecretCRUD) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "deletion not implemented", http.StatusNotImplemented)
}

func (rw *VRDBSecretCRUD) Update(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "updating not implemented", http.StatusNotImplemented)
}

func (rw *VRDBSecretCRUD) Read(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "reading not implemented", http.StatusNotImplemented)
}
