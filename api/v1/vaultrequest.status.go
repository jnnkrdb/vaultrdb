package v1

type DeployedObject struct {
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
}

type deployedObjectList []DeployedObject

// remove a specific object from the
func (dol deployedObjectList) RemoveObject(kind, namespace string) deployedObjectList {

	if len(dol) == 0 {
		return deployedObjectList{}
	}

	var result []DeployedObject
	for i := range dol {

		if dol[i].Kind == kind && dol[i].Namespace == namespace {
			continue
		}

		result = append(result, dol[i])
	}

	return result
}

// append a new status object
func (do deployedObjectList) Append(kind, namespace string) deployedObjectList {
	if kind == "" || namespace == "" {
		return do
	}
	return append(do, DeployedObject{Kind: kind, Namespace: namespace})
}

// check if a specific item is in the status object
func (do deployedObjectList) Contains(kind, namespace string) bool {
	if len(do) == 0 {
		return false
	}

	for i := range do {
		if do[i].Kind == kind && do[i].Namespace == namespace {
			return true
		}
	}

	return false
}
