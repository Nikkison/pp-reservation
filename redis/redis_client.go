package main

import (
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
)

func redisSet(key string, value string, c redis.Conn) {
	c.Do("SET", key, value)
}

func redisGet(key string, c redis.Conn) string {
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return s
}

func redisConnection() redis.Conn {
	const port = ":6379"

	//redisに接続
	c, err := redis.Dial("tcp", port)
	if err != nil {
		panic(err)
	}
	return c
}

func main() {
	// 1. sudo apt-get -y install redis-server
	// 2. sudo service redis startd
	// 3. redis-cli
	// 3.1 go get github.com/garyburd/redigo/redis
	c := redisConnection()
	defer c.Close()

	var key = "KEY"
	var val = "VALUE"
	redisSet(key, val, c)
	s := redisGet(key, c)
	fmt.Println(s)
	// 4. keys *
}
