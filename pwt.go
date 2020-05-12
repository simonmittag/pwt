package pwt

import (
	"fmt"
	"net"
	"os"
	"time"
	"github.com/rs/zerolog/log"
)

const Version string = "v0.1.0"

func Zzz(host string, port uint16, timeSeconds int) {
	for i:=0;i<timeSeconds;i++{
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		defer conn.Close()
		if err != nil {
			time.Sleep(time.Duration(1*time.Second))
		} else {
			log.Info().Msgf("pwt[%s] %s:%d connected after %d seconds", Version, host, port, i)
			os.Exit(0)
		}
	}
	log.Info().Msgf("pwt[%s] %s:%d not connected after %d seconds, aborting...", Version, host, port, timeSeconds)
	os.Exit(-1)
}