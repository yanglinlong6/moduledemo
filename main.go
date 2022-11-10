package main

import (
	"fmt"
	"moduledemo/common"
	"moduledemo/config"
	"moduledemo/mypackage" // 导入同一项目下的mypackage包
	"moduledemo/yanglinlong"
)

func main() {
	mypackage.New()
	hello := yanglinlong.YangHello(2)
	fmt.Println("main", hello)

	fmt.Println("杨龙龙")
	// 初始化nacos
	config.GetNacosConfig()
	// 初始化数据库(mysql)
	common.InitMysql()
	//host := "0.0.0.0"
	//port := 8899
	//
	//// 注册所有路由
	//r := routes.InitRoutes()
	//
	//srv := &http.Server{
	//	Addr:    fmt.Sprintf("%s:%d", host, port),
	//	Handler: r,
	//}
	//
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		fmt.Println("listen: %s\n", err)
	//	}
	//}()
	//
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	fmt.Println("Server forced to shutdown:", err)
	//}
	// 获取一个 gin 框架实例
	//r := gin.Default()
	// 配置路由对应
	//r.GET("/ping", func(ctx *gin.Context) {
	//	ctx.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//// 注册所有路由
	//r := routes.InitRoutes()
	// 服务启动
	//defer r.Run(":8899")

}
