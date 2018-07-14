package main


import (
	_"management/routers"
    "management/models"
    "github.com/astaxie/beego"
	
    // "fmt"
     "github.com/astaxie/beego/orm"

   _ "github.com/go-sql-driver/mysql"
	
)
   

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	
    orm.Debug = true
	orm.RunSyncdb("default", false, true)
	
   beego.Run()
}




/*

package main

import (
    "encoding/base64"
    "fmt"
)

func base64Encode(src []byte) []byte {
    return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
    return base64.StdEncoding.DecodeString(string(src))
}

func main() {
    // encode
    hello := "你好，世界！ hello world"
    debyte := base64Encode([]byte(hello))
    fmt.Println(debyte)
    // decode
    enbyte, err := base64Decode(debyte)
    if err != nil {
        fmt.Println(err.Error())
    }

    if hello != string(enbyte) {
        fmt.Println("hello is not equal to enbyte")
    }

    fmt.Println(string(enbyte))
}



*/


















/*
    o := orm.NewOrm()
    user := models.User{Name: "slene"}
    id, err := o.Insert(&user)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)
 
    user1 := models.User{Name: "toom",Id:999}
    id, err3 := o.Insert(&user1)
    fmt.Printf("ID: %d, ERR: %v\n", id, err3)
 
    user2 := models.User{Id: 33}
    id, err4 := o.Insert(&user2)
    fmt.Printf("ID: %d, ERR: %v\n", id, err4)
 
    user3 :=models.User{Name: "mary"}
    id, err5 := o.Insert(&user3)
    fmt.Printf("ID: %d, ERR: %v\n", id, err5)
 
    user.Name = "astaxie"
    num, err6 := o.Update(&user)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err6)
 
    u := models.User{Id: user.Id}
    err1 := o.Read(&u)
    fmt.Printf("ERR: %v\n", err1)
        
    num, err2 := o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err2)
	*/
	

	



 
   /*   

 
 
 
 
   profile := new(models.Profile)
    profile.Id = 30 

    user := new(models.User)
    user.Profile = profile
    user.Name = "slene " 

    fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
	*/





































/*




package main

import (
	"fmt"
	_ "management/routers"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	
	//"database/sql"

)
type User struct{
	Id int	
	Nick string
/*	Password string
Email string
	Tel int
	
}
func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/management?charset=utf8")
	//orm.RegisterModel(new(models.Article))
}

func main() {
	/*
	o := orm.NewOrm()
    o.Using("default") // 默认使用 default，你可以指定为其他数据库

    profile := new(Profile)
    profile.Age = 30

    user := new(User)
    user.Profile = profile
    user.Name = "slene"

    fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
	
	


	o := orm.NewOrm()
user := User{Id: 1}

err := o.Read(&user)

if err == orm.ErrNoRows {
    fmt.Println("查询不到")
} else if err == orm.ErrMissPK {
    fmt.Println("找不到主键")
} else {
    fmt.Println(user.Id, user.Nick)
}
}

*/