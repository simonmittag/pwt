package main

import (
	"flag"
	"fmt"
	"github.com/simonmittag/pwt"
	"os"
	"strconv"
	"strings"
)

const default_host = "localhost"
const default_port = 80

func main() {

	timeSeconds := flag.Int("w", 10, "time wait in seconds")
	modeVersion := flag.Bool("v", false, "print pwt version")
	flag.Parse()
	host, port := parseArgs(flag.Args())

	if *modeVersion {
		printVersion()
	} else {
		wait(host, port, *timeSeconds)
	}
}

func printVersion() {
	fmt.Printf("pwt[%s]\n", pwt.Version)
}

func wait(host string, port int, timeSeconds int) {
	ok := pwt.Zzz(host, uint16(port), timeSeconds)
	if ok {
		os.Exit(0)
	} else {
		os.Exit(-1)
	}
}

func parseArgs(args []string) (string, int) {
	var host string = default_host
	var port int = default_port

	if len(args) == 1 {
		dest := args[0] //because of Ipv6
		if (strings.Contains(dest, ":") && !strings.Contains(dest, "::")) || strings.Contains(dest, "]:") {
			ci := strings.LastIndex(args[0], ":")
			host = args[0][0:ci]
			port, _ = strconv.Atoi(args[0][ci+1:])
		} else {
			host = args[0]
		}
	} else if len(args) == 2 {
		host = args[0]
		port, _ = strconv.Atoi(args[1])
	}
	return host, port
}
