package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"inventory-bee/conf"
	"inventory-bee/controllers"
)

func init() {
	conf.ConnectDB()
	categoryController := controllers.NewCategoryController()

	beego.Router("/categories", categoryController, "get:Index")

	beego.Router("/categories/create", categoryController, "get:ShowCreate;post:SubmitCreate")

	beego.Router("/categories/:id", categoryController, "get:ShowEdit")
	beego.Router("/categories/edit", categoryController, "post:SubmitEdit")
	beego.Router("/categories/delete/:id", categoryController, "post:SubmitDelete")
}
