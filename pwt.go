package pwt

import (
	"fmt"
	"net"
	"os"
	"time"
)

const Version string = "v0.1.0"

func Zzz(host string, port uint16, timeSeconds int) {
	fmt.Printf("pwt[%s] attempting to dial %s on %d ", Version, host, port)
	for i := 0; i < timeSeconds; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			time.Sleep(time.Duration(1 * time.Second))
			fmt.Print(".")
		} else {
			fmt.Printf("\npwt[%s] succesfully connected to %s:%d after %d seconds\n", Version, host, port, i)
			conn.Close()
			os.Exit(0)
		}
	}
	fmt.Printf("\npwt[%s] aborting attempted dial to %s on %d after %d seconds, exiting\n", Version, host, port, timeSeconds)
	os.Exit(-1)
}
