package gogetdockerimage

import (
	"testing"

	"github.com/EricNeid/go-getdockerimage/internal/verify"
)

func TestParseDestination(t *testing.T) {
	// action
	_, err := ParseDestination("ssh://user@10.20.300.400:22/home/user/dir")
	// verify
	verify.Ok(t, err)
}
