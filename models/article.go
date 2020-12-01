package models

import (
	"database/sql"
	"errors"
	"goblog/app"
	"goblog/utils"
	"strconv"
)

// Article  对应一条文章数据
type Article struct {
	BaseModel
	Title string `db:"title"`
	Body  string `db:"body"`
}

// GetArticleByID 根据 id 获取 article
func GetArticleByID(id int64) (*Article, error) {
	article := Article{}
	err := app.DB().Get(&article, "SELECT * FROM articles WHERE id = ?", id)
	return &article, err
}

// GetAllAricles 获取全部 article
func GetAllAricles() (articles []Article, err error) {
	articles = make([]Article, 0)
	err = app.DB().Select(&articles, "SELECT * FROM articles")
	return articles, err
}

// Create 创建 article
func (a *Article) Create() (int64, error) {
	var (
		result sql.Result
		id     int64
		err    error
	)

	result, err = app.DB().NamedExec(`INSERT INTO articles (title, body) VALUES (:title, :body)`, a)
	if err != nil {
		return 0, err
	}

	if id, err = result.LastInsertId(); id > 0 {
		return id, nil
	}

	return 0, errors.New("article 创建失败")
}

// Update 更新 article
func (a *Article) Update() error {
	var (
		result sql.Result
		count  int64
		err    error
	)

	if a.ID == 0 {
		return errors.New("ID not found")
	}

	result, err = app.DB().NamedExec(`UPDATE articles SET title = :title, body = :body WHERE id = :id`, a)
	if err != nil {
		return err
	}

	if count, err = result.RowsAffected(); count > 0 {
		return nil
	}

	return errors.New("article 更新失败")
}

// Delete 删除 article
func (a *Article) Delete() error {
	var (
		result sql.Result
		count  int64
		err    error
	)

	if a.ID == 0 {
		return errors.New("ID not found")
	}

	result, err = app.DB().NamedExec(`DELETE FROM articles WHERE id = :id`, a.ID)
	if err != nil {
		return err
	}

	if count, err = result.RowsAffected(); count > 0 {
		return nil
	}

	return errors.New("article 删除失败")
}

// Link article link
func (a *Article) Link() string {
	showURL, err := app.Router().Get("articles.show").URL("id", strconv.FormatInt(a.ID, 10))
	if err != nil {
		utils.CheckError(err)
		return ""
	}
	return showURL.String()
}
