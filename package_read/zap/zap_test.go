package test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"go.uber.org/zap"
)

var logger *zap.Logger

func logmsg() {
	logger, _ = zap.NewProduction()
}

func GetWeb(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}

}

func TestZap(t *testing.T) {
	path := "./root"
	s, err := os.Stat(path)
	fmt.Println(s, err)
	logmsg()
	defer logger.Sync()
	url := "www.baidu.com"
	GetWeb(url)
}
