package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
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
func GetOwelistById(id int64) (v *Owelist, err error) {
	o := orm.NewOrm()
	v = &Owelist{Id: id}
	if err = o.QueryTable(new(Owelist)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
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

// GetAllOwelist retrieves all Owelist matches certain condition. Returns empty list if
// no records exist
func GetAllOwelist(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Owelist))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Owelist
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateOwelist updates Owelist by Id and returns error if
// the record to be updated doesn't exist
func UpdateOwelistById(m *Owelist) (err error) {
	o := orm.NewOrm()
	v := Owelist{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
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
