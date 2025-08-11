package dbmodel

type Assign struct {
	Id           uint   `gorm:"primarykey" json:"id"`
	StartTime    int64  `gorm:"default:0;comment:start_time" json:"start_time"`
	EndTime      int64  `gorm:"default:0;comment:createtime" json:"end_time"`
	AddTime      int64  `gorm:"default:0;comment:add_time" json:"add_time"`
	NormalCarMax int    `gorm:"default:0;comment:normal_car_max" json:"normal_car_max"`
	BigCarMax    int    `gorm:"default:0;comment:big_car_max" json:"big_car_max"`
	NormalCarMin int    `gorm:"default:0;comment:normal_car_min" json:"normal_car_min"`
	BigCarBound  int    `gorm:"default:0;comment:big_car_bound" json:"big_car_bound"`
	BigCarShop   int    `gorm:"default:0;comment:big_car_shop" json:"big_car_shop"`
	HopCar       int    `gorm:"default:0;comment:hop_car" json:"hop_car"`
	Status       int    `gorm:"default:0;comment:0待生成 1已生成" json:"status"`
	Excel        string `gorm:"type:varchar(255) not null;comment:excel" json:"excel"`
	NoOrdersn    string `gorm:"type:text not null;comment:no_ordersn" json:"no_ordersn"`
	LineText     string `gorm:"type:text not null;comment:line_text" json:"line_text"`
	AllOrdersn   string `gorm:"type:text not null;comment:all_ordersn" json:"all_ordersn"`
	Total        int    `gorm:"default:0;comment:total" json:"total"`
	OrderNum     int    `gorm:"default:0;comment:order_num" json:"order_num"`
	ShopLocation string `gorm:"type:text not null;comment:line_text" json:"shop_location"`
	ShopNum      int    `gorm:"default:0;comment:order_num" json:"shop_num"`
	ProcureExcel string `gorm:"type:varchar(255) not null;comment:procure_excel" json:"procure_excel"`
	CheckExcel   string `gorm:"type:varchar(255) not null;comment:check_excel" json:"check_excel"`
}

func (Assign) TableName() string {
	return "ims_ewei_shop_order_assign"
}
