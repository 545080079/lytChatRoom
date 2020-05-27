package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func redisTest(conn redis.Conn){
	conn.Do("set","lytmm", "13859984262")
    res, _ := redis.String(conn.Do("get", "lytmm"))
	fmt.Println("res = ", res)
}

var pool *redis.Pool

func init(){
	fmt.Println("init redis...")
	pool = &redis.Pool{
		MaxIdle: 8,		//Max空闲连接数
		MaxActive: 0,	//Max连接数
		IdleTimeout: 300, //Max Time out
		Dial: func()(redis.Conn, error){
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main(){
	conn := pool.Get()	//get a link
	defer conn.Close()
	redisTest(conn)
}