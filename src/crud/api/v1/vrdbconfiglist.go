package v1

import (
	"net/http"
)

type VRDBConfigListCRUD struct{}

func (rw *VRDBConfigListCRUD) Create(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "creating not implemented", http.StatusNotImplemented)
}

func (rw *VRDBConfigListCRUD) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "deletion not implemented", http.StatusNotImplemented)
}

func (rw *VRDBConfigListCRUD) Update(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "updating not implemented", http.StatusNotImplemented)
}

func (rw *VRDBConfigListCRUD) Read(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "reading not implemented", http.StatusNotImplemented)
}
