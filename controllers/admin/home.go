package admin

type HomeController struct {
	BaseController
}

func (c *HomeController) Home() {

	c.TplName = "admin/home/home.html"
}
