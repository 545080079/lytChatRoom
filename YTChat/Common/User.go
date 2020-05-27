package Common

type User struct{

	//为了序列化、反序列化成功，必须保证
	//用户信息json字符串和结构体字段对应的tag要一致
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}