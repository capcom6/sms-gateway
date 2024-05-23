package health

import (
	httpclient "net/http"
	"time"

	"github.com/capcom6/go-infra-fx/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func testHealth(shutdowner fx.Shutdowner, logger *zap.Logger, config http.Config) {
	client := httpclient.Client{
		Timeout: 1 * time.Second,
	}

	_, err := client.Get("http://" + config.Listen + "/health")
	if err != nil {
		if err := shutdowner.Shutdown(fx.ExitCode(1)); err != nil {
			logger.Error("Failed to shutdown", zap.Error(err))
		}
		return
	}

	if err := shutdowner.Shutdown(); err != nil {
		logger.Error("Failed to shutdown", zap.Error(err))
	}
}
