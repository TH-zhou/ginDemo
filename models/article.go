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

//func GetArticles(pageNum, pageSize int, maps interface{}) []Article {
//	var articles []Article
//	DBEngine
//}

