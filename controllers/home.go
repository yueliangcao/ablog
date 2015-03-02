package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/yueliangcao/ablog/models"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Prepare() {
	c.Layout = "_layout.tpl"
}

func (c *HomeController) Index() {
	var err error
	c.TplNames = "home_index.tpl"

	c.Data["articles"], err = models.GetHomeArticle(100, 1)
	if err != nil {
		beego.Warning(err.Error())
	}

	c.Data["active"] = "home"
	c.Data["title"] = "羊咩咩"
}

func (c *HomeController) Article() {
	var err error
	var article *models.Article

	c.TplNames = "home_article.tpl"
	id, _ := c.GetInt(":id")

	article, err = models.GetOneArticle(id)
	if err != nil {
		beego.Warning(err.Error())
	}

	err = models.IncrArticlePv(id)
	if err != nil {
		beego.Warning(err.Error())
	}

	c.Data["article"] = article
	c.Data["active"] = "home"
	c.Data["title"] = article.Title
}

func (c *HomeController) Tag() {
	var err error
	c.TplNames = "home_archives.tpl"
	name := c.GetString(":name")

	c.Data["archives"], c.Data["years"], err = models.GetArticleArchives(name)
	if err != nil {
		beego.Warning(err.Error())
	}

	c.Data["active"] = "home"
	c.Data["title"] = fmt.Sprintf("标签-%s", name)
}

func (c *HomeController) Archives() {
	var err error
	c.TplNames = "home_archives.tpl"

	c.Data["archives"], c.Data["years"], err = models.GetArticleArchives("")
	if err != nil {
		beego.Warning(err.Error())
	}

	c.Data["active"] = "archives"
	c.Data["title"] = "时间轴"
}

func (c *HomeController) Tags() {
	var err error
	c.TplNames = "home_tags.tpl"

	c.Data["tags"], err = models.GetAllTag()
	if err != nil {
		beego.Warning(err.Error())
	}

	c.Data["active"] = "tags"
	c.Data["title"] = "标签贴"
}

func (c *HomeController) AboutMe() {
	c.TplNames = "home_aboutme.tpl"

	c.Data["active"] = "aboutme"
	c.Data["title"] = "关于我"
}
