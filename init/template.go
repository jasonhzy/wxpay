package init

import (
	"time"
	"github.com/astaxie/beego"
)

//时间转换
func convertT(num float64) (out string){
	tm := time.Unix(int64(num / 1000), 0)
	out = tm.Format("2006-01-02 03:04:05")
	return
}

//自定义模板函数
func init() {
	beego.AddFuncMap("convertT", convertT) //时间转换
}