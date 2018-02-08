package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
)

//用户表
type User struct {
	Id            int64
	Username      string
	Password      string
	Repasswd  	  string    `orm:"-"`
	Nickname      string
	Age			  int
	Sex			  int
	Create_time   time.Time
	Update_time   time.Time
}

func init(){
	//表前缀
	tb_prefix := beego.AppConfig.String("tb_prefix")
	orm.RegisterModelWithPrefix(tb_prefix, new(User))
}

func GetList(page int64, page_size int64, sort string) (users []orm.Params, count int64){
	o := orm.NewOrm()

	//var user []User
	//o.Raw("select * from user").QueryRows(&user)
	//fmt.Printf("===%s", user[0].Username)
	//var user User
	//o.Raw("select * from user").QueryRow(&user)
	//fmt.Printf("===%s", user.Username)

	qs := o.QueryTable(new(User))
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&users)
	count, _ = qs.Count()
	return users, count
}

func GetUserById(user User) error {
	o := orm.NewOrm()
	return o.Read(&user)
}

func Add(user User) (int64, error){
	o := orm.NewOrm()
	//在传入对象时，只能转入对象的指针，这样才能保证通过反射获取相关数据是正确合理的。
	return o.Insert(&user)
}

func UpdateBycolumn(user User, column string) (int64, error){
	o := orm.NewOrm()
	return o.Update(&user, column)
}
