package pwt

import (
	"testing"
)

func TestZzzPass(t *testing.T) {
	ok := Zzz("google.com", 80, 1)
	if !ok {
		t.Error("google.com should return ok <1s but didn't")
	}
}

func TestZzzFail(t *testing.T) {
	ok := Zzz("thisisn'treallyadomainandcan'tpasspwt.com", 80, 3)
	if ok {
		t.Error("fake domain should't return ok <3s but did")
	}
}