package controllers

import (
	"github.com/revel/revel"
)

//App is to create a controller
type App struct {
	*revel.Controller
}

//Index  helps in rendering
func (c App) Index() revel.Result {
	return c.Render()
}
