package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/simonmittag/pwt"
	"os"
	"strconv"
	"strings"
)

const default_port = 80

type Mode uint8

const (
	Server Mode = 1 << iota
	Version
	Usage
)

func main() {
	mode := Server
	host := ""
	port := default_port

	timeSeconds := flag.Int("w", 10, "time wait in seconds")
	v := flag.Bool("v", false, "print pwt version")
	h := flag.Bool("h", false, "print usage instructions")
	flag.Usage = printUsage
	err := ParseFlags()
	if err != nil {
		mode = Usage
	} else {
		a := flag.Args()
		host, port, err = parseArgs(a)

		if *v {
			mode = Version
		} else if err != nil || *h {
			mode = Usage
		} else {
			mode = Server
		}
	}

	switch mode {
	case Server:
		wait(host, port, *timeSeconds)
	case Usage:
		printUsage()
	case Version:
		printVersion()
	}
}

func printUsage() {
	printVersion()
	fmt.Printf("Usage: pwt [-h]|[-v]|[-w n] host[:port]\n")
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Printf("pwt[%s]\n", pwt.Version)
}

func wait(host string, port int, timeSeconds int) {
	ok := pwt.Zzz(host, uint16(port), timeSeconds)
	if !ok {
		os.Exit(-1)
	}
}

func parseArgs(args []string) (string, int, error) {
	var host = ""
	var port int = default_port

	if len(args) == 1 {
		dest := args[0] //because of Ipv6
		if (strings.Contains(dest, ":") && !strings.Contains(dest, "::")) || strings.Contains(dest, "]:") {
			ci := strings.LastIndex(args[0], ":")
			host = args[0][0:ci]
			p, _ := strconv.Atoi(args[0][ci+1:])
			if p > 0 && p < 65535 {
				port = p
			}
		} else {
			host = dest
		}
		return host, port, nil
	} else {
		return "", 0, errors.New("invalid host or port")
	}
}

// ParseFlags parses the command line args, allowing flags to be
// specified after positional args.
func ParseFlags() error {
	return ParseFlagSet(flag.CommandLine, os.Args[1:])
}

// ParseFlagSet works like flagset.Parse(), except positional arguments are not
// required to come after flag arguments.
func ParseFlagSet(flagset *flag.FlagSet, args []string) error {
	var positionalArgs []string
	for {
		if err := flagset.Parse(args); err != nil {
			return err
		}
		// Consume all the flags that were parsed as flags.
		args = args[len(args)-flagset.NArg():]
		if len(args) == 0 {
			break
		}
		// There's at least one flag remaining and it must be a positional arg since
		// we consumed all args that were parsed as flags. Consume just the first
		// one, and retry parsing, since subsequent args may be flags.
		positionalArgs = append(positionalArgs, args[0])
		args = args[1:]
	}
	// Parse just the positional args so that flagset.Args()/flagset.NArgs()
	// return the expected value.
	// Note: This should never return an error.
	return flagset.Parse(positionalArgs)
}
