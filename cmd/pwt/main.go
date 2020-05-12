package main

import (
	"flag"
	"github.com/simonmittag/pwt"
)

func main() {
	host := flag.String("t", "localhost", "the target host name, defaults to localhost")
	port := flag.Uint("p", 8080, "the port, defaults to 8080")
	timeSeconds := flag.Int("w", 60, "time wait in seconds, defaults to 60")
	flag.Parse()


	pwt.Zzz(*host, uint16(*port), *timeSeconds)
}
