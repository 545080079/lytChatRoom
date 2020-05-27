package model

import (
	"awesomeProject/YTChat/Common"
	"net"
)

type CurUser struct{
	Conn net.Conn
	message Common.User
	UserId int
	UserName string
}