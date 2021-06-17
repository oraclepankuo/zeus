package admin

type BatchexecController struct {
	BaseController
}

func (c *BatchexecController) Ssh() {
	c.TplName = "admin/batchexec/ssh-index.html"
}

func (c *BatchexecController) SaltStackApi() {
	c.TplName = "admin/batchexec/saltstack-index.html"
}
