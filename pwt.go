package pwt

import (
	"fmt"
	"net"
	"os"
	"time"
)

const Version string = "v0.1.0"

func Zzz(host string, port uint16, timeSeconds int) {
	fmt.Printf("pwt dialling %s:%d ", Version, host, port)
	for i := 0; i < timeSeconds; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			time.Sleep(time.Duration(1 * time.Second))
			fmt.Print(".")
		} else {
			fmt.Printf(" success", Version, host, port)
			if i>0 {
				fmt.Printf(" in %d seconds\n", i)
			} else {
				fmt.Printf(" in <1 second\n")
			}
			conn.Close()
			os.Exit(0)
		}
	}
	fmt.Printf(" aborted after %d seconds\n", Version, host, port, timeSeconds)
	os.Exit(-1)
}
