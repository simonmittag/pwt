package main

import (
	"flag"
	"fmt"
	"github.com/simonmittag/pwt"
	"os"
	"strconv"
	"strings"
)

func main() {
	host := "localhost"
	port := 80
	timeSeconds := flag.Int("w", 10, "time wait in seconds")
	modeVersion := flag.Bool("v", false, "print pwt version")
	flag.Parse()

	args := flag.Args()
	if len(args) == 1 {
		if strings.Contains(args[0], ":") {
			ph := strings.Split(args[0], ":")
			host = ph[0]
			port, _ = strconv.Atoi(ph[1])
		} else {
			host = args[0]
		}
	}

	if *modeVersion {
		fmt.Printf("pwt[%s]\n", pwt.Version)
	} else {
		ok := pwt.Zzz(host, uint16(port), *timeSeconds)
		if ok {
			os.Exit(0)
		} else {
			os.Exit(-1)
		}
	}
}
