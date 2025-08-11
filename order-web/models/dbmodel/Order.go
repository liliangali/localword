package dbmodel

type Order struct {
	Id           uint         `gorm:"primarykey" json:"id"`
	Status       int          `gorm:"default:0;comment:status" json:"status"`
	Openid       string       `gorm:"type:varchar(255) not null;comment:openid" json:"openid"`
	Ordersn      string       `gorm:"type:varchar(255) not null;comment:ordersn" json:"ordersn"`
	Createtime   int64        `gorm:"default:0;comment:创建时间" json:"createtime"`
	Finishtime   int64        `gorm:"default:0;comment:创建时间" json:"finishtime"`
	Paytime      int64        `gorm:"default:0;comment:创建时间" json:"paytime"`
	Sendtime     int64        `gorm:"default:0;comment:创建时间" json:"sendtime"`
	Address      string       `gorm:"type:varchar(2250) not null;comment:地址" json:"address"`
	OrderGoods   []OrderGoods `gorm:"foreignKey:Orderid" json:"orderGoods"`
	Location     string       `gorm:"-" json:"location"` // 通过 struct 读写会忽略该字段
	Realname     string       `gorm:"-" json:"realname"`
	Mobile       string       `gorm:"-" json:"mobile"`
	ShortAddress string       `gorm:"-" json:"short_address"`
}

func (Order) TableName() string {
	return "ims_ewei_shop_order"
}

type OrderGoods struct {
	Id           uint    `gorm:"primarykey" json:"id"`
	Orderid      uint    `gorm:"default:0;comment:orderid" json:"orderid"`
	Goodsid      uint    `gorm:"default:0;comment:goodsid" json:"goodsid"`
	Goods        Goods   `gorm:"foreignKey:Goodsid" json:"Goods"`
	Total        int     `gorm:"default:0;comment:total" json:"total"`
	Oldprice     float64 `gorm:"type:decimal(10,2);comment:oldprice" json:"oldprice"`
	AveragePrice float64 `gorm:"-" json:"average_price"`
}

func (OrderGoods) TableName() string {
	return "ims_ewei_shop_order_goods"
}
