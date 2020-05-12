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
		_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			time.Sleep(time.Duration(1*time.Second))
		} else {
			log.Info().Msgf("pwt %s:%d connected after %i seconds", host, port, i)
			os.Exit(0)
		}
	}
	log.Info().Msgf("pwt %s:%d not connected after %i seconds, aborting...", host, port, timeSeconds)
	os.Exit(-1)
}