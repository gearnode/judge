package storage

// DB represents the database generic interface.
type DB interface {
	DescribeAll(table string) ([]interface{}, error)
	Describe(table string, id string) (interface{}, error)
	Put(table string, id string, object interface{}) error
	Delete(table string, id string) error
}
