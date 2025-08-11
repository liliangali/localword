package config

type RedisConfig struct {
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Password    string `mapstructure:"password" json:"password"`
	DBid        int    `mapstructure:"dbid" json:"dbid"`
	Expire      int    `mapstructure:"expire" json:"expire"`
	RandExpire  int    `mapstructure:"randexpire" json:"randexpire"`
	Sqlhost     string `mapstructure:"sqlhost" json:"sqlhost"`
	Sqluser     string `mapstructure:"sqluser" json:"sqluser"`
	Sqlpassword string `mapstructure:"sqlpassword" json:"sqlpassword"`
	Sqlport     int    `mapstructure:"sqlport" json:"sqlport"`
	Sqldbname   string `mapstructure:"sqldbname" json:"sqldbname"`
	Exceldir    string `mapstructure:"exceldir" json:"exceldir"`
	Maxgpt      int    `mapstructure:"maxgpt" json:"maxgpt"`
	Chatgptkey  string `mapstructure:"chatgptkey" json:"chatgptkey"`
	Model       string `mapstructure:"model" json:"model"`
	Maxcartoon  int    `mapstructure:"maxcartoon" json:"maxcartoon"`
}
