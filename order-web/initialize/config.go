package initialize

import (
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"localword/order-web/global"
	"log"
	"os"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("DBAPI_DEBUG")
	configFilePrefix := "redis"
	configFileName := fmt.Sprintf("order-web/%s.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("order-web/%s-debug.yaml", configFilePrefix)
	}

	fmt.Println(configFileName)
	v := viper.New()
	v.SetConfigFile(configFileName)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//fmt.Println(v.Get("name"))
	//NacosConfig *config.NacosConfig = &config.NacosConfig{}
	//NacosConfigs *config.NacosConfig = config.NacosConfig{}

	if err := v.Unmarshal(global.RedisConfig); err != nil {
		panic(err)
	}

	//global.OpenApiClient = arkruntime.NewClientWithApiKey(global.RedisConfig.Chatgptkey)

	//global.PassList = ReadPass()

	//defaultConfig := openai.DefaultConfig(global.RedisConfig.Chatgptkey)
	//defaultConfig.BaseURL = global.RedisConfig.ChatgptUrl
	//global.OpenApiClient = openai.NewClientWithConfig(defaultConfig)
	//global.OpenApiClient = openai.NewClient(global.RedisConfig.Chatgptkey)

	//global.ExcelDir = global.RedisConfig.Exceldir

	//config_t2s := "order-web/opccdata/config/t2s.json"
	//global.Opeccc = opencc.NewConverter(config_t2s)
	//fmt.Println(global.PassList)
	//global.ReplaceSentence = ReplaceSentence()
}

func Iptxt() []string {
	fp, err := os.Open("order-web/ipb.txt")
	if err != nil {
		fmt.Println(err) //打开文件错误
		return nil
	}
	buf := bufio.NewScanner(fp)
	var passList []string
	for {
		if !buf.Scan() {
			break //文件读完了,退出for
		}
		line := buf.Text() //获取每一行
		passList = append(passList, line)
	}
	return passList
}

func WatchReplaceFile() {
	// 创建一个新的监视器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	// 启动一个 goroutine 以处理文件系统事件
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Printf("Event received: %s\n", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Printf("File modified: %s\n", event.Name)
					global.IpData = Iptxt()
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Printf("File created: %s\n", event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					fmt.Printf("File removed: %s\n", event.Name)
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					fmt.Printf("File renamed: %s\n", event.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Printf("Error received: %s\n", err)
			}
		}
	}()
	// 添加要监控的文件或目录
	filename := "order-web/ipb.txt" // 替换为你想监控的文件
	err = watcher.Add(filename)
	if err != nil {
		log.Fatal(err)
	}
}

func ReplaceSentence() []string {
	fp, err := os.Open("order-web/replace.txt")
	if err != nil {
		fmt.Println(err) //打开文件错误
		return nil
	}
	buf := bufio.NewScanner(fp)
	var passList []string
	for {
		if !buf.Scan() {
			break //文件读完了,退出for
		}
		line := buf.Text() //获取每一行
		passList = append(passList, line)
	}
	return passList
}

func ReadPass() []string {
	fp, err := os.Open("order-web/pass.txt")
	if err != nil {
		fmt.Println(err) //打开文件错误
		return nil
	}
	buf := bufio.NewScanner(fp)
	var passList []string
	for {
		if !buf.Scan() {
			break //文件读完了,退出for
		}
		line := buf.Text() //获取每一行
		passList = append(passList, line)
	}
	return passList
}
