package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Money_20230706_143704 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Money_20230706_143704{}
	m.Created = "20230706_143704"

	migration.Register("Money_20230706_143704", m)
}

// Run the migrations
func (m *Money_20230706_143704) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Money_20230706_143704) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
