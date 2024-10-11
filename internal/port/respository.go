package port

// Repository represents the repository port
type Repository interface {
	Begin(base string) (interface{}, error)
	Insert(tx interface{}, base, object string, vals *map[string]interface{}) (int64, error)
	Get(tx interface{}, base, object string, vals *map[string]interface{}) (*[]map[string]interface{}, error)
	Commit(interface{}) error
	Rollback(interface{}) error
	Close() error
}
