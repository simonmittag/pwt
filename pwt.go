package pwt

import (
	"fmt"
	"net"
	"os"
	"time"
)

const Version string = "v0.1.0"

func Zzz(host string, port uint16, timeSeconds int) {
	fmt.Printf("pwt[%s] attempting connection to %s:%d", Version, host, port)
	for i := 0; i < timeSeconds; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			time.Sleep(time.Duration(1 * time.Second))
			fmt.Print(".")
		} else {
			fmt.Printf("pwt[%s] %s:%d connected after %d seconds", Version, host, port, i)
			conn.Close()
			os.Exit(0)
		}
	}
	fmt.Printf("pwt[%s] %s:%d not connected after %d seconds, aborting...", Version, host, port, timeSeconds)
	os.Exit(-1)
}
