package article

import (
	"goblog/pkg/model"
	"goblog/pkg/types"
)

// Get 通过 ID 获取文章
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

// GetAll 获取全部文章
func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil

	// 上面代码相当于原生 sql 查询的
	/*
		1. 执行查询语句，返回一个结果集
		rows, err := db.Query("SELECT * from articles")
		logger.LogError(err)
		defer rows.Close()

		var articles []Article
		2. 循环读取结果
		for rows.Next() {
			var article Article
			2.1 扫码每一行的结果并赋值到一个 article 对象中
			err := rows.Scan(&article.ID, &article.Title, &article.Body)
			logger.LogError(err)
			2.2 将 article 追加到 articles 的这个数组中
			articles = append(articles, article)
		}

		2.3 检测遍历时是否发生错误
		err = rows.Err()
		logger.LogError(err)
	*/
}
