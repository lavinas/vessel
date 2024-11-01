package entity

// User is a struct that represents a user
type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	IsEmployed bool   `json:"isEmployed"`
}

// TableName is a method that returns the table name
func (u *User) TableName() string {
	return "user"
}

// RestPath is a method that returns the REST path
func (u *User) RestPath() string {
	return "user"
}
