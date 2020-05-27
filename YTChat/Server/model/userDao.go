package model

import (
	"awesomeProject/YTChat/Common"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type UserDao struct{
	pool *redis.Pool
}


//饿汉单例：提前创建UserDao实例提供服务
var (
	MyUserDao *UserDao
)

//工厂模式，创建UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao){
	userDao = &UserDao{
		pool: pool,
	}
	return
}


func (t *UserDao) getUserById(conn redis.Conn, id int) (user *Common.User, err error){
	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil{
		if err == redis.ErrNil{
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &Common.User{
	}
	//redis存的json格式的res需要反序列化-->结构体
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal error: ", err)
		return
	}

	return
}

func (t *UserDao) Login(userId int, userPwd string) (user *Common.User, err error){
	conn := t.pool.Get()
	defer conn.Close()

	user, err = t.getUserById(conn, userId)
	if err != nil{
		return
	}
	if user.UserPwd != userPwd{
		err = ERROR_USER_PWD
		return
	}

	return
}