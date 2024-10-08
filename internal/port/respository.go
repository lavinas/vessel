package port

// Repository represents the repository port
type Repository interface {
	Begin(base string) (interface{}, error)
	InsertAuto(tx interface{}, base, object string, fields *[]string, vals *[]string) (int64, error)
	GetId(tx interface{}, base, object string, id int64, fields *[]string) (*[]interface{}, error)
	GetField(tx interface{}, base, object, field, value string, fields *[]string) (*[]interface{}, error)
	Commit(interface{}) error
	Rollback(interface{}) error
	Close() error
}
