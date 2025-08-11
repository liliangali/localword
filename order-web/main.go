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
	//初始化router
	Router := initialize.Routers()
	//注册consul
	port := 8084
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
	//接受终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
