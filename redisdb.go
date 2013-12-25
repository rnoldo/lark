package lark

import (
	"github.com/codegangsta/martini"
	"github.com/garyburd/redigo/redis"
)

func RedisDB() martini.Handler {
	conn, err = redis.Dial("localhost", ":6379")
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		c.Map(conn)
		defer conn.Close()
		c.Next()
	}
}
