package main

import (
	"github.com/codegangsta/martini"
	"lark/handlers"
	"lark/orm"
	"lark/redisdb"
)

func main() {
	m = martini.New()

	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(RedisDb())
	m.Use(Orm())

	r := martini.NewRouter()
	r.Post("/exps", NewExp)
	r.Get("/exps/:id", ExpById)
	r.Get("/exps/:name", ExpBtName)
	r.Get("/exps", Exps)
	r.Post("/exps/:name/winner", SetExpWinner)
	r.Post("/exps/:name/reset", ResetExp)

	r.Get("/participate/")

	m.MapTo(db, (*RedisDB)(nil))
	m.MapTo(orm, (*Orm)(nil))

	m.Action(r.Handle())
	m.Run()
}
