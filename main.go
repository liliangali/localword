package main

import (
	"fmt"
	"go.uber.org/zap"
	"localword/order-web/initialize"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//初始化usr_client
	//initialize.InitSrvConn()
	initialize.InitConfig()
	initialize.InitMysqlConn()
	//初始化router
	Router := initialize.Routers()
	//注册consul
	port := 8091

	//debug
	//text := "作为AI语言模型 🎈只有高中毕业的翔平，好不容易在一家条件不错的企业谋得职位。然而，休息时间他总被女上司随意差遣，心中虽有不满😤，却也只能一再忍耐。某天，他的手机里突然出现一个陌生应用程式，这让他获得了…… (注：这里为了满足字数和添加表情包的要求，对内容进行了一定程度的简化和调整。) "
	//
	//// 正则表达式匹配带括号且内部有超过15个字符的部分
	//re := regexp.MustCompile(`（[^）]{15,}）`)
	//
	//// 使用正则表达式删除匹配的部分
	//result := re.ReplaceAllString(text, "")
	//
	//fmt.Println("aaaaa", result)

	//if strings.Contains(text, "作为AI语言模型") {
	//	fmt.Println("333333")
	//}

	//register_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	//serviceId := global.ServerConfig.Name
	//err := register_client.Register(global.ServerConfig.Host, port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	//if err != nil {
	//	zap.S().Panic("注册失败", err.Error())
	//}

	/*
		1. S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2. 日志是分级别的，debug， info ， warn， error， fetal
		3. S函数和L函数很有用， 提供了一个全局的安全访问logger的途径
	*/
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
			zap.S().Panic("启动失败", err.Error())
		}
	}()

	//ticker := time.NewTicker(1 * time.Minute)
	//defer ticker.Stop() // 确保程序退出时停止Ticker
	//// 使用一个 goroutine 来处理每次触发的事件
	//go func() {
	//	for {
	//		select {
	//		case t := <-ticker.C:
	//			// 每次触发时执行的任务
	//			fmt.Println("Task executed at:", t)
	//			order.ResetVolClient()
	//			// 在这里添加你的任务逻辑
	//		}
	//	}
	//}()

	//定时任务
	//c := cron.New(cron.WithSeconds())
	//_, err = c.AddFunc("0 0 4 * * *", order.DoDayBill)
	//if err != nil {
	//	fmt.Println("Error scheduling task", err)
	//	return
	//}
	//c.Start()
	//定时任务

	//initialize.WatchReplaceFile()

	//接受终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
