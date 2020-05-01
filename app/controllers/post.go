package controllers

import (
	"errors"
	"myapp/app/models"

	"github.com/revel/revel"
)

//Post struct is the one which will be used in this controller
type Post struct {
	*revel.Controller
}

//Index Returning the complete set of posts
func (c Post) Index() revel.Result {
	posts := []models.Post{}

	result := DB.Order("id desc").Find(&posts)
	err := result.Error
	if err != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}

	return c.Render(posts)
}

//Create will create a new Post
func (c Post) Create() revel.Result {
	post := models.Post{
		Body: c.Params.Form.Get("body"),
	}

	ret := DB.Create(&post)
	if ret.Error != nil {
		return c.RenderError(errors.New("Record Create failure." + ret.Error.Error()))
	}

	return c.Redirect("/posts")
}

//Delete a single Post record
func (c Post) Delete() revel.Result {
	id := c.Params.Route.Get("id")
	posts := []models.Post{}
	ret := DB.Delete(&posts, id)
	if ret.Error != nil {
		return c.RenderError(errors.New("Record Delete failure." + ret.Error.Error()))
	}
	return c.Redirect("/posts")
}

//RedirectToPosts for testing purpose
func (c Post) RedirectToPosts() revel.Result {
	return c.Redirect("/posts")
}
