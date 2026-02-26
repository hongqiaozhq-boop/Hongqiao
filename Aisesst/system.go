package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// mysql_connect 连接数据库（参考之前的方法）
func mysql_connect() *sql.DB {
	connect := "root:123456@tcp(127.0.0.1:3306)/zhq?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", connect)
	if err != nil {
		fmt.Println("数据库配置错误:", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		_ = db.Close()
		return nil
	}
	fmt.Println("数据库连接成功")
	return db
}

// register 注册新用户
// 数据库表结构：id（自增）, name, password, statu（默认值）
func register() {
	// 1. 连接数据库
	db := mysql_connect()
	if db == nil {
		fmt.Println("数据库获取失败，无法进行注册")
		return
	}
	defer db.Close()
	fmt.Println("数据库连接成功，准备注册...\n")

	for {
		// 2. 获取用户输入
		var inputName, inputPwd string
		fmt.Print("请输入账号: ")
		fmt.Scanln(&inputName)
		fmt.Print("请输入密码: ")
		fmt.Scanln(&inputPwd)

		// 验证输入不能为空
		if inputName == "" {
			fmt.Println("账号不能为空，请重新输入\n")
			continue
		}
		if inputPwd == "" {
			fmt.Println("密码不能为空，请重新输入\n")
			continue
		}

		// 3. 检查账号是否已存在
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM user WHERE name = ?", inputName).Scan(&count)
		if err != nil {
			fmt.Println("检查账号失败:", err, "，请重新输入\n")
			continue
		}
		if count > 0 {
			fmt.Println("该账号已存在，请使用其他账号\n")
			continue
		}

		// 4. 插入新用户
		// id 字段自增，statu 字段使用默认值
		result, err := db.Exec("INSERT INTO user (name, password) VALUES (?, ?)", inputName, inputPwd)
		if err != nil {
			fmt.Println("注册失败:", err, "，请重新输入\n")
			continue
		}

		// 5. 获取插入的用户ID
		lastID, err := result.LastInsertId()
		if err != nil {
			fmt.Println("获取用户ID失败:", err, "，但注册已成功")
			fmt.Printf("注册成功！账号: %s\n", inputName)
			return
		}

		fmt.Printf("注册成功！用户ID: %d, 账号: %s\n", lastID, inputName)
		return
	}
}

// main 测试注册功能
func main() {
	register()
}
