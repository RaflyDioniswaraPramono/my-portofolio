package sosmed

type Sosmed struct {
	Id            int    `gorm:"primaryKey;AUTO_INCREMENT;column:id" json:"id"`
	SosmedName    string `gorm:"column:sosmed_name" json:"sosmed_name"`
	SosmedAccount string `gorm:"column:sosmed_account" json:"sosmed_account"`
	Url           string `gorm:"column:url" json:"url"`
}
