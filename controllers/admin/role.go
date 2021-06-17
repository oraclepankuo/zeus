package admin

import (
	"zeus/models"
	"strconv"
	"strings"
)

type RoleController struct {
	BaseController
}

func (c *RoleController) Get() {

	role := []models.Role{}
	models.DB.Find(&role)
	c.Data["roleList"] = role
	c.TplName = "admin/role/index.html"
}

func (c *RoleController) Add() {
	c.TplName = "admin/role/add.html"
}

func (c *RoleController) DoAdd() {
	title := strings.Trim(c.GetString("title"), "")
	description := strings.Trim(c.GetString("description"), "")

	if title == "" {
		c.Error("标题不能为空", "/role/add")
		return
	}

	role := models.Role{}
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(models.GetUnix())

	err := models.DB.Create(&role).Error
	if err != nil {
		c.Error("增加角色失败", "/role/add")
	} else {
		c.Success("增加角色成功", "/role")
	}
}

func (c *RoleController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/role/")
		return
	}

	role := models.Role{Id: id}
	models.DB.Find(&role)
	c.Data["role"] = role
	c.TplName = "admin/role/edit.html"
}

func (c *RoleController) DoEdit() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/role/")
		return
	}

	title := strings.Trim(c.GetString("title"), "")
	description := strings.Trim(c.GetString("description"), "")

	if title == "" {
		c.Error("标题不能为空", "/role/add")
		return
	}
	role := models.Role{Id: id}
	models.DB.Find(&role)

	role.Title = title
	role.Description = description

	err2 := models.DB.Save(&role).Error

	if err2 != nil {
		c.Error("修改数据失败", "/role/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改角色成功", "/role")
	}

}

func (c *RoleController) Delete() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/role/")
		return
	}

	role := models.Role{Id: id}
	models.DB.Delete(&role)

	c.Success("删除角色成功", "/role")

}

func (c *RoleController) Auth() {
	roleId, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/role")
		return
	}

	access := []models.Access{}
	models.DB.Preload("AccessItem").Where("module_id=0").Find(&access)

	roleAccess := []models.RoleAccess{}
	models.DB.Where("role_id=?", roleId).Find(&roleAccess)

	roleAccessMap := make(map[int]int)

	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = v.AccessId
	}

	for i := 0; i < len(access); i++ {
		_, ok := roleAccessMap[access[i].Id]
		if ok {
			access[i].Checked = true
		}

		for j := 0; j < len(access[i].AccessItem); j++ {

			if _, ok := roleAccessMap[access[i].AccessItem[j].Id]; ok {
				access[i].AccessItem[j].Checked = true
			}
		}
	}
	c.Data["accessList"] = access
	c.Data["roleId"] = roleId
	c.TplName = "admin/role/auth.html"
}
func (c *RoleController) DoAuth() {
	//1、获取参数post传过来的角色id 和 权限切片
	roleId, err := c.GetInt("role_id")
	if err != nil {
		c.Error("传入参数错误", "/role")
		return
	}
	accessNode := c.GetStrings("access_node")

	//2、修改角色权限---删除当前角色下面的所有权限

	roleAccess := models.RoleAccess{}
	models.DB.Where("role_id=?", roleId).Delete(&roleAccess)

	//3、执行增加数据

	for _, v := range accessNode {
		accessId, _ := strconv.Atoi(v)
		roleAccess.AccessId = accessId
		roleAccess.RoleId = roleId
		models.DB.Create(&roleAccess)
	}
	c.Success("授权成功", "/role/auth?id="+strconv.Itoa(roleId))
}
