package main

import (
	"fmt"

	redis "github.com/gomodule/redigo/redis"
)

func main() {
	redisCon, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer redisCon.Close()

	redisCon.Do("SET", "key1", "hello world")
	res, err := redis.String(redisCon.Do("GET", "key1"))
	fmt.Printf("%#v\n", res)
}
