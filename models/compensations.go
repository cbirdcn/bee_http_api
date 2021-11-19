package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Compensations struct {
	Id         int       `orm:"column(id);auto"`
	ServerId   int       `orm:"column(server_id)"`
	Title      string    `orm:"column(title);size(255)"`
	Content    string    `orm:"column(content);size(5000)"`
	Note       string    `orm:"column(note);size(255)" description:"面向使用者的备忘录"`
	ExecutedAt time.Time `orm:"column(executed_at);type(datetime);null"`
	CreatedAt  time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(datetime);auto_now"`
}

func (t *Compensations) TableName() string {
	return "compensations"
}

func init() {
	orm.RegisterModel(new(Compensations))
}

// AddCompensations insert a new Compensations into database and returns
// last inserted Id on success.
func AddCompensations(m *Compensations) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCompensationsById retrieves Compensations by Id. Returns error if
// Id doesn't exist
func GetCompensationsById(id int) (v *Compensations, err error) {
	o := orm.NewOrm()
	v = &Compensations{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCompensations retrieves all Compensations matches certain condition. Returns empty list if
// no records exist
func GetAllCompensations(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Compensations))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
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

	var l []Compensations
	qs = qs.OrderBy(sortFields...)
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

// UpdateCompensations updates Compensations by Id and returns error if
// the record to be updated doesn't exist
func UpdateCompensationsById(m *Compensations) (err error) {
	o := orm.NewOrm()
	v := Compensations{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCompensations deletes Compensations by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCompensations(id int) (err error) {
	o := orm.NewOrm()
	v := Compensations{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Compensations{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
