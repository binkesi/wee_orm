package schema

import (
	"hooks/dialect"
	"log"
	"testing"
)

type User struct {
	Name string `weeorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		log.Fatal("Parse User struct failed")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		log.Fatal("Parse Name primary key failed")
	}
}
