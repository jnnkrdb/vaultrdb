package vaultrequest

import (
	"github.com/go-logr/logr"
	jnnkrdbdev1 "github.com/jnnkrdb/vaultrdb/api/v1"
)

// validate the data, provided by the dataref fields. Returns a map, usable for
// secrets/configmaps data, a [valid] bool value, indicating whether the checks were
// valid or not and an error value
func receiveDataMap(_log logr.Logger, mp map[string]jnnkrdbdev1.DataMap) (map[string]string, bool, error) {

	// if the datamap [mp] is empty, then there is no error while validation,
	// but the validation check will be false
	if length := len(mp); length <= 0 {
		_log.Info("empty datamaps are not allowed", "len(datamap)", length)
		return nil, false, nil
	}

	var objectData = make(map[string]string)

	// start processing the datamap fields
	for dmKey, dmValue := range mp {
		_log = _log.WithValues("datamapkey", dmKey)

		// get the data for the key
		if dat, e := dmValue.GetData(_log); e != nil {
			_log.Error(e, "invalid datamap object")
			return nil, false, e
		} else {
			objectData[dmKey] = dat
		}
	}

	return objectData, true, nil
}
