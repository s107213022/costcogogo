package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Owelist_20230724_124048 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Owelist_20230724_124048{}
	m.Created = "20230724_124048"

	migration.Register("Owelist_20230724_124048", m)
}

// Run the migrations
func (m *Owelist_20230724_124048) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE owelist(`id` int(11) NOT NULL AUTO_INCREMENT,`creditor` int(11) DEFAULT NULL,`items` varchar(128) NOT NULL,`money` int(11) DEFAULT NULL,`unit` int(11) DEFAULT NULL,`date` datetime NOT NULL,`finish` int(11) DEFAULT NULL,`debtor` int(11) DEFAULT NULL,PRIMARY KEY (`id`),FOREIGN KEY(creditor) REFERENCES user(id),FOREIGN KEY(debtor) REFERENCES user(id))")
}

// Reverse the migrations
func (m *Owelist_20230724_124048) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `owelist`")
}
