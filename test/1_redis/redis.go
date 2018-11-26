package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	//连接redis
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	//通过Do函数，发送redis命令
	v, err := c.Do("SET", "name", "red")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("v:", v, err)
	v, err = redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("v:", v)

	//操作列表
	c.Do("lpush", "redlist", "qqq")
	c.Do("lpush", "redlist", "www")
	c.Do("lpush", "redlist", "eee")
	_, err = c.Do("del", "redlist")
	//读取列表
	values, _ := redis.Values(c.Do("lrange", "redlist", "0", "100"))
	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}

	// 或者
	var v1 string
	redis.Scan(values, &v1)
	fmt.Println("v1：", v1)

	size, err := c.Do("DBSIZE")
	fmt.Printf("size is %d \n", size)
}
