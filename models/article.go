package models

import (
	"database/sql"
	"errors"
)

// Article  对应一条文章数据
type Article struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
	Body  string `db:"body"`
}

func createArticleTable() {
	schema := `CREATE TABLE IF NOT EXISTS articles(
		id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		body longtext COLLATE utf8mb4_unicode_ci`

	DB.MustExec(schema)
}

// GetArticleByID 根据 id 获取 article
func GetArticleByID(id string) (article Article, err error) {
	article = Article{}
	err = DB.Get(&article, "SELECT * FROM articles WHERE id = ?", id)
	return article, err
}

// GetAllAricles 获取全部 article
func GetAllAricles() (articles []Article, err error) {
	articles = make([]Article, 0)
	err = DB.Select(&articles, "SELECT * FROM articles")
	return articles, err
}

// Create 创建 article
func (a *Article) Create() error {
	var (
		result sql.Result
		id     int64
		err    error
	)

	result, err = DB.NamedExec(`INSERT INTO articles (id, title, body) VALUES (:id, :title, :body)`, a)
	if err != nil {
		return err
	}

	if id, err = result.LastInsertId(); id > 0 {
		return nil
	}

	return errors.New("article 创建失败")
}

// Update 更新 article
func (a *Article) Update() error {
	var (
		result sql.Result
		id     int64
		err    error
	)

	if a.ID == 0 {
		return errors.New("ID not found")
	}

	result, err = DB.NamedExec(`UPDATE articles SET title = :title, body = :body WHERE id = :id`, a)
	if err != nil {
		return err
	}

	if id, err = result.RowsAffected(); id > 0 {
		return nil
	}

	return errors.New("article 更新失败")
}
