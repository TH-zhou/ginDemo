package models

type Tag struct {
	Id         int    `xorm:"pk autoincr" json:"id"`
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 根据条件获取所有tag
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	DBEngine.Where(maps).Limit(pageSize, pageNum).Find(&tags)

	return
}

// 获取tag总数
func GetTagTotal(maps interface{}) (count int64) {
	tag := new(Tag)
	count, _ = DBEngine.Where(maps).Count(tag)

	return
}

// 判断tag是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	DBEngine.Where("name = ?", name).Get(&tag)
	if tag.Id > 0 {
		return true
	}

	return false
}

// 添加标签
func AddTag(name string, state int, createdBy string) bool {
	tag := &Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	}
	//tag := new(Tag)
	//tag.Name = name
	//tag.CreatedBy = createdBy
	//tag.State = state
	tagID, err := DBEngine.Insert(tag)
	if err != nil {
		return false
	}
	if tagID > 0 {
		return true
	}

	return false
}

// 通过id判断tag是否存在
func ExistTagById(id int) bool {
	if isExist, err := DBEngine.Exist(&Tag{Id: id}); isExist && err == nil {
		return true
	}
	return false
}

// 编辑标签
func EditTag(id int, maps interface{}) bool {
	if affected, err := DBEngine.Table(new(Tag)).ID(id).Update(maps); affected > 0 && err == nil {
		return true
	}

	return false
}

// 删除标签
func DeleteTag(id int) bool {
	affected, err := DBEngine.ID(id).Delete(&Tag{})
	if affected > 0 && err == nil {
		return true
	}

	return false
}
