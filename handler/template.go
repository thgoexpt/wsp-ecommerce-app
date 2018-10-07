package handler

import "html/template"

var t,e = template.ParseGlob("template/*")

func Validate() {
	if e != nil {
		panic(e)
	}
}