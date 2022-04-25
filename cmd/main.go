package main

import (
	"fmt"
	"github/mongofs/simBench"
	"time"
)

func main() {
	config := simBench.InitConfig()
	RunServer(config)
	if config.KeepTime == 0 {
		select {}
	}else {
		time.Sleep(time.Duration(config.KeepTime) * time.Second)
	}
	fmt.Println("exit process")
}


// RunServer 启动服务
func RunServer(cof *simBench.Config) {
	sb := simBench.NewBench(cof)
	sb.Run()
}
