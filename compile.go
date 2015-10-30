package main

import (
	"log"
	"os"
	"path"

	"chca/conf"
	"chca/utils"
	"neutron/template"
)

var data = map[string]interface{}{
	"title":       conf.SiteTitle(),
	"subtitle":    conf.SiteSubTitle(),
	"description": conf.SiteDescription(),
	"author":      conf.Author(),
	"avatar":      conf.Avatar(),
	"github":      conf.Github(),
	"weibo":       conf.Weibo(),
}

func Compile() {
	checkFile()
	copy()

	LoadArticle()
	// 创建页面
    CompileHome()
    CompileArticle()
    CompileArchive()
	CompileTagPage()
	CompileCatePage()
	CompileCategory()
	CompileTag()
}

// 编译主页
func CompileHome() {

	data["artlist"] = GetAllArt()
	data["cate"] = GetCate()
	data["tpl"] = conf.DirTheme() + "/layout/index.html"

	err := utils.MkDir(conf.DirHtml())

	if err != nil {
		panic(err)
	}

	homepath := path.Join(conf.DirHtml(), "index.html")

	htmlfile, err := os.Create(homepath)
	if err != nil {
		panic(err)
	}

	t, _ := template.New(conf.DirTheme() + "/layout/main.html")
	t = t.Funcs(template.FuncMap{"unescaped": utils.Unescaped, "cmonth": utils.CMonth, "format": utils.Format, "count": utils.Count, "lt": utils.Lt, "gt": utils.Gt, "eq": utils.Eq})
	t.Walk(conf.DirTheme()+`/layout`, ".html")
	t.Execute(htmlfile, data)
}

// 编译文章页
func CompileArticle() {
	artlist := GetAllArt()

	for _, art := range artlist {
		data["tpl"] = conf.DirTheme() + "/layout/page.html"

		data["article"] = art
		data["cate"] = GetCate()

		url := CreatePostLink(art)
		filepath := path.Join(conf.DirHtml(), url)

		err := utils.MkDir(filepath)

		if err != nil {
			panic(err)
		}

		filename := path.Join(filepath, "index.html")

		htmlfile, err := os.Create(filename)

		if err != nil {
			panic(err)
		}

		t, _ := template.New(conf.DirTheme() + "/layout/main.html")
		t = t.Funcs(template.FuncMap{"unescaped": utils.Unescaped, "cmonth": utils.CMonth, "format": utils.Format, "count": utils.Count, "lt": utils.Lt, "gt": utils.Gt, "eq": utils.Eq})
		t.Walk(conf.DirTheme()+`/layout`, ".html")
		t.Execute(htmlfile, data)
	}
}

// 编译归档页
func CompileArchive() {

	data["archive"] = GetArchive()
	data["cate"] = GetCate()
	data["tpl"] = conf.DirTheme() + "/layout/archive.html"

	filepath := path.Join(conf.DirHtml(), "archive")

	err := utils.MkDir(filepath)

	if err != nil {
		panic(err)
	}

	filename := path.Join(filepath, "index.html")

	htmlfile, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	t, _ := template.New(conf.DirTheme() + "/layout/main.html")
	t = t.Funcs(template.FuncMap{"unescaped": utils.Unescaped, "cmonth": utils.CMonth, "format": utils.Format, "count": utils.Count, "lt": utils.Lt, "gt": utils.Gt, "eq": utils.Eq})
	t.Walk(conf.DirTheme()+`/layout`, ".html")
	t.Execute(htmlfile, data)
}

// 编译cate导航页
func CompileCatePage() {

	data["cate"] = GetCate()
	data["tpl"] = conf.DirTheme() + "/layout/category.html"

	filepath := path.Join(conf.DirHtml(), "category")

	err := utils.MkDir(filepath)

	if err != nil {
		panic(err)
	}

	filename := path.Join(filepath, "index.html")

	htmlfile, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	t, _ := template.New(conf.DirTheme() + "/layout/main.html")
	t = t.Funcs(template.FuncMap{"unescaped": utils.Unescaped, "cmonth": utils.CMonth, "format": utils.Format, "count": utils.Count, "lt": utils.Lt, "gt": utils.Gt, "eq": utils.Eq})
	t.Walk(conf.DirTheme()+`/layout`, ".html")
	t.Execute(htmlfile, data)
}

