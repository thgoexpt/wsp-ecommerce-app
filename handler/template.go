package handler

import "html/template"

var t *template.Template
var e error

func init() {
	fm := template.FuncMap{
		"Iterate": func(count int) []int {
			var i int
			var Items []int
			for i = 1; i <= (count); i++ {
				Items = append(Items, i)
			}
			return Items
		},
	}
	t, e = template.New("default").Funcs(fm).ParseGlob("template/*")
}

func Validate() {
	if e != nil {
		panic(e)
	}
}
