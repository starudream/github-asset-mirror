package config

import (
	"os"
	"testing"

	"github.com/starudream/go-lib/testx"
)

func TestMain(m *testing.M) {
	Default()

	os.Exit(m.Run())
}

func TestFormatURL(t *testing.T) {
	data := map[string]any{
		"ver":      "v1.6.0",
		"os":       "linux",
		"arch":     "amd64",
		"platform": "client",
	}
	v, err := FormatURL(data)
	testx.P(t, err, v)
}
