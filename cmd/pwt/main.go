package main

import (
	"flag"
	"github.com/simonmittag/pwt"
)

func main() {
	port := flag.String("p", "8080", "the port, defaults to 8080")
	host := flag.String("h", "localhost", "the host name, defaults to localhost")
	time := flag.Int("t", 60, "time wait in seconds, defaults to 60")
}
