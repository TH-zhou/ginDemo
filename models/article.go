package models

type Article struct {
	Id         int    `xorm:"pk autoincr" json:"id"`
	TagId      int    `xorm:"index" json:"tag_id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedOn  int    `json:"created_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedOn int    `json:"modified_on"`
	ModifiedBy string `json:"modified_by"`
	DeletedOn  int    `json:"deleted_on"`
	State      int    `json:"state"`
}


// 根据条件获取所有文章
func GetArticles(pageNum, pageSize int, maps interface{}) []Article {
	var articles []Article
	DBEngine.Where(maps).Limit(pageSize, pageNum).Find(&articles)

	return articles
}

// 根据条件获取文章总数
func GetArticleToTal(maps interface{}) (count int) {
	DBEngine.Table(new(Article)).Where(maps).Count(&count)

	return
}

// 通过id获取单篇文章
func GetArticleById(id int) (article *Article) {
	DBEngine.ID(id).Get(article)

	return
}

// 添加文章
func AddArticle(data map[string]interface{}) bool {
	article := &Article{
		TagId: data["tag_id"].(int),
	}
	if articleId, err := DBEngine.Insert(article); articleId > 0 && err == nil {
		return true
	}

	return false
}

func EditArticle(id int, maps interface{}) bool {
	return false
}

