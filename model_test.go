package main

import (
	"database/sql"
	"testing"
)

func TestModel(t *testing.T) {
	a := App{}
	dbinfo := "postgres://" + DB_USER + ":" + DB_PASSWORD + "@localhost/" + DB_NAME
	a.DB, _ = sql.Open("postgres", dbinfo)
	materials, _ := getMaterials(a.DB)
	if len(materials) < 0 {
		t.Errorf("Expected at least a result")
	}

}
