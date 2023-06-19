package v1

import (
	"fmt"
	"strings"
)

type DeployedObjects []string

func fromNamespaceKind(kind, namespace string) string {
	return fmt.Sprintf("%s/%s", kind, namespace)
}

func (do DeployedObjects) GetKindAndNamespace(namespacekind string) (string, string) {
	if len(namespacekind) == 0 {
		return "", ""
	}
	var splitted = strings.Split(namespacekind, "/")
	return splitted[0], splitted[1]
}

// remove a specific object from the
func (do DeployedObjects) RemoveObject(kind, namespace string) DeployedObjects {

	if len(do) == 0 {
		return []string{}
	}

	var result DeployedObjects
	for i := range do {
		if do[i] != fromNamespaceKind(kind, namespace) {
			result = append(result, do[i])
		}
	}

	return result
}

// append a new status object
func (do DeployedObjects) Append(kind, namespace string) DeployedObjects {
	return append(do, fromNamespaceKind(kind, namespace))
}

// check if a specific item is in the status object
func (do DeployedObjects) Contains(kind, namespace string) bool {
	if len(do) == 0 {
		return false
	}

	for i := range do {
		if do[i] == fromNamespaceKind(kind, namespace) {
			return true
		}
	}

	return false
}
