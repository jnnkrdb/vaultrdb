package v1

// remove a specific object from the
func Remove(list []DeployedObject, kind, namespace string) []DeployedObject {

	if len(list) == 0 {
		return []DeployedObject{}
	}

	var result []DeployedObject
	for i := range list {

		if list[i].Kind == kind && list[i].Namespace == namespace {
			continue
		}

		result = append(result, list[i])
	}

	return result
}

// append a new status object
func Append(list []DeployedObject, kind, namespace string) []DeployedObject {
	if kind == "" || namespace == "" {
		return list
	}
	return append(list, DeployedObject{Kind: kind, Namespace: namespace})
}

// check if a specific item is in the status object
func Contains(list []DeployedObject, kind, namespace string) bool {
	if len(list) == 0 {
		return false
	}

	for i := range list {
		if list[i].Kind == kind && list[i].Namespace == namespace {
			return true
		}
	}

	return false
}
