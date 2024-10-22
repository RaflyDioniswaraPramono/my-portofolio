package content

type Content struct {
	Id          int    `gorm:"primaryKey;AUTO_INCREMENT;not null;colum:id" json:"id"`
	ContentText string `gorm:"colum:content_text" json:"content_text"`
	LinkPath    string `gorm:"colum:link_path" json:"link_path"`
}
