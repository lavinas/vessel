package repository

import (
	"os"
	"time"

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

func TestInsertGet(t *testing.T) {
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
	vals := map[string]interface{}{
		"name":        "test",
		"description": "desc",
		"created_at":  time.Now(),
	}
	id, err := mysql.Insert(tx, "assets", "class", &vals)
	if err != nil {
		t.Error(err)
	}
	if id == 0 {
		t.Error("id generated is 0")
	}
	vals = map[string]interface{}{
		"id": id,
	}
	row, err := mysql.Get(tx, "assets", "class", &vals)
	if err != nil {
		t.Error(err)
	}
	if row == nil {
		t.Fatal("row == nil")
	}
	if (*row)[0]["id"].(int64) != id {
		t.Error("row[0][\"id\"] != id")
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
