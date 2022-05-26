package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	_ "github.com/Astronaut-X-X/TaskList/back_end/model"
	_ "github.com/Astronaut-X-X/TaskList/back_end/router"
)

func main() {
	fmt.Println("Start service ...")
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	log.Println("Server exiting")
}
