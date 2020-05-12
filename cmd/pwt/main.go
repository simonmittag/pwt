package main

import (
	"flag"
	"github.com/simonmittag/pwt"
)

func main() {
	host := flag.String("t", "localhost", "the destination host name")
	port := flag.Uint("p", 8080, "the destination port")
	timeSeconds := flag.Int("w", 60, "time wait in seconds")
	flag.Parse()


	pwt.Zzz(*host, uint16(*port), *timeSeconds)
}
