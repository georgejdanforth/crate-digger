package models

type Search struct {
	Entity string `form:"entity"`
	Query  string `form:"query"`
}
