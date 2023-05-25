package main

import (
	"os"
	"testing"
	"time"
)

//TODO: cannot run these tests because multiple invocations of flag.Arg()
//func TestMainFuncWithVersion(t *testing.T) {
//	os.Args = append([]string{""}, "-v")
//	main()
//}
//
//func TestMainFuncWithUsage(t *testing.T) {
//	os.Args = append([]string{""}, "-h")
//	main()
//}

func TestMainFuncWithPositionalArgsBeforeFlagArgs(t *testing.T) {
	//first arg is command
	before := time.Now()
	os.Args = []string{"pwt", "www.google.com:888", "-w", "4"}
	main()

	after := time.Now()
	got := after.Sub(before).Seconds()
	want := time.Duration(time.Second * 4).Seconds()
	if float64(got) < float64(want) {
		t.Errorf("wanted at least %v seconds, but got %v", want, got)
	}

}

func TestParseArgsHostNameColonPort(t *testing.T) {
	host, port, err := parseArgs([]string{"somehost:8083"})
	if host != "somehost" {
		t.Errorf("did not resolve host want somehost, got %s", host)
	}
	if port != 8083 {
		t.Errorf("did not resolve port want 8083 got %d", port)
	}
	if err != nil {
		t.Errorf("should not have errored, but got %v", err)
	}
}

func TestParseArgsHostName(t *testing.T) {
	host, port, err := parseArgs([]string{"myhost.com"})
	if host != "myhost.com" {
		t.Errorf("did not resolve host want myhost.com got %s", host)
	}
	if port != default_port {
		t.Errorf("did not resolve port want %d, got %d", default_port, port)
	}
	if err != nil {
		t.Errorf("should not have errored, but got %v", err)
	}
}

func TestParseArgsIpv4(t *testing.T) {
	host, port, err := parseArgs([]string{"127.0.0.1"})
	if host != "127.0.0.1" {
		t.Errorf("did not resolve host want 127.0.0.1 got %s", host)
	}
	if port != default_port {
		t.Errorf("did not resolve port want %d, got %d", default_port, port)
	}
	if err != nil {
		t.Errorf("should not have errored, but got %v", err)
	}
}

func TestParseArgsIpv4ColonPort(t *testing.T) {
	host, port, err := parseArgs([]string{"127.0.0.1:8081"})
	if host != "127.0.0.1" {
		t.Errorf("did not resolve host want 127.0.0.1 got %s", host)
	}
	if port != 8081 {
		t.Errorf("did not resolve port want %d, got %d", 8081, port)
	}
	if err != nil {
		t.Errorf("should not have errored, but got %v", err)
	}
}

func TestParseArgsIpv6(t *testing.T) {
	host, port, err := parseArgs([]string{"[::1]"})
	if host != "[::1]" {
		t.Errorf("did not resolve host want [::1] got %s", host)
	}
	if port != default_port {
		t.Errorf("did not resolve port want %d, got %d", default_port, port)
	}
	if err != nil {
		t.Errorf("should not have errored, but got %v", err)
	}
}

func TestParseArgsIpv6ColonPort(t *testing.T) {
	host, port, err := parseArgs([]string{"[::1]:8087"})
	if host != "[::1]" {
		t.Errorf("did not resolve host want [::1] got %s", host)
	}
	if port != 8087 {
		t.Errorf("did not resolve port want %d, got %d", 8087, port)
	}
	if err != nil {
		t.Errorf("should not have errored, but got %v", err)
	}
}

func TestParseArgsHostAndPort(t *testing.T) {
	_, _, err := parseArgs([]string{"hostname.com", "8089"})
	if err == nil {
		t.Errorf("should have failed %v", err)
	}
}

func TestWait(t *testing.T) {
	wait("google.com", 80, 1)
}

func TestPrintVersion(t *testing.T) {
	printVersion()
}
