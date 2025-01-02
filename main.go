package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "inventory-bee/routers"
)

func main() {
	err := beego.AddFuncMap("isString", func(i interface{}) bool {
		_, ok := i.(string)
		return ok
	})

	if err != nil {
		return
	}

	beego.Run()
}
