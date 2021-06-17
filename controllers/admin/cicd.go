package admin

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bndr/gojenkins"
)

type CicdController struct {
	BaseController
}

func arrayToString(arr []string) string {

	var result string
	for _, i := range arr {
		result += i
	}
	return result

}

func (c *CicdController) JenkinsApi() {

	ctx := context.Background()
	jenkins := gojenkins.CreateJenkins(nil, "http://10.202.0.30:8080", "admin", "Pa$$w0rd")
	_, err1 := jenkins.Init(ctx)
	if err1 != nil {
		beego.Error(err1)
	} else {
		beego.Info("连接Jenkins 成功")
	}

	//所有工程列表
	jobs, err2 := jenkins.GetAllJobNames(ctx)
	if err2 != nil {
		fmt.Println(err2)
	}

	JobsListLen := len(jobs)
	JobListName := make([]string, JobsListLen, JobsListLen)

	for index, job := range jobs {
		JobListName[index] = job.Name
	}

	c.Data["JobListName"] = JobListName
	c.TplName = "admin/cicd/jenkinsci-index.html"

	//jobName := "testA1"
	//builds, err := jenkins.GetAllBuildIds(ctx, jobName)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, build := range builds {
	//	buildId := build.Number
	//	data, err := jenkins.GetBuild(ctx, jobName, buildId)
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	if "SUCCESS" == data.GetResult() {
	//		fmt.Println("This build succeeded")
	//	}
	//}

	//beego.Info(jobs)
	//beego.Info(builds)

	//models.Job_build(jenkins, ctx, jobName)
	//models.Job_status(jenkins,ctx,jobName)

}

func (c *CicdController) GitlabApi() {
	c.TplName = "admin/cicd/jenkinsci-index.html"
}
