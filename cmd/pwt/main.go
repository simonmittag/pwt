package main

import (
	"flag"
	"github.com/simonmittag/pwt"
	"os"
)

func main() {
	host := flag.String("t", "localhost", "the destination host name")
	port := flag.Uint("p", 80, "the destination port")
	timeSeconds := flag.Int("w", 10, "time wait in seconds")
	flag.Parse()

	ok := pwt.Zzz(*host, uint16(*port), *timeSeconds)
	if ok {
		os.Exit(0)
	} else {
		os.Exit(-1)
	}
}
