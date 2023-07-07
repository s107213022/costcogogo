package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Money_20230706_145232 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Money_20230706_145232{}
	m.Created = "20230706_145232"

	migration.Register("Money_20230706_145232", m)
}

// Run the migrations
func (m *Money_20230706_145232) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE money(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`time` datetime NOT NULL,`items` varchar(128) NOT NULL,`price` int(11) DEFAULT NULL,`content` varchar(128) NOT NULL,`owe` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Money_20230706_145232) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `money`")
}
