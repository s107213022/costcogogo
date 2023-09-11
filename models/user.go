package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64  `orm:"auto"`
	Name     string `orm:"size(128)"`
	Account  string `orm:"size(128)"`
	Password string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func CheckLogin(account, password string) (loggedIn bool, userID int64, err error) {
	o := orm.NewOrm()
	user := User{Account: account, Password: password}
	err = o.Read(&user, "Account", "Password")
	if err == nil {
		loggedIn = true
		userID = user.Id
		return loggedIn, userID, nil
	}
	if err == orm.ErrNoRows {
		return false, 0, nil
	}
	return false, 0, err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int64) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetNameById(id int64) (name string, err error) {
	o := orm.NewOrm()
	user := User{Id: id}
	err = o.Read(&user)
	if err == nil {
		name = user.Name
		return name, nil
	}
	return "", err
}

func GetUserById(id int64) (user *User, err error) {
	o := orm.NewOrm()
	user = &User{Id: id}
	err = o.Read(user)
	if err == nil {
		return user, nil
	}
	return nil, err
}

func GetUserByName(name string) (user *User, err error) {
	o := orm.NewOrm()
	user = &User{Name: name}
	err = o.Read(user, "Name")
	if err == nil {
		return user, nil
	}
	return nil, err
}

// GetUserExceptSelf 除了自己之外的所有人
func GetNamesExceptSelf(selfID int64) (names []string, err error) {
	o := orm.NewOrm()
	var users []*User
	_, err = o.QueryTable(new(User)).
		Exclude("Id", selfID).
		All(&users)
	if err == nil {
		for _, user := range users {
			names = append(names, user.Name)
		}
		return names, nil
	}
	return nil, err
}
