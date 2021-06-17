package admin

import (
	"zeus/models"
	"strconv"
)

type AssetController struct {
	BaseController
}

func (c *AssetController) Host() {
	//当前页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 9

	//实现搜索功能
	keyword := c.GetString("keyword")
	where := "1=1"
	if len(keyword) > 0 {
		where += " AND title like \"%" + keyword + "%\""
	}
	//查询数据
	hostsList := []models.Host{}
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&hostsList)

	//判断是否有数据，如果没有数据跳转到上一页
	if len(hostsList) == 0 {
		prvPage := page - 1
		if prvPage == 0 {
			prvPage = 1
		}
		c.Goto("/goods?page=" + strconv.Itoa(prvPage))
	}

	//查询goods表里面的数量
	var count int
	models.DB.Where(where).Table("host").Count(&count)
	var num int

	if count%pageSize == 0 {
		num = 0
	} else if count%pageSize != 0 {
		num = pageSize
	}

	c.Data["count"] = count
	c.Data["hostsList"] = hostsList
	c.Data["num"] = num
	c.Data["page"] = page
	c.Data["keyword"] = keyword
	c.TplName = "admin/asset/host-index.html"
}

func (c *AssetController) HostAdd() {
	c.TplName = "admin/asset/host-add.html"
}

func (c *AssetController) HostEdit() {
	c.TplName = "admin/asset/host-edit.html"
}

func (c *AssetController) Phost() {
	//当前页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 9

	//实现搜索功能
	keyword := c.GetString("keyword")
	where := "1=1"
	if len(keyword) > 0 {
		where += " AND title like \"%" + keyword + "%\""
	}
	//查询数据
	phostsList := []models.Phost{}
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&phostsList)

	//判断是否有数据，如果没有数据跳转到上一页
	if len(phostsList) == 0 {
		prvPage := page - 1
		if prvPage == 0 {
			prvPage = 1
		}
		c.Goto("/goods?page=" + strconv.Itoa(prvPage))
	}

	//查询goods表里面的数量
	var count int
	models.DB.Where(where).Table("phost").Count(&count)
	var num int

	if count%pageSize == 0 {
		num = 0
	} else if count%pageSize != 0 {
		num = pageSize
	}

	c.Data["count"] = count
	c.Data["phostsList"] = phostsList
	c.Data["num"] = num
	c.Data["page"] = page
	c.Data["keyword"] = keyword

	c.TplName = "admin/asset/phost-index.html"
}

func (c *AssetController) PhostAdd() {
	c.TplName = "admin/asset/phost-add.html"
}

func (c *AssetController) PhostEdit() {
	c.TplName = "admin/asset/phost-edit.html"
}

func (c *AssetController) Vhost() {
	//当前页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 9

	//实现搜索功能
	keyword := c.GetString("keyword")
	where := "1=1"
	if len(keyword) > 0 {
		where += " AND title like \"%" + keyword + "%\""
	}
	//查询数据
	vhostsList := []models.Vhost{}
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&vhostsList)

	//判断是否有数据，如果没有数据跳转到上一页
	if len(vhostsList) == 0 {
		prvPage := page - 1
		if prvPage == 0 {
			prvPage = 1
		}
		c.Goto("/goods?page=" + strconv.Itoa(prvPage))
	}

	//查询goods表里面的数量
	var count int
	models.DB.Where(where).Table("vhost").Count(&count)
	var num int

	if count%pageSize == 0 {
		num = 0
	} else if count%pageSize != 0 {
		num = pageSize
	}

	c.Data["count"] = count
	c.Data["vhostsList"] = vhostsList
	c.Data["num"] = num
	c.Data["page"] = page
	c.Data["keyword"] = keyword
	c.TplName = "admin/asset/vhost-index.html"
}

func (c *AssetController) VhostAdd() {
	c.TplName = "admin/asset/vhost-add.html"
}

func (c *AssetController) VhostEdit() {
	c.TplName = "admin/asset/vhost-edit.html"
}

