package admincontroller

import (
	"fmt"
	"gold/models"
	"io"
	"os"
	"time"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (a *ArticleController) Add() {
	categoryresult := models.CategoryAllList()
	clist := make([]map[string]interface{}, len(categoryresult))

	for k, v := range categoryresult {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		clist[k] = row
	}

	a.Data["categorys"] = clist
	a.TplName = "admin/article/add.html"
}

func (a *ArticleController) List() {

	page, _ := a.GetInt("page")
	code, _ := a.GetInt("code")

	// model 查询
	pagesize := 10
	categoryresult := models.CategoryAllList()
	clist := make(map[int]string)
	result, total := models.List(page, pagesize)
	list := make([]map[string]interface{}, 10)

	for _, v := range categoryresult {

		clist[v.Id] = v.Name
	}

	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["title"] = v.Title
		row["sub_title"] = v.Sub_title
		row["cid"] = v.Cid
		row["status"] = v.Status
		row["index"] = v.Index
		row["image"] = v.Image
		row["c_name"] = clist[v.Cid]
		list[k] = row
	}

	pagenum := total/int64(pagesize) + 1
	pag := make(map[int]int)
	for index := 1; index <= int(pagenum); index++ {
		pag[index] = index
	}

	a.Data["list"] = list
	a.Data["count"] = pag
	a.Data["count1"] = pagenum
	a.Data["code"] = code
	a.Data["page"] = page
	a.TplName = "admin/article/list.html"
}

func (a *ArticleController) Edit() {
	id, _ := a.GetInt("id")

	res, _ := models.ArticleById(id)
	categoryresult := models.CategoryAllList()
	clist := make([]map[string]interface{}, len(categoryresult))

	for k, v := range categoryresult {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		clist[k] = row
	}

	a.Data["categorys"] = clist
	a.Data["data"] = res
	a.TplName = "admin/article/edit.html"
}

func (a *ArticleController) EditOne() {
	var img string
	id, _ := a.GetInt("id")
	title := a.GetString("title")
	stitle := a.GetString("sub_title")
	cid, _ := a.GetInt("cid")
	status, _ := a.GetInt("status")
	index, _ := a.GetInt("index")
	content := a.GetString("content")

	article, _ := models.ArticleById(id)
	article.Title = title
	article.Sub_title = stitle
	article.Cid = cid
	article.Status = status
	article.Index = index
	article.Content = content
	img = article.Image

	// 获取上传文件
	f, h, err := a.GetFile("filename")
	if err == nil {
		defer f.Close()
		path := CreateDateDir()
		a.SaveToFile("filename", path+h.Filename)
		img = "/" +  path + h.Filename
	}

	article.Image = img

	if err := article.Update(); err != nil {
		a.Ctx.Redirect(302, "/admin/article/list?page=1&code=1")
	}

	a.Ctx.Redirect(302, "/admin/article/list?page=1&code=2")
}

func (a *ArticleController) AddOne() {
	var img string
	title := a.GetString("title")
	stitle := a.GetString("sub_title")
	cid, _ := a.GetInt("cid")
	status, _ := a.GetInt("status")
	index, _ := a.GetInt("index")
	content := a.GetString("content")

	// 获取上传文件
	f, h, err := a.GetFile("filename")
	if err == nil {
		defer f.Close()
		path := CreateDateDir()
		a.SaveToFile("filename", path+h.Filename)
		img = "/" +  path + h.Filename
	}

	// 准备article指针 插入数据数据
	article := new(models.Article)
	article.Title = title
	article.Sub_title = stitle
	article.Cid = cid
	article.Status = status
	article.Index = index
	article.Content = content
	article.Image = img
	article.Created = time.Now().Format("2006-01-02 15:04:05")
	article.Updated = time.Now().Format("2006-01-02 15:04:05")

	id, err := models.Addarticle(article)
	if err != nil || id == 0 {
		a.Ctx.Redirect(302, "/admin/article/list?page=1&code=1")
	}
	a.Ctx.Redirect(302, "/admin/article/list?page=1&code=2")
}

// 单个文件上传
func (i *ArticleController) Uploadone() {
	f, h, err := i.GetFile("filename")
	if err != nil {
		fmt.Print("3333")
	}
	defer f.Close()

	path := CreateDateDir()

	i.SaveToFile("filename", path+h.Filename)
	mystruct := make(map[string]interface{}, 0)
	var arr [1]string
	arr[0] = "/" +  path + h.Filename

	mystruct["errno"] = 0
	mystruct["data"] = arr
	i.Data["json"] = &mystruct
	i.ServeJSON()
}

// 多文件上传
func (a *ArticleController) Upload() {
	var img [5]string
	mystruct := make(map[string]interface{}, 0)
	mystruct["errno"] = 0
	path := CreateDateDir()
	files, err := a.GetFiles("filename")
	if err != nil {
		mystruct["msg"] = "44444"
		return
	}
	for i, _ := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			mystruct["msg"] = "3333"
		}
		//create destination file making sure the path is writeable.
		fname := path + files[i].Filename
		dst, err := os.Create(fname)
		defer dst.Close()
		if err != nil {
			mystruct["msg"] = "11111111"
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			mystruct["msg"] = "222"
		}
		img[i] =  "/" + fname
	}
	mystruct["data"] = img
	a.Data["json"] = &mystruct
	a.ServeJSON()
}

// 创建文件夹
func CreateDateDir() string {
	folderName := time.Now().Format("2006-01-02")
	folderPath := "static/upload/" + string(folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	return folderPath + "/"
}
