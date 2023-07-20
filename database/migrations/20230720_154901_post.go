package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Post_20230720_154901 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Post_20230720_154901{}
	m.Created = "20230720_154901"

	migration.Register("Post_20230720_154901", m)
}

// Run the migrations
func (m *Post_20230720_154901) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE post(`id` int(11) NOT NULL AUTO_INCREMENT,`creditor` int(11) DEFAULT NULL,`items` varchar(128) NOT NULL,`unit` int(11) DEFAULT NULL,`date` datetime NOT NULL,`finish` int(11) DEFAULT NULL,`debtor` int(11) DEFAULT NULL,PRIMARY KEY (`id`),FOREIGN KEY (`creditor`) REFERENCES `user` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,FOREIGN KEY (`debtor`) REFERENCES `user` (`id`) ON DELETE SET NULL ON UPDATE CASCADE)")
}

// Reverse the migrations
func (m *Post_20230720_154901) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `post`")
}
