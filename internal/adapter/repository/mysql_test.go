package repository

import (
	"reflect"
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
		"name":       "test3",
		"created_at": time.Now(),
		"value":      10.2,
		"value2":     100,
	}
	id, err := mysql.Insert(tx, "assets", "test", &vals)
	if err != nil {
		t.Error(err)
	}
	if id == 0 {
		t.Error("id generated is 0")
	}
	vals = map[string]interface{}{
		"id": id,
	}
	row, err := mysql.Get(tx, "assets", "test", &vals)
	if err != nil {
		t.Error(err)
	}
	if row == nil {
		t.Fatal("row == nil")
	}
	t.Error(1, reflect.TypeOf((*row)[0]["id"]).String())
	t.Error(2, reflect.TypeOf((*row)[0]["name"]).String())
	t.Error(3, reflect.TypeOf((*row)[0]["value"]).String())
	t.Error(4, reflect.TypeOf((*row)[0]["value2"]).String())
	t.Error(5, reflect.TypeOf(5.2).String())


	/*
	if (*row)[0]["id"].(int64) != id {
		t.Error("row[0][\"id\"] != id")
	}
	if (*row)[0]["name"].(string) != "test" {
		t.Error("row[0][\"name\"] != \"test\"")
	}
	if (*row)[0]["description"].(string) != "desc" {
		t.Error("row[0][\"description\"] != \"desc\"")
	}
	if (*row)[0]["created_at"].(time.Time).IsZero() {
		t.Error("row[0][\"created_at\"].IsZero()")
	}
	*/
	err = mysql.DeleteId(tx, "assets", "test", id)
	if err != nil {
		t.Error(err)
	}

	err = mysql.Commit(tx)
	if err != nil {
		t.Error(err)
	}
}
