package model

type News struct {
	Model
	Show       int    `json:"show,omitempty"  form:"show"`
	Type       int    `json:"type,omitempty" form:"type"`
	SeeNum     int    `json:"see_num,omitempty"  gorm:"column:see_num" form:"see_num"`
	Hot        int    `json:"hot,omitempty" form:"hot"`
	Zan        int    `json:"zan,omitempty" form:"zan"`
	CateId     int    `json:"cate_id,omitempty"  gorm:"column:cate_id" form:"cateId"`
	Tags       string `json:"tags,omitempty" form:"tags"`
	Content    string `json:"content,omitempty" form:"content"`
	SeoKeyword string `json:"seo_keyword,omitempty"  gorm:"column:seo_keyword" form:"seo_keyword"`
	Title      string `json:"title,omitempty" form:"title"`
	Thumb      string `json:"thumb,omitempty" form:"thumb"`
	Desc       string `json:"desc,omitempty" form:"desc"`
}

func NewNews() *News {
	return &News{}
}

func (n News) TableName() string {
	return "ssf_news"
}

// GetById 通过ID获取内容
func (n *News) GetById(id string) (*News, int64) {
	res := db.Where("id=?", id).First(&n)
	return n, res.RowsAffected
}

func (n *News) AddNews() int64 {
	res := db.Create(n).RowsAffected
	return res
}
