package admin

import (
	"errors"
	"os"
	"path"
	"strconv"
	"strings"
	"zeus/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Success(message string, redirect string) {
	//http://localhost:8080/beego_admin/goods?page=2
	c.Data["message"] = message
	if strings.Contains(redirect, "http") {
		c.Data["redirect"] = redirect
	} else {
		c.Data["redirect"] = "/" + beego.AppConfig.String("adminPath") + redirect
	}
	c.TplName = "admin/public/success.html"
}

func (c *BaseController) Error(message string, redirect string) {
	c.Data["message"] = message
	if strings.Contains(redirect, "http") {
		c.Data["redirect"] = redirect
	} else {
		c.Data["redirect"] = "/" + beego.AppConfig.String("adminPath") + redirect
	}
	c.TplName = "admin/public/error.html"
}
func (c *BaseController) Goto(redirect string) {
	c.Redirect("/"+beego.AppConfig.String("adminPath")+redirect, 302)
}

//本地上传
func (c *BaseController) LoclUploadImg(picName string) (string, error) {
	//1、获取上传的文件
	f, h, err := c.GetFile(picName)
	if err != nil {
		return "", err
	}
	//2、关闭文件流
	defer f.Close()
	//3、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(h.Filename)

	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {

		return "", errors.New("图片后缀名不合法")
	}
	//4、创建图片保存目录  static/upload/20200623
	day := models.GetDay()
	dir := "static/upload/" + day

	if err := os.MkdirAll(dir, 0666); err != nil {
		return "", err
	}
	//5、生成文件名称   144325235235.png
	fileUnixName := strconv.FormatInt(models.GetUnixNano(), 10)
	//static/upload/20200623/144325235235.png
	saveDir := path.Join(dir, fileUnixName+extName)
	//6、保存图片
	c.SaveToFile(picName, saveDir)

	return saveDir, nil
}
