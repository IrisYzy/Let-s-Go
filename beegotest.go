package main

import "github.com/astaxie/beego"

type Maincontroller struct {
	beego.Controller
}

func (this *Maincontroller) Get() {
	this.Ctx.WriteString("hello world")
}

func main()  {
	beego.Router("/",&Maincontroller{})
	beego.Run()
}