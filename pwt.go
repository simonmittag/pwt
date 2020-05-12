package pwt

import (
	"fmt"
	"net"
	"time"
)

const Version string = "v0.1.1"

var dialler = &net.Dialer{
	Timeout:   1 * time.Second,
	KeepAlive: 1 * time.Second,
}

func Zzz(host string, port uint16, timeSeconds int) bool {
	fmt.Printf("pwt waiting for %s:%d ", host, port)
	for i := 0; i < timeSeconds; i++ {

		conn, err := dialler.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			time.Sleep(time.Duration(1 * time.Second))
			fmt.Print(".")
		} else {
			if i > 0 {
				fmt.Printf(" connected in %d seconds\n", i)
			} else {
				fmt.Printf("connected in <1 second\n")
			}
			conn.Close()
			return true
		}
	}
	fmt.Printf(" aborted after %d seconds\n", timeSeconds)
	return false
}
