package main

import "github.com/jgarc172/gocsprog/lab6-1/myclass"

func main() {
	mc := myclass.New()
	mc.GetFields()
	mc.SetFields()
	mc.GetFields()
	myclass.ModFields()
	mc.GetFields()
	mc = myclass.New()
	mc.GetFields()
}
