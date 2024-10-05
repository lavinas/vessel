package handler

// ClassCreateCmd represents the class create command
type ClassCreateCmd struct {
	Name        string `arg:"-n,--name" help:"The name of the class"`
	Description string `arg:"-d,--description" help:"The description of the class"`
}

// ClassGetCmd represents the class get command
type ClassGetCmd struct {
	ID int64 `arg:"-i,--id" help:"The id of the class"`
}
