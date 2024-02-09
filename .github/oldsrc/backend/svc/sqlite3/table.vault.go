package sqlite3

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	_log = ctrl.Log.WithName("database")
)

type Pair struct {
	UID   string `json:"uid"`
	Value string `json:"value"`
}

// get all vaultpairs from the database
func SelectAllPairs() ([]Pair, error) {
	if rows, err := _DBConn.Query("SELECT uid, data FROM vault;"); err != nil {
		_log.Error(err, "error receiving all pairs", "function", "SelectAllPairs")
		return []Pair{}, err
	} else {
		defer rows.Close()
		var res []Pair
		var uid, data string
		for rows.Next() {
			if err := rows.Scan(&uid, &data); err != nil {
				_log.Error(err, "error scanning the current row", "function", "SelectAllPairs", "uid", uid)
				return []Pair{}, err
			}
			res = append(res, Pair{UID: uid, Value: data})
		}
		return res, nil
	}
}

// select a specific pair by the uid
func SelectPairByUID(uid string) (Pair, error) {
	var res Pair
	if err := _DBConn.QueryRow("SELECT uid, data FROM vault WHERE uid=$1;", uid).Scan(&res.UID, &res.Value); err != nil {
		_log.Error(err, "error receiving specific pair", "function", "SelectPairByUID", "uid", uid)
		return Pair{}, err
	}
	return res, nil
}

// insert a new pair with specific data
func InsertPair(p Pair) (Pair, error) {
	if _, err := _DBConn.Exec("INSERT INTO vault (uid, data) VALUES ($1, $2);", p.UID, p.Value); err != nil {
		_log.Error(err, "error inserting a new pair", "function", "InsertPair", "uid", p.UID, "data.length", len(p.Value))
		return Pair{}, err
	}
	return p, nil
}

// update the data from a specific pair
func UpdatePair(p Pair) (Pair, error) {
	if _, err := _DBConn.Exec("UPDATE vault SET data=$1 WHERE uid=$2;", p.Value, p.UID); err != nil {
		_log.Error(err, "error updating an existing pair", "function", "UpdatePair", "uid", p.UID, "data.length", len(p.Value))
		return Pair{}, err
	}
	return p, nil
}

// remove an existing pair by uid
func DeletePair(uid string) error {
	var exists bool = false
	if err := _DBConn.QueryRow("SELECT exists(SELECT 1 FROM vault WHERE uid=$1);", uid).Scan(&exists); err != nil {
		_log.Error(err, "error checking exisitence of a pair", "function", "DeletePair", "uid", uid)
		return err
	}
	if exists {
		if _, err := _DBConn.Exec("DELETE FROM vault WHERE uid=$1;", uid); err != nil {
			_log.Error(err, "error deleting a pair", "function", "DeletePair", "uid", uid)
			return err
		}
	}
	return nil
}
