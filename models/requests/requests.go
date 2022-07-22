package requests

type Some struct {
	Name string `form:"name" binding:"required"`
	Url  string `form:"url" binding:"required,correctUrl"`
}
