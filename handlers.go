package lark

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
	"strconv"
)

//新建一个实验
func NewExp(r *http.Request, orm *orm, params martini.Params) (int string) {

}

//获取实验列表
func Exps(orm *Orm, r render.Render) {
	var exps []Experiment
	_, err = orm.Select(&exps, "select * from experiments order by exp_id")
	if err != nil {
		r.JSON(500, "Internl Error")
	}
	r.JSON(200, exps)
}

// r.Get("/exps/:name")
func ExpByName(orm *Orm, r render.Render, name string) {
	var exp = Experiment
	err = orm.SelectOne(&exp, "select * from experiments where exp_name=?", name)

}

// r.Get("/exps/:id")
func ExpById(orm *Orm, r render.Render, id int64) {
	var exp = Experiment
	err = orm.SelectOne(&exp, "select * from experiments where exp_id=?", id)
	if err != nil {
		r.JSON(500, "error")
	}

	r.JSON(200, exp)
}

//r.Post("exps/:name/winner", SetExpWinner)
func SetExpWinner(req *http.Request, orm *Orm, r render.Render, name string) {
	var exp = Experiment
	err = orm.SelectOne(&exp, "select * from experiments where exp_name=?", name)
	if err != nil {

	}
	exp.winner = req.FormValue(alt_name)

	err = orm.Update(&exp)

	if err != nil {
		r.Error("500")
	}

	r.JSON(200, exp)
}

//r.Post("exps/:name/reset", ResetExp)
func ResetExp(req *http.Request, orm Orm, r render.Render, name string) {

}

//r.Post("exps/:name/winner/reset", ResetExpWinner)
func ResetExpWinner(req *http.Request, orm Orm, r render.Render, name string) {

}

//r.Post("exps/:name/delete", DeleteExp)
func DeleteExp(req *http.Request, orm Orm, r render.Render, name string) {

}

//用户主页
func UserHome(orm *Orm, r render.Render) {

}

func Archived(orm *Orm, r render.Render) {
	experiments = orm.GetAllExps()
	r.HTML(200, "dashbord", exps)
	r.JSON(200, exps)
}
