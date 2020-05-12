package pwt

import (
	"fmt"
	"net"
	"time"
)

func pwt(host string, port int) {
	for {
		_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			time.Sleep(time.Duration(1*time.Second))
		} else {
			break
		}
	}
}