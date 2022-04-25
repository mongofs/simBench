package simBench

import (
	"fmt"

	"github.com/gorilla/websocket"
	"go.uber.org/atomic"
	"math/rand"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().Unix()))
)

type Bench struct {

	// metric count
	success atomic.Int32
	fail    atomic.Int32
	online  atomic.Int32
	retry   atomic.Int32

	// message count
	singleMessageCount  atomic.Int64
	allUserMessageCount atomic.Int64

	oToken string //output token
	config *Config

	closeMonitor chan string
}

func NewBench(conf *Config)*Bench{
	return &Bench{
		success:             atomic.Int32{},
		fail:                atomic.Int32{},
		online:              atomic.Int32{},
		retry:               atomic.Int32{},
		singleMessageCount:  atomic.Int64{},
		allUserMessageCount: atomic.Int64{},
		oToken:              "",
		config:              conf,
		closeMonitor:        make(chan string,10),
	}
}

func (s *Bench) Run() {
	tokens := s.getValidateKey()
	for k, v := range tokens {
		if k%s.config.Concurrency == 0 && k != 0 {
			time.Sleep(1 * time.Second)
			if s.success.Load() != 0 {
				fmt.Printf("Current number of successfully established connections %v ,fail %v \n\r", s.success.Load(), s.fail.Load())
			}
		}
		go s.CreateClient(s.config.Host, v, s.config.TagNum, k)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("Current number of successfully established connections %v ,fail %v \n\r", s.success.Load(), s.fail.Load())
	s.monitor()
}

var limiter = time.NewTicker(50 * time.Microsecond)

func (s *Bench) CreateClient(Host, token string, tagN, id int) error {
	<-limiter.C
	dialer := websocket.Dialer{}
	url := fmt.Sprintf(Host+"?token=%s&tag=%s", token, s.createTags(tagN))
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		s.fail.Inc()
		fmt.Printf("error occurs during runtime id : %v, url : %s ,err :%s\r\n", id, url, err.Error())
		return nil
	}
	s.success.Inc()
	s.online.Inc()
	defer conn.Close()
	go func() {
		for {
			time.Sleep( 10 * time.Second)
			conn.WriteJSON("{test: 1}")
		}
	}()
	for {
		messageType, messageData, err := conn.ReadMessage()
		if nil != err {
			s.closeMonitor <- token
			fmt.Printf("connction occure err : %v\r\n", err)
		}
		s.allUserMessageCount.Inc()
		if token == s.oToken {
			s.singleMessageCount.Inc()
			switch messageType {
			case websocket.TextMessage:
				fmt.Printf(" recieve the message content : %v , \r\n", string(messageData))
			case websocket.BinaryMessage:
			default:
			}
		}
	}
}

func (s *Bench) monitor() {
	go func() {
		t := time.NewTicker(time.Duration(5) * time.Second)
		for {
			select {
			case <-t.C:
				fmt.Printf("Current Open connections %v ,Retry connections %v, message count  %v ,all message count %v \n\r",
					s.online.Load(), s.retry.Load(),s.singleMessageCount.Load(),s.allUserMessageCount.Load())
			case token := <-s.closeMonitor:
				fmt.Printf("client %v is closed \r\n", token)
				s.retry.Inc()
				s.online.Dec()
				go s.CreateClient(s.config.Host, token, s.config.TagNum, 99999)
			}
		}
	}()
}

func (s *Bench) getValidateKey() []string {
	var tokens []string
	for i := 0; i < s.config.Number; i++ {
		tokens = append(tokens, RandString(20))
	}
	s.oToken = tokens[len(tokens)/2]
	return tokens
}

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
