package gogetdockerimage

import (
	"testing"

	"github.com/EricNeid/go-getdockerimage/internal/verify"
)

func TestParseDestination(t *testing.T) {
	// action
	res, err := ParseDestination("ssh://user:pass@10.20.300.400:22/home/user/dir")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "user", res.User)
	verify.Equals(t, "pass", res.Pass)
	verify.Equals(t, "10.20.300.400:22", res.Addr)
}

func TestParseDestination_noPasswordSet(t *testing.T) {
	// action
	res, err := ParseDestination("ssh://user@10.20.300.400:22/home/user/dir")
	// verify
	verify.Ok(t, err)
	verify.Equals(t, "user", res.User)
	verify.Equals(t, "", res.Pass)
	verify.Equals(t, "10.20.300.400:22", res.Addr)
}
