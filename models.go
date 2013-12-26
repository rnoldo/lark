package lark

import (
	""
)

type Client struct {
	Id int64
}

//实验
type Experiment struct {
	Id int64 'db:"exp_id"'
	Created int64
	Name   string
	Description string
	Winner *Alternative
	Alts []Alternative
}

//AB选项
type Alternative struct {
	experiment Experiment
	name       string
	is_winner bool
}