// 编译category页面
func CompileCategory() {

	cates := GetCate()
	data["cate"] = cates

	for _, cate := range cates {

		data["ptitle"] = cate.Name
		data["content"] = cate.Posts
		data["count"] = cate.Count
		data["tpl"] = conf.DirTheme() + "/layout/post.html"

		filepath := path.Join(conf.DirHtml(), "category", cate.Name)

		err := utils.MkDir(filepath)

		if err != nil {
			panic(err)
		}

		filename := path.Join(filepath, "index.html")

		htmlfile, err := os.Create(filename)

		if err != nil {
			panic(err)
		}

		t, _ := template.New(conf.DirTheme() + "/layout/main.html")
		t = t.Funcs(template.FuncMap{"unescaped": utils.Unescaped, "cmonth": utils.CMonth, "format": utils.Format, "count": utils.Count, "lt": utils.Lt, "gt": utils.Gt, "eq": utils.Eq})
		t.Walk(conf.DirTheme()+`/layout`, ".html")
		t.Execute(htmlfile, data)
	}

}

// 编译tag导航页
func CompileTagPage() {

	data["cate"] = GetCate()
	data["tag"] = GetTag()
	data["tpl"] = conf.DirTheme() + "/layout/tag.html"

	filepath := path.Join(conf.DirHtml(), "tag")

	err := utils.MkDir(filepath)

	if err != nil {
		panic(err)
	}

	filename := path.Join(filepath, "index.html")

	htmlfile, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	t, _ := template.New(conf.DirTheme() + "/layout/main.html")
	t = t.Funcs(template.FuncMap{"unescaped": utils.Unescaped, "cmonth": utils.CMonth, "format": utils.Format, "count": utils.Count, "lt": utils.Lt, "gt": utils.Gt, "eq": utils.Eq})
	t.Walk(conf.DirTheme()+`/layout`, ".html")
	t.Execute(htmlfile, data)
}

// 编译tag页面
func CompileTag() {

	tags := GetTag()
	data["cate"] = GetCate()

	for _, tag := range tags {

		data["ptitle"] = tag.Name
		data["content"] = tag.Posts
		data["count"] = tag.Count
		data["tpl"] = conf.DirTheme() + "/layout/post.html"

		filepath := path.Join(conf.DirHtml(), "tag", tag.Name)

		err := utils.MkDir(filepath)

		if err != nil {
			panic(err)
		}

		filename := path.Join(filepath, "index.html")

		htmlfile, err := os.Create(filename)

		if err != nil {
			panic(err)
		}

		t, _ := template.New(conf.DirTheme() + "/layout/main.html")
		t = t.Funcs(template.FuncMap{"unescaped": utils.Unescaped, "cmonth": utils.CMonth, "format": utils.Format, "count": utils.Count, "lt": utils.Lt, "gt": utils.Gt, "eq": utils.Eq})
		t.Walk(conf.DirTheme()+`/layout`, ".html")
		t.Execute(htmlfile, data)
	}

}

func CrearteMark(filename string) string {
	file := path.Join(conf.DirMark(), filename+".md")

	_, err := os.Stat(file)
	if !os.IsNotExist(err) {
		log.Println("已存在文件")
		os.Exit(1)
	}

	src, err := utils.CreateFile(conf.DirMark(), filename+".md")
	if err != nil {
		panic(err)
	}

	return src
}

func copy() {

	// copy 配置文件
	/*_, err := utils.CopyFile("conf.ini", path.Join(conf.DirHtml(), "conf.ini"))
	  if err != nil {
	      panic(err)
	  }*/

	err := utils.CopyDir(path.Join(conf.DirTheme(), "assets"), path.Join(conf.DirHtml(), "assets"))
	if err != nil {
		panic(err)
	}

}

func checkFile() {
	if _, err := os.Stat(conf.DirTheme()); os.IsNotExist(err) {
		panic("需要先初始化并添加模板文件")
	}

	if _, err := os.Stat(conf.DirStor()); os.IsNotExist(err) {
		panic("需要先初始化")
	}
}
