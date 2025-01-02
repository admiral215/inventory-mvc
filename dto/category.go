package dto

type CategoryCreate struct {
	Name string `form:"name" valid:"Required;MinSize(3);MaxSize(100)"`
}

type CategoryUpdate struct {
	Id   uint   `form:"id" valid:"Required"`
	Name string `form:"name" valid:"Required;MinSize(3);MaxSize(100)"`
}
