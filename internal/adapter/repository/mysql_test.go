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
	vals := map[string]*string{
		"id": &[]string{"99999999"}[0],
		"name": &[]string{"test"}[0],
		"description": nil,
		"created_at": &[]string{"2021-01-01"}[0],
	}
	id, err := mysql.InsertAuto(tx, "assets", "class", &vals)
	if err != nil {
		t.Error(err)
	}
	if id != 99999999 {
		t.Error("id != 99999999")
	}
	err = mysql.DeleteId(tx, "assets", "class", id)
	if err != nil {
		t.Error(err)
	}
	err = mysql.Commit(tx)
	if err != nil {
		t.Error(err)
	}	
}