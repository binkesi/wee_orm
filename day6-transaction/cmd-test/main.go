package main

import (
	"fmt"
	weeorm "transaction"
	"transaction/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := weeorm.NewEngine("sqlite3", "wee.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXIST User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Sungn", "Wumx").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
	row := s.Raw("SELECT Name FROM User LIMIT 1;").QueryRow()
	var name string
	row.Scan(&name)
	log.Info(name)
}
