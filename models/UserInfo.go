package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/Go-SQL-Driver/MySQL"
	"fmt"
	"reflect"
	"strconv"
)

type UserInfo struct {
	Id             int64   `orm:"column(id);pk;auto"           json:"id"             description:"主键ID号"`
	UserName       string  `orm:"column(user_name);size(100)"  json:"user_name"      description:"抖音用户名"`
	TotalRewards   float64 `orm:"column(total_rewards)"        json:"total_rewards"  description:"累计打赏总金额"`
	SettledAmounts float64 `orm:"column(settled_amounts)"      json:"settled_amounts"description:"当前可提现金额"`
	HeadLogo       string  `orm:"column(head_logo);size(100)"  json:"head_logo"      description:"用户头像"`
	CustomerId     string  `orm:"column(customer_id);size(100)"json:"customer_id"    description:"用户抖音ID号"`
	OpenId         string  `orm:"column(open_id);size(100)"    json:"open_id"        description:"用户在当前小程序的 ID号（即唯一标识ID）"`
}

func init() {
	orm.RegisterModel(new(UserInfo))
}
func (u *UserInfo) TableName() string {
	return "users"
}
func QueryUserInfo(open_id string) (UserInfo,error){
	o := orm.NewOrm()
	o.Using("default")
	var userInfo UserInfo
	err := o.Raw("select * from users where open_id = ?",open_id).QueryRow(&userInfo)
	if err != nil {
		return userInfo,err
	}
	return userInfo,nil
}
func InsertOrUpdateUserInfo0(user_info UserInfo)(string,error){
	var str string
	user_info.Id,_ = QueryNoUseId()
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil{
		return "insert or update fail",err
	}
	var maps []orm.Params
	num, err := o.Raw("select * from users where open_id = ?",user_info.OpenId).Values(&maps)
	if err == nil && num > 0{
		for _,item := range maps{
			user_info.Id,_ = strconv.ParseInt(fmt.Sprintf("%v",item["id"]), 10, 64)
		}
		
		_,err = o.Update(&user_info)
		if err != nil {
			err = o.Rollback()
			return "update fail",err
		}
		str = "update success"
	}else {
		_,err = o.Insert(&user_info)
		if err != nil {
			err = o.Rollback()
			return "insert fail",err
		}
		str = "insert success"
	}

	err = o.Commit()
	return str, nil

}
func InsertOrUpdateUserInfo1(user_info UserInfo)(string,error){
	var str string
	var isUpdate bool
	user_info.Id,_ = QueryNoUseId()
	userMaps := TraversalStruct(user_info)
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil{
		return "insert or update fail",err
	}

	var maps []orm.Params
	//num, err := o.Raw("select * from users where open_id = ?",user_info.OpenId).Values(&maps)

	num, err := o.Raw("select * from users where open_id = ?",user_info.OpenId).Values(&maps)

	//fmt.Println(num,err,user_info.OpenId)
	if err == nil && num > 0{
		for _,item := range maps{
			userMaps["id"] = item["id"]
			for i,v := range item{
				fmt.Printf("%s:%v-%v\n",i,v,userMaps[i])
				//fmt.Println(reflect.DeepEqual(v,userMaps[i]))
				//v1 := reflect.ValueOf(v)
				//v2 := reflect.ValueOf(userMaps[i])
				//fmt.Println(v1,v2,v1 != v2)
				v1 := fmt.Sprintf("%v",v)
				v2 := fmt.Sprintf("%v",userMaps[i])
				fmt.Println(v1 != v2)
				if v1 != v2 {
					//fmt.Println(v != userMaps[i])
					//这行很重要，缺少这行会使o.Update(user_info)即便执行成功也没效果
					user_info.Id,_ = strconv.ParseInt(fmt.Sprintf("%v",item["id"] ), 10, 64)
					_,err = o.Update(&user_info)
					if err != nil {
						err = o.Rollback()
						return "update fail",err
					}
					isUpdate = true
					str = "update success"
					break
				}
			}
		}
		//_,err = o.Update(&user_info)
		////fmt.Println(err)
		//if err != nil {
		//	err = o.Rollback()
		//	return "update fail",err
		//}
		if !isUpdate{
			str = "nothing update"
		}

	}else {
		_,err = o.Insert(&user_info)
		if err != nil {
			err = o.Rollback()
			return "insert fail",err
		}
		str = "insert success"
	}

	err = o.Commit()
	return str, nil
}
//遍历结构体字段
func TraversalStruct(obj interface{})map[string]interface{}{
	maps := make(map[string]interface{})
	//maps0 := new(map[string]interface{})
	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj) //获取reflect.Type类型

	//判断是不是结构体
	kd := val.Kind() //获取到obj对应的类别
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return nil
	}
	//获取到该结构体有几个字段
	num := val.NumField()
	fmt.Printf("该结构体有%d个字段\n", num) //4个
	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		//获取到struct标签，需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		fmt.Printf("Field %d:字段标签:%s值=%v\n",i,tagVal,val.Field(i))

		//如果该字段有tag标签就显示，否则就不显示
		//if tagVal != ""{
		//	//fmt.Printf("Field %d:tag=%v\n",i,tagVal)
		//}
		maps[tagVal] = val.Field(i)
	}
	return maps
}
//查询未占用的id
func QueryNoUseId()(int64,error){
	var id int64 = 1
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("select * from users").Values(&maps)
	if err != nil{
		return num,err
	}
	for _,item := range maps{
		idValue,_ := strconv.ParseInt(fmt.Sprintf("%v",item["id"]),10,64)
		//fmt.Println(id == idValue)
		if id == idValue{
			//fmt.Println("***")
			id ++
		}
	}

	return id,nil
}
