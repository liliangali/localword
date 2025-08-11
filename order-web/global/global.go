package global

import (
	"github.com/go-ego/gse/hmm/idf"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"gorm.io/gorm"
	"localword/order-web/config"
)

var (
	RedisConfig     *config.RedisConfig = &config.RedisConfig{}
	PassList        []string
	GlobalDB        *gorm.DB
	OpenApiClient   *arkruntime.Client
	ExcelDir        string
	GseTe           idf.TagExtracter
	ReplaceSentence []string
	IpData          []string
	VolErrorStatus  bool
)
