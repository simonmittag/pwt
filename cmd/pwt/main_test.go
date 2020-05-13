package main

import (
	"testing"
)

func TestParseArgsHostNameColonPort(t *testing.T) {
	host, port := parseArgs([]string{"somehost:8083"})
	if host != "somehost" {
		t.Errorf("did not resolve host want somehost, got %s", host)
	}
	if port != 8083 {
		t.Errorf("did not resolve port want 8083 got %d", port)
	}
}

func TestParseArgsHostName(t *testing.T) {
	host, port := parseArgs([]string{"myhost.com"})
	if host != "myhost.com" {
		t.Errorf("did not resolve host want myhost.com got %s", host)
	}
	if port != default_port {
		t.Errorf("did not resolve port want %d, got %d", default_port, port)
	}
}

func TestParseArgsIpv4(t *testing.T) {
	host, port := parseArgs([]string{"127.0.0.1"})
	if host != "127.0.0.1" {
		t.Errorf("did not resolve host want 127.0.0.1 got %s", host)
	}
	if port != default_port {
		t.Errorf("did not resolve port want %d, got %d", default_port, port)
	}
}

func TestParseArgsIpv4ColonPort(t *testing.T) {
	host, port := parseArgs([]string{"127.0.0.1:8081"})
	if host != "127.0.0.1" {
		t.Errorf("did not resolve host want 127.0.0.1 got %s", host)
	}
	if port != 8081 {
		t.Errorf("did not resolve port want %d, got %d", 8081, port)
	}
}

func TestParseArgsIpv6(t *testing.T) {
	host, port := parseArgs([]string{"[::1]"})
	if host != "[::1]" {
		t.Errorf("did not resolve host want [::1] got %s", host)
	}
	if port != default_port {
		t.Errorf("did not resolve port want %d, got %d", default_port, port)
	}
}

func TestParseArgsIpv6ColonPort(t *testing.T) {
	host, port := parseArgs([]string{"[::1]:8087"})
	if host != "[::1]" {
		t.Errorf("did not resolve host want [::1] got %s", host)
	}
	if port != 8087 {
		t.Errorf("did not resolve port want %d, got %d", 8087, port)
	}
}

func TestParseArgsHostAndPort(t *testing.T) {
	host, port := parseArgs([]string{"hostname.com", "8089"})
	if host != "hostname.com" {
		t.Errorf("did not resolve host want [::1] got %s", host)
	}
	if port != 8089 {
		t.Errorf("did not resolve port want %d, got %d", 8087, port)
	}
}
