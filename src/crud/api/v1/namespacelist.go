package v1

import (
	"net/http"
)

type NamespaceListCRUD struct{}

func (rw *NamespaceListCRUD) Create(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "creating not implemented", http.StatusNotImplemented)
}

func (rw *NamespaceListCRUD) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "deletion not implemented", http.StatusNotImplemented)
}

func (rw *NamespaceListCRUD) Update(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "updating not implemented", http.StatusNotImplemented)
}

func (rw *NamespaceListCRUD) Read(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "reading not implemented", http.StatusNotImplemented)
}
