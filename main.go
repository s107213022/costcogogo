package main

import (
	_ "costcogogo/routers"
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var DBARGS = struct {
	Driver string
	Source string
	Debug  string
}{
	os.Getenv("ORM_DRIVER"),
	os.Getenv("ORM_SOURCE"),
	os.Getenv("ORM_DEBUG"),
}

func init() {
	Debug, _ := orm.StrTo(DBARGS.Debug).Bool()

	fmt.Printf("Debug=%v\n", Debug)
	if Debug {
		fmt.Printf("ORM config=%v\n", DBARGS)
	}
	if DBARGS.Driver == "" || DBARGS.Source == "" {
		fmt.Println("Please set ORM_DRIVER/ORM_SOURCE")
		os.Exit(2)
	}

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:1qaz2wsx@tcp(127.0.0.1)/costcogogo?charset=utf8", 30)
	orm.RunSyncdb("default", false, true)
}

func main() {
	// userNotAdded := models.User{Name: "John Doe", Account: "123", Password: "456"}
	// userid, err := models.AddUser(&userNotAdded)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(userid)
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
