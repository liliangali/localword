package dbmodel

type Goods struct {
	Id    int    `gorm:"primarykey" json:"id"`
	Title string `gorm:"type:varchar(255) not null;comment:title" json:"title"`
}

func (Goods) TableName() string {
	return "ims_ewei_shop_goods"
}
