package main

import (
	"github.com/astaxie/beego"
	_ "tarobackend/routers"
	_ "tarobackend/utils"
)

func main() {
	beego.Run()
}
