package utils

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

var logger *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"log.log"}
	logger, _ = config.Build()
	defer logger.Sync()
}

func LoggerMiddleware() gin.HandlerFunc {
	return ginzap.Ginzap(logger, time.RFC3339, true)
}

func RecoverMiddleware() gin.HandlerFunc {
	return ginzap.RecoveryWithZap(logger, true)
}
