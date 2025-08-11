package dbmodel

type SeoWord struct {
	Id          uint   `gorm:"primarykey" json:"id"`
	Title       string `gorm:"type:varchar(255) not null;comment:title" json:"title"`
	ExtendTitle string `gorm:"type:varchar(255) not null;comment:extend_title" json:"extend_title"`
}

func (SeoWord) TableName() string {
	return "seo_words"
}

type Address struct {
	Id       uint   `gorm:"primarykey" json:"id"`
	Openid   string `gorm:"type:varchar(255) not null;comment:openid" json:"openid"`
	Realname string `gorm:"type:varchar(255) not null;comment:realname" json:"realname"`
	Mobile   string `gorm:"type:varchar(255) not null;comment:mobile" json:"mobile"`
	Province string `gorm:"type:varchar(255) not null;comment:province" json:"province"`
	City     string `gorm:"type:varchar(255) not null;comment:city" json:"city"`
	Area     string `gorm:"type:varchar(255) not null;comment:area" json:"area"`
	Address  string `gorm:"type:varchar(255) not null;comment:address" json:"address"`
	Lng      string `gorm:"type:varchar(255) not null;comment:lng" json:"lng"`
	Lat      string `gorm:"type:varchar(255) not null;comment:lat" json:"lat"`
}

func (Address) TableName() string {
	return "ims_ewei_shop_member_address"
}

type ExcelData struct {
	Id       uint   `gorm:"primarykey" json:"id"`
	Content1 string `gorm:"type:text not null;comment:content1" json:"content1"`
	Content2 string `gorm:"type:text not null;comment:content2" json:"content2"`
	Ftype    int    `gorm:"default:0;comment:ftype" ftype:"total"`
}

func (ExcelData) TableName() string {
	return "exceldata"
}

type Volsetting struct {
	Id           uint   `gorm:"primarykey" json:"id"`
	Volkey       string `gorm:"type:varchar(500) not null;comment:volkey" json:"volkey"`
	Volmodel     string `gorm:"type:varchar(2250) not null;comment:volmodel" json:"volmodel"`
	Loginaccount string `gorm:"type:varchar(2250) not null;comment:loginaccount" json:"loginaccount"`
	IsLose       int    `gorm:"default:0;comment:is_lose" json:"is_lose"`
}

func (Volsetting) TableName() string {
	return "volsetting"
}
