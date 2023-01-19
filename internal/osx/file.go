package osx

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/starudream/go-lib/log"

	"github.com/starudream/github-asset-mirror/internal/unitx"
)

var cli = &http.Client{Timeout: time.Minute}

func SaveFile(url string, filepath string) ([]byte, error) {
	resp, err := cli.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if sl := resp.Header.Get("Content-Length"); sl != "" {
		il, ce := strconv.Atoi(sl)
		if ce == nil {
			log.Info().Msgf("download file size: %s", unitx.HumanSize(float64(il)))
		}
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, os.WriteFile(filepath, bs, 0644)
}

func ProxyURL(proxy, url string) string {
	switch proxy {
	case "ghproxy":
		return "https://ghproxy.com/" + url
	case "fastgit":
		return strings.ReplaceAll(url, "https://github.com/", "https://download.fastgit.org/")
	}
	return url
}

func ExistFile(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && !fi.IsDir()
}
