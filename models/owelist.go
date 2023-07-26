package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Owelist struct {
	Id       int64  `orm:"auto"`
	Creditor *User  `orm:"rel(fk)"`
	Items    string `orm:"size(128)"`
	Money    int
	Unit     int
	Date     time.Time `orm:"type(datetime)"`
	Finish   int
	Debtor   *User `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Owelist))
}

// AddOwelist insert a new Owelist into database and returns
// last inserted Id on success.
func AddoweList(m *Owelist) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOwelistById retrieves Owelist by Id. Returns error if
// Id doesn't exist
func GetOwelistById(id int64) (*Owelist, error) {
	o := orm.NewOrm()
	owelist := &Owelist{Id: id}
	err := o.QueryTable(new(Owelist)).Filter("Id", id).RelatedSel("Creditor", "Debtor").One(owelist)
	if err != nil {
		return nil, err
	}
	return owelist, nil
}

func GetOwelistByCreditorID(creditorID int64) (oweList []*Owelist, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Owelist))

	// 過濾Creditor的ID
	qs = qs.Filter("Creditor", creditorID)

	// 載入相關的Creditor和Debtor模型以獲取它們的名稱
	qs = qs.RelatedSel("Creditor", "Debtor")

	// 獲取所有符合條件的Owelist記錄
	_, err = qs.All(&oweList)
	if err != nil {
		return nil, err
	}
	return oweList, nil
}

func UpdateOwelist(owelist *Owelist) error {
	o := orm.NewOrm()
	_, err := o.Update(owelist)
	return err
}

// DeleteOwelist deletes Owelist by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOwelist(id int64) (err error) {
	o := orm.NewOrm()
	v := Owelist{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Owelist{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
