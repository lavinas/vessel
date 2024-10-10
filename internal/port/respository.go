package port

// Repository represents the repository port
type Repository interface {
	Begin(base string) (interface{}, error)
	InsertAuto(tx interface{}, base, object string, fields *[]string, vals *[]string) (int64, error)
	Get(tx interface{}, base, object, key, value string) (*map[string]int, *[][]*string, error)
	Commit(interface{}) error
	Rollback(interface{}) error
	Close() error
}
