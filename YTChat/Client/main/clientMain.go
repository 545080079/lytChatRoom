package main

import (
	"awesomeProject/YTChat/Client/processor"
	"fmt"
	"os"
)

func main(){

	for true{
		fmt.Println("***\t\t\t\tTT聊天室\t\t\t\t***")
		fmt.Println("***\t\t\t1. 登录\t\t\t\t***")
		fmt.Println("***\t\t\t2. 注册\t\t\t\t***")
		fmt.Println("***\t\t\t3. 退出\t\t\t\t***")

		var key int
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			var userId int
			var userPwd string
			fmt.Println("输入用户ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("输入用户密码")
			fmt.Scanf("%s", &userPwd)
			//完成登录
			up := &processor.UserProcessor{
			}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("Register")
		case 3:
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Error Input!")
		}

	}


}
