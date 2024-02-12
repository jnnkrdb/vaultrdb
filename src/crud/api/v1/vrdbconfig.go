package v1

import (
	"net/http"
)

type VRDBConfigCRUD struct{}

func (rw *VRDBConfigCRUD) Create(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "creating not implemented", http.StatusNotImplemented)
}

func (rw *VRDBConfigCRUD) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "deletion not implemented", http.StatusNotImplemented)
}

func (rw *VRDBConfigCRUD) Update(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "updating not implemented", http.StatusNotImplemented)
}

func (rw *VRDBConfigCRUD) Read(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "reading not implemented", http.StatusNotImplemented)
}
