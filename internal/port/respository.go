package port

// Repository represents the repository port
type Repository interface {
	Close() error
	Check() error
	Begin(base string) (interface{}, error)
	Commit(interface{}) error
	Rollback(interface{}) error
	InsertAuto(tx interface{}, base, object string, fields *[]string, vals *[]string) (int64, error)
	GetId(tx interface{}, base, object string, id int64, fields *[]string) (*[]string, error)
	GetField(tx interface{}, base, object, field, value string, fields *[]string) (*[]string, error)
}
