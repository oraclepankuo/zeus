package admin

import (
	"zeus/models"
	"strconv"
	"strings"
)

type ManagerController struct {
	BaseController
}

func (c *ManagerController) Get() {
	manager := []models.Manager{}
	models.DB.Preload("Role").Find(&manager)

	c.Data["managerList"] = manager
	c.TplName = "admin/manager/index.html"
}

func (c *ManagerController) Add() {
	role := []models.Role{}
	models.DB.Find(&role)
	c.Data["roleList"] = role
	c.TplName = "admin/manager/add.html"
}

func (c *ManagerController) DoAdd() {
	roleId, err1 := c.GetInt("role_id")
	if err1 != nil {
		c.Error("非法求情", "/manager/add")
		return
	}
	username := strings.Trim(c.GetString("username"), " ")
	password := strings.Trim(c.GetString("password"), " ")
	mobile := strings.Trim(c.GetString("mobile"), " ")
	email := strings.Trim(c.GetString("email"), " ")

	if len(username) < 2 || len(password) < 6 {
		c.Error("用户名或密码不合法", "/manager/add")
		return
	}

	managerList := []models.Manager{}
	models.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		c.Error("用户名已经存在", "/manager/add")
		return
	}

	manager := models.Manager{}
	manager.Username = username
	manager.Password = models.Md5(password)
	manager.Mobile = mobile
	manager.Email = email
	manager.Status = 1
	manager.AddTime = int(models.GetUnix())
	manager.RoleId = roleId
	err := models.DB.Create(&manager).Error
	if err != nil {
		c.Error("增加管理员失败", "/manager/add")
		return
	}

	c.Success("增加管理员成功", "/manager")

}

func (c *ManagerController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("非法请求", "/manager")
		return
	}
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)
	c.Data["managerList"] = manager

	role := []models.Role{}
	models.DB.Find(&role)
	c.Data["roleList"] = role

	c.TplName = "admin/manager/edit.html"
}
func (c *ManagerController) DoEdit() {

	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("非法求情", "/manager")
		return
	}

	roleId, err2 := c.GetInt("role_id")
	if err2 != nil {
		c.Error("非法求情", "/manager")
		return
	}

	password := strings.Trim(c.GetString("password"), " ")
	mobile := strings.Trim(c.GetString("mobile"), " ")
	email := strings.Trim(c.GetString("email"), " ")

	manager := models.Manager{Id: id}
	models.DB.Find(&manager)
	manager.RoleId = roleId
	manager.Mobile = mobile
	manager.Email = email
	if password != "" {
		if len(password) < 6 {
			c.Error("密码长度不合法,小于6位数", "/manager/edit?id="+strconv.Itoa(id))
			return
		}
		manager.Password = models.Md5(password)
	}

	err3 := models.DB.Save(&manager).Error
	if err3 != nil {
		c.Error("修改数据失败", "/manager/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改数据成功", "/manager")
	}

}

func (c *ManagerController) Delete() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/manager")
		return
	}

	manager := models.Manager{Id: id}
	models.DB.Delete(&manager)

	c.Success("删除管理员成功", "/manager")
}
