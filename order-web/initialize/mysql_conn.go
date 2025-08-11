package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"localword/order-web/global"
)

func InitMysqlConn() {
	//dsn := "root:@tcp(127.0.0.1:3306)/gorm_class?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.RedisConfig.Sqluser, global.RedisConfig.Sqlpassword, global.RedisConfig.Sqlhost,
		global.RedisConfig.Sqlport, global.RedisConfig.Sqldbname,
	)
	zap.S().Infof("读取mysql配置：%v", dsn)

	fmt.Println(dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256, //默认字符串的长度
	}), &gorm.Config{
		//Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "ims_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: false,  // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		DisableForeignKeyConstraintWhenMigrating: true, //禁用 外键约束
	})
	if err != nil {
		panic(err)
	}
	//GLOBAL_DB = dbmodel
	global.GlobalDB = db
	//order.InitModel()
	//TestDogCreate()
	//CreateTest()
	//One2One()
	//AutoMigrateTable()
}

func AutoMigrateTable() {
	global.GlobalDB.AutoMigrate(
	//&dbmodel.Menu{},
	//&dbmodel.Role{},
	)
}
