package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type SLBLoader struct {
	interval int
	needStop chan bool
}

func (s *SLBLoader) run() {
	println("begin to get SLB from dns url")
	ips, _ := net.LookupIP("t50ngj89fi0qo.oceanbase.aliyuncs.com")
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
		}
	}
}

func NewSLBLoader(interval int) *SLBLoader {
	s := new(SLBLoader)
	s.interval = interval
	s.needStop = make(chan bool)
	return s
}

func (s *SLBLoader) start() {
	ticker := time.NewTicker(time.Duration(s.interval) * time.Millisecond)
	go func() {
		for {
			select {
			case <-s.needStop:
				ticker.Stop()
				println("stop timer task")
				return
			case <-ticker.C:
				s.run()
			}
		}
	}()
}

func (s *SLBLoader) stop() {
	s.needStop <- true
}

func main() {
	println("hello world")
	s := NewSLBLoader(5000)
	//s := nil
	//s.start()
	//time.Sleep(20 * 1000 * time.Millisecond)
	//println("begin to stop timer task")
	//s.stop()
	//println("finish to stop timer task")
	//time.Sleep(10 * 1000 * time.Millisecond)
	//println("finish main")
	ss := fmt.Sprintf("%#v", s)
	io.WriteString(os.Stdout, ss)
}
