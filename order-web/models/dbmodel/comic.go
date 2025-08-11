package dbmodel

type Volcano struct {
	Id         uint   `gorm:"primarykey" json:"id"`
	Volkey     string `gorm:"type:varchar(1000) not null;default:'';comment:volkey" json:"volkey"`
	Volcontent string `gorm:"type:varchar(2250) not null;default:'';comment:volcontent" json:"volcontent"`
	Voltype    int    `gorm:"default:0;comment:voltype" json:"voltype"`
}

func (Volcano) TableName() string {
	return "volcano"
}
