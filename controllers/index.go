package controllers

import (
	"detect/utils"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	s := beego.AppConfig.String("sitespath")
	sits, e := ioutil.ReadFile(s)
	if e != nil {
		beego.Error("read sites error", e)
	}
	kpath := beego.AppConfig.String("keywordspath")
	keywords, e := ioutil.ReadFile(kpath)
	if e != nil {
		beego.Error("read keywords error", e)
	}
	c.Data["sites"] = string(sits)
	c.Data["keywords"] = string(keywords)
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	var code int
	var msg string
	defer utils.Retjson(c.Ctx, &msg, &code)
	input := c.Input()
	var sites string
	var keywords string
	for _, v := range input["sites"] {
		sites += v + "\n"
	}
	s := beego.AppConfig.String("sitespath")
	var site = []byte(sites)
	err := ioutil.WriteFile(s, site, 666)
	if err != nil {
		beego.Error("write  sites error", err)
		msg = "保存失败"
		return
	}

	for _, v := range input["keywords"] {
		keywords += v + "\n"
	}
	kpath := beego.AppConfig.String("keywordspath")
	var k = []byte(keywords)
	err1 := ioutil.WriteFile(kpath, k, 666)
	if err1 != nil {
		beego.Error("write  keywords error", err1)
		msg = "保存失败"
		return
	}
	c.Redirect(beego.URLFor("MainController.get"), 301)
}
