package dbmodel

type Location struct {
	Id          uint   `gorm:"primarykey" json:"id"`
	Origin      string `gorm:"type:varchar(255) not null;comment:origin" json:"origin"`
	Destination string `gorm:"type:varchar(255) not null;comment:destination" json:"destination"`
	Distance    int    `gorm:"default:0;comment:distance" json:"distance"`
	Addtime     int64  `gorm:"default:0;comment:addtime" json:"addtime"`
	Strategy    int    `gorm:"default:0;comment:strategy" json:"strategy"`
}

func (Location) TableName() string {
	return "ims_ewei_shop_member_location"
}