func (c *AssetController) NetworkDevice() {
	//当前页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 7

	//实现搜索功能
	keyword := c.GetString("keyword")
	where := "1=1"
	if len(keyword) > 0 {
		where += " AND title like \"%" + keyword + "%\""
	}
	//查询数据
	networkDevicesList := []models.NetworkDevice{}
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&networkDevicesList)

	//判断是否有数据，如果没有数据跳转到上一页
	if len(networkDevicesList) == 0 {
		prvPage := page - 1
		if prvPage == 0 {
			prvPage = 1
		}
		c.Goto("/goods?page=" + strconv.Itoa(prvPage))
	}

	//查询goods表里面的数量
	var count int
	models.DB.Where(where).Table("networkdevice").Count(&count)
	var num int

	if count%pageSize == 0 {
		num = 0
	} else if count%pageSize != 0 {
		num = pageSize
	}

	c.Data["count"] = count
	c.Data["networkDevicesList"] = networkDevicesList
	c.Data["num"] = num
	c.Data["page"] = page
	c.Data["keyword"] = keyword
	c.TplName = "admin/asset/networkdevice-index.html"
}

func (c *AssetController) NetworkDeviceAdd() {
	c.TplName = "admin/asset/networkdevice-add.html"
}

func (c *AssetController) NetworkDeviceEdit() {
	c.TplName = "admin/asset/networkdevice-edit.html"
}

func (c *AssetController) DomainName() {
	//当前页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 7

	//实现搜索功能
	keyword := c.GetString("keyword")
	where := "1=1"
	if len(keyword) > 0 {
		where += " AND title like \"%" + keyword + "%\""
	}
	//查询数据
	domainNameList := []models.DomainNameList{}
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&domainNameList)

	//判断是否有数据，如果没有数据跳转到上一页
	if len(domainNameList) == 0 {
		prvPage := page - 1
		if prvPage == 0 {
			prvPage = 1
		}
		c.Goto("/goods?page=" + strconv.Itoa(prvPage))
	}

	//查询goods表里面的数量
	var count int
	models.DB.Where(where).Table("domainnameList").Count(&count)
	var num int

	if count%pageSize == 0 {
		num = 0
	} else if count%pageSize != 0 {
		num = pageSize
	}

	c.Data["count"] = count
	c.Data["domainNameList"] = domainNameList
	c.Data["num"] = num
	c.Data["page"] = page
	c.Data["keyword"] = keyword
	c.TplName = "admin/asset/domainname-index.html"
}

func (c *AssetController) DomainNameAdd() {
	c.TplName = "admin/asset/domainname-add.html"
}

func (c *AssetController) DomainNameEdit() {
	c.TplName = "admin/asset/domainname-edit.html"
}

func (c *AssetController) Certificate() {
	//当前页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 7

	//实现搜索功能
	keyword := c.GetString("keyword")
	where := "1=1"
	if len(keyword) > 0 {
		where += " AND title like \"%" + keyword + "%\""
	}
	//查询数据
	certificateList := []models.CertificateList{}
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&certificateList)

	//判断是否有数据，如果没有数据跳转到上一页
	if len(certificateList) == 0 {
		prvPage := page - 1
		if prvPage == 0 {
			prvPage = 1
		}
		c.Goto("/goods?page=" + strconv.Itoa(prvPage))
	}

	//查询goods表里面的数量
	var count int
	models.DB.Where(where).Table("certificatelist").Count(&count)
	var num int

	if count%pageSize == 0 {
		num = 0
	} else if count%pageSize != 0 {
		num = pageSize
	}

	c.Data["count"] = count
	c.Data["certificateList"] = certificateList
	c.Data["num"] = num
	c.Data["page"] = page
	c.Data["keyword"] = keyword
	c.TplName = "admin/asset/certificate-index.html"
}

func (c *AssetController) CertificateAdd() {
	c.TplName = "admin/asset/certificate-add.html"
}

func (c *AssetController) CertificateEdit() {
	c.TplName = "admin/asset/certificate-edit.html"
}
