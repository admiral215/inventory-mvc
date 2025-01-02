package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"inventory-bee/dto"
	"inventory-bee/helpers"
	"inventory-bee/repositories"
	"inventory-bee/services"
	"strconv"
)

type CategoryController struct {
	beego.Controller
	Service services.CategoryService
}

func NewCategoryController() *CategoryController {
	repo := repositories.NewCategoryRepository()
	service := services.NewCategoryService(repo)
	return &CategoryController{
		Service: service,
	}
}

func (c *CategoryController) Index() {
	helpers.GetFlash(&c.Controller)

	search := c.GetString("search", "")

	categories, err := c.Service.GetAllBySearch(search)
	if err != nil {
		c.Data["error"] = err.Error()
	}

	c.Data["search"] = search
	c.Data["categories"] = categories
	c.TplName = "categories/index.html"
}

func (c *CategoryController) ShowCreate() {
	c.Data["Dto"] = dto.CategoryCreate{}
	c.TplName = "categories/create.html"
}

func (c *CategoryController) SubmitCreate() {
	createDto := dto.CategoryCreate{}

	if err := c.ParseForm(&createDto); err != nil {
		c.Data["Errors"] = []string{err.Error()}
		c.TplName = "categories/create.html"
		return
	}

	err := helpers.ValidateDto(&createDto)
	if err != nil {
		c.Data["Errors"] = err
		c.Data["Dto"] = createDto
		c.TplName = "categories/create.html"
		return
	}

	err = c.Service.Create(&createDto)
	if err != nil {
		c.Data["Errors"] = err
		c.TplName = "categories/create.html"
		return
	}

	c.Redirect("/categories", 302)
}

func (c *CategoryController) ShowEdit() {
	helpers.GetFlash(&c.Controller)

	c.TplName = "categories/edit.html"

	categoryIdStr := c.Ctx.Input.Param(":id")

	categoryId, err := strconv.ParseUint(categoryIdStr, 10, 32)

	category, err := c.Service.GetById(uint(categoryId))
	if err != nil {
		c.Data["Errors"] = err.Error()
		return
	}

	c.Data["Category"] = category
}

func (c *CategoryController) SubmitEdit() {
	categoryId, _ := c.GetUint8("id")
	flash := beego.NewFlash()
	redirectUrl := fmt.Sprintf("/categories/%d", categoryId)

	handleError := func(message string, err any) {
		flash.Error(message, err)
		flash.Store(&c.Controller)
		c.Redirect(redirectUrl, 302)
	}

	updateDto := dto.CategoryUpdate{}
	if err := c.ParseForm(&updateDto); err != nil {
		handleError("error parsing form", err)
		return
	}

	err := helpers.ValidateDto(&updateDto)
	if err != nil {
		handleError("invalid input", err)
		return
	}

	err = c.Service.Edit(&updateDto)
	if err != nil {
		handleError("error updating data", err)
		return
	}

	flash.Success("Successfully Edited data")
	flash.Store(&c.Controller)

	c.Redirect("/categories", 302)
}

func (c *CategoryController) SubmitDelete() {
	flash := beego.NewFlash()

	categoryIdStr := c.Ctx.Input.Param(":id")
	categoryId, err := strconv.ParseUint(categoryIdStr, 10, 32)
	if err != nil {
		flash.Error("Invalid category id")
	}

	err = c.Service.DeleteById(uint(categoryId))
	if err != nil {
		flash.Error("Failed to delete category")
	}

	c.Redirect("/categories", 302)
}
