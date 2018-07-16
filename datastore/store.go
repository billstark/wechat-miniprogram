package datastore

// Store is an interface that wraps most common methods of a datastore,
// which are create, retrieve, update and delete
type Store interface {

	// Creates a record. all arguments are included in args
	Create(args interface{}) (interface{}, error)

	// Retrieves record(s).
	Retrieve(args interface{}) (interface{}, error)

	// Updates a record. id should be specified
	Update(id string, args interface{}) (interface{}, error)

	// Deletes a record. id should be specified
	Delete(id string, args interface{}) (interface{}, error)
}
