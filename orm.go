package lark

import (
	"database/sql"
	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

func Orm() martini.Handler {
	return func(c martini.Context) {
		orm, err = initGorp()
		if err != nil {
			panic(err)
		}
		c.Map(orm)
		defer orm.Db.Close()
		c.Next()
	}
}

func initGorp() (*gorp.DbMap, error) {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Experiment{}, "experiments").SetKeys(true, "Id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap, err
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
