//package main
//
//import "fmt"
//
//func main() {
//	var hard string = "今天天气不错"
//	fmt.Println(hard)
//}

//
//func Enter() {
//	date := "2026-1-28"
//	page := "学习go语言"
//	enterher := "你好，今天是%s,我正在%s"
//	enterbro := "当前的学习进度为，刚开始!"
//	Introduction := fmt.Sprintf(enterher, date, page)
//	fmt.Println(Introduction + enterbro)
//}
//
//func unchanged() {
//	data := 123
//	sun_data := data * 321
//	fmt.Println(sun_data)
//}
//
//func splice() {
//	str_one := "今年过年"
//	str_two := "不收礼，"
//	str_three := "收礼只收脑白金！"
//	Spring_Festival := fmt.Sprintf(str_one + str_two + str_three)
//	fmt.Println(Spring_Festival)
//
//}

//func login() {
//	var name string
//	var password int
//
//	fmt.Print("请输入账号：")
//	fmt.Scanln(&name)
//
//	fmt.Print("请输入密码：")
//	fmt.Scanln(&password)
//
//	if name == "张三" && password == 123456 {
//		fmt.Println("登录成功")
//	} else {
//		fmt.Println("登录失败")
//	}
//
//}

//func enter_int() {
//	var intsr []int
//	var intse []int
//	for i := 1; i <= 100; i++ {
//		if i >= 50 {
//			intsr = append(intsr, i)
//		}
//		if i < 50 {
//			intse = append(intse, i)
//		}
//
//	}
//	fmt.Println(intsr)
//	fmt.Println(intse)
//	fmt.Println(len(intse) - 2)
//}

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func mysql_connect() *sql.DB {
	//格式："用户名:密码@tcp(IP:端口)/数据库名?参数"
	connect := "root:123456@tcp(127.0.0.1:3306)/zhq?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", connect)
	if err != nil {
		fmt.Println("格式错误")
		return nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("链接失败")
		_ = db.Close()
		return nil
	} else {
		fmt.Println("链接成功")
	}
	return db
}

func login() {
	// 1. 获取用户输入
	var inputName, inputPwd string
	fmt.Print("请输入账号: ")
	fmt.Scanln(&inputName)
	fmt.Print("请输入密码: ")
	fmt.Scanln(&inputPwd)

	// 2. 连接数据库
	db := mysql_connect()
	if db == nil {
		fmt.Println("连接失败")
		return
	}
	defer db.Close()

	// 3. 直接查这个账号
	var dbPwd string
	err := db.QueryRow("SELECT password FROM user WHERE name = ?", inputName).Scan(&dbPwd)

	if err == sql.ErrNoRows {
		fmt.Println("账号不存在")
		return
	}
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}

	// 4. 验证密码
	if dbPwd == inputPwd {
		fmt.Println("登录成功！")
	} else {
		fmt.Println("密码错误！")
	}
}

//func mysql_user() {
//	type user struct {
//		name     string
//		password string
//	}
//
//	db := mysql_connect()
//	if db == nil {
//		fmt.Println("你的链接失败了")
//		return
//	}
//
//	defer db.Close()
//	rows, err := db.Query("select name,password from user")
//	if err != nil {
//		fmt.Println("查询失败")
//		return
//	}
//	defer rows.Close()
//	var users []user
//
//	for rows.Next() {
//		var u user
//		err := rows.Scan(&u.name, &u.password)
//		if err != nil {
//			fmt.Println("扫描失败")
//			return
//		}
//		users = append(users, u)
//	}
//	fmt.Println(users)

//}

//	func login() {
//		var name string
//		var password string
//		var hard = []string{"admin", "123456"}
//
//		hard_name := hard[0]
//		hard_password := hard[1]
//
//		fmt.Print("请输入账号:")
//		fmt.Scanln(&name)
//
//		fmt.Print("请输入密码：")
//		fmt.Scanln(&password)
//
//		if name != hard_name {
//			fmt.Println("账号不存在！")
//		}
//
//		if name == hard_name && password != hard_password {
//			fmt.Println("密码错误")
//		}
//
//		if name == hard_name && password == hard_password {
//			fmt.Println("登录成功，账号为" + name)
//		}
//	}
func main() {
	login()
}
