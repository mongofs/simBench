package main


import (
	"fmt"
	"github.com/gorilla/websocket"
	"math/rand"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().Unix()))
)

type sBench struct {
	config *config

}


func (s *sBench)Run (){
	tokens := s.getValidateKey()
	counter := 0
	for k, v := range tokens {
		if k%s.config.concurrency == 0 {
			counter += s.config.concurrency
			fmt.Printf("%v connections established successfully ,Total current connections is %v  \n\r",s.config.concurrency,counter)
			time.Sleep(1 * time.Second)
		}
		go s.CreateClient(s.config.host, v,s.config.tagNum ,k )
	}
}




// RandString 生成随机字符串做Token, 注意的是这里生成token的规则，
// 需要你能够在validate的接口实现中自己能解出来
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}


func  (s *sBench)getValidateKey()[]string  {
	var tokens []string
	for i := 0; i < s.config.number; i++ {
		tokens = append(tokens, RandString(20))
	}
	return tokens
}




// CreateMockClient 图形界面化也可以使用这个网站进行查看 http://www.baidu.com/conn?token=1080&version=v.10
// 模拟连接，在此包内可
func (s *sBench)CreateClient(Host, token string, tagN ,id int ) error {
	dialer := websocket.Dialer{}
	url := fmt.Sprintf(Host+"?token=%s&tag=%s", token, s.createTags(tagN))
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		fmt.Printf("error occurs during runtime id : %v, url : %s ,err :%s\r\n",id ,url,err.Error())
		return nil
	}
	defer conn.Close()
	for {
		messageType, messageData, err := conn.ReadMessage()
		if nil != err {
			return err
		}
		switch messageType {
		case websocket.TextMessage:
			fmt.Printf("recieve the message content : %v \r\n", string(messageData))
		case websocket.BinaryMessage:
		default:
		}
	}
}