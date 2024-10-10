package repository

import (
	"os"

	"testing"
)

func TestNewMySql(t *testing.T) {
	mysql, err := NewMySql(os.Getenv("MYSQL_DNS"), "")
	if err != nil {
		t.Error(err)
	}
	err = mysql.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestInsertAuto(t *testing.T) {
	mysql, err := NewMySql(os.Getenv("MYSQL_DNS"), "")
	if err != nil {
		t.Error(err)
	}
	defer mysql.Close()
	tx, err := mysql.Begin("assets")
	if err != nil {
		t.Error(err)
	}
	defer mysql.Rollback(tx)
	_, err = mysql.InsertAuto(tx, "assets", "class", &[]string{"test"}, &[]string{"test"})
	if err != nil {
		t.Error(err)
	}
}