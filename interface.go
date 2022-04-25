package simBench



// 监控

type  Monitor interface {
	Online () int // 查看在线用户

}

type Iface interface {

	// 创建num 个用户，tag1
	CreateConn (num ,tag int, )

}