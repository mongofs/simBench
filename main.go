package main

import (
	"time"
)

func main() {
	config := InitConfig()
	RunServer(config)
	time.Sleep(time.Duration(config.keepTime) * time.Second)
}


// RunServer 启动服务
func RunServer(cof *config) {
	sb := sBench{config: cof}
	sb.Run()
}
