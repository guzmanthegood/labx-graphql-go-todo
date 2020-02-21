package model

// dataStore global variable
var ds DataStore

// set global variable
func SetDataStore(dataStore DataStore) {
	ds = dataStore
}

// get unique client
func GetDataStore() DataStore {
	return ds
}
