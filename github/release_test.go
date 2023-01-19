package github

import (
	"testing"
	"time"

	"github.com/starudream/go-lib/testx"
)

func TestGetLatestRelease(t *testing.T) {
	xLatestReleaseExpire = time.Second / 2

	resp1, err := GetLatestRelease()
	testx.P(t, err, resp1)

	resp2, err := GetLatestRelease()
	testx.P(t, err, resp2)

	time.Sleep(time.Second)

	resp3, err := GetLatestRelease()
	testx.P(t, err, resp3)
}
