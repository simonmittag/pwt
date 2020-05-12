package main

import (
	"flag"
	"github.com/simonmittag/pwt"
)

func main() {
	host := flag.String("t", "localhost", "the destination host name")
	port := flag.Uint("p", 80, "the destination port")
	timeSeconds := flag.Int("w", 10, "time wait in seconds")
	flag.Parse()

	pwt.Zzz(*host, uint16(*port), *timeSeconds)
}
