package main

import (
	"github.com/spf13/pflag"
	"time"
)

func main() {
	config := InitConfig()
	pflag.Parse()
	RunServer(config)
	time.Sleep(time.Duration(config.keepTime) * time.Second)
}


// RunServer 启动服务
func RunServer(cof *config) {
	sb := sBench{config: cof}
	sb.Run()
}
