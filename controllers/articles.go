package controllers

import (
	"database/sql"
	"fmt"
	"goblog/models"
	"goblog/utils"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"unicode/utf8"

	"github.com/gorilla/mux"
)

// Articles controllers
type Articles struct {
	Router *mux.Router
}

// ArticlesFormData 创建博文表单数据
type ArticlesFormData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

func (*Articles) validateArticleFormData(title string, body string) map[string]string {
	errors := make(map[string]string)
	// 验证标题
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度需介于 3-40"
	}

	// 验证内容
	if body == "" {
		errors["body"] = "内容不能为空"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "内容长度需大于或等于 10 个字节"
	}

	return errors
}

func (*Articles) modelError(err error, w http.ResponseWriter) {
	utils.CheckError(err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "500 服务器内部错误")
}

func (a *Articles) checkModelError(err error, w http.ResponseWriter) {
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 文章未找到")
	} else {
		a.modelError(err, w)
	}
}

func (a *Articles) getArticleFromRoute(w http.ResponseWriter, r *http.Request) (article *models.Article, id int64, err error) {
	sid := utils.GetRouteVariable("id", r)
	iid, err := strconv.Atoi(sid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 文章未找到")
		return nil, 0, err
	}
	id = int64(iid)

	article, err = models.GetArticleByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
			return nil, 0, err
		}

		a.modelError(err, w)
		return nil, 0, err
	}

	return article, id, err
}

// Show article 详情
func (a *Articles) Show(w http.ResponseWriter, r *http.Request) {
	article, _, err := a.getArticleFromRoute(w, r)
	if err != nil {
		return
	}

	// 4. 读取成功，显示文章
	tmpl, err := template.New("show.gohtml").
		Funcs(template.FuncMap{
			"RouteName2URL": utils.RouteName2URL,
			"Int64ToString": utils.Int64ToString,
		}).
		ParseFiles("resources/views/articles/show.gohtml")
	utils.CheckError(err)

	tmpl.Execute(w, article)
}

// Index article 列表
func (*Articles) Index(w http.ResponseWriter, r *http.Request) {
	articles, err := models.GetAllAricles()
	utils.CheckError(err)

	tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
	utils.CheckError(err)
	tmpl.Execute(w, articles)
}

// Create article 创建表单
func (a *Articles) Create(w http.ResponseWriter, r *http.Request) {
	storeURL, _ := a.Router.Get("articles.store").URL()
	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
	utils.CheckError(err)
	tmpl.Execute(w, ArticlesFormData{
		Title:  "",
		Body:   "",
		URL:    storeURL,
		Errors: nil,
	})
}

// Store article 创建
func (a *Articles) Store(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	// 验证标题/内容
	errors := a.validateArticleFormData(title, body)

	// 检查是否有错误
	if len(errors) == 0 {
		article := &models.Article{
			Title: title,
			Body:  body,
		}
		newID, err := article.Create()
		if err == nil {
			fmt.Fprint(w, "插入成功，ID 为"+strconv.FormatInt(newID, 10))
		} else {
			a.modelError(err, w)
		}
	} else {
		storeURL, _ := a.Router.Get("articles.store").URL()

		tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
		utils.CheckError(err)
		tmpl.Execute(w, ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    storeURL,
			Errors: errors,
		})
	}
}

// Edit article 编辑表单
func (a *Articles) Edit(w http.ResponseWriter, r *http.Request) {
	article, id, err := a.getArticleFromRoute(w, r)
	if err != nil {
		return
	}

	updateURL, _ := a.Router.Get("articles.update").URL("id", strconv.Itoa(int(id)))
	tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
	utils.CheckError(err)
	tmpl.Execute(w, ArticlesFormData{
		Title:  article.Title,
		Body:   article.Body,
		URL:    updateURL,
		Errors: nil,
	})
}

// Update article 编辑
func (a *Articles) Update(w http.ResponseWriter, r *http.Request) {
	_, id, err := a.getArticleFromRoute(w, r)
	if err != nil {
		return
	}

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")
	errors := a.validateArticleFormData(title, body)

	if len(errors) == 0 {
		article := &models.Article{
			Title: title,
			Body:  body,
		}
		article.ID = id

		if article.Update() != nil {
			a.modelError(err, w)
			return
		}

		showURL, _ := a.Router.Get("articles.show").URL("id", strconv.Itoa(int(id)))
		http.Redirect(w, r, showURL.String(), http.StatusFound)
	} else {
		updateURL, _ := a.Router.Get("articles.update").URL("id", strconv.Itoa(int(id)))
		tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
		utils.CheckError(err)
		tmpl.Execute(w, ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    updateURL,
			Errors: errors,
		})
	}
}

// Delete article 删除
func (a *Articles) Delete(w http.ResponseWriter, r *http.Request) {
	article, _, err := a.getArticleFromRoute(w, r)
	if err != nil {
		return
	}

	if article.Delete() != nil {
		a.modelError(err, w)
	} else {
		indexURL, _ := a.Router.Get("articles.index").URL()
		http.Redirect(w, r, indexURL.String(), http.StatusFound)
	}
}
