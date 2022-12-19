package bootstrap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"mindmap-go/utils/config"
	"time"
)

func DebugTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func NewLogger(cfg *config.Config) *zap.Logger {
	var l *zap.Logger
	var err error
	if cfg.App.Production {
		l, err = zap.NewProduction()
	} else {
		conf := zap.NewDevelopmentConfig()
		conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		conf.EncoderConfig.EncodeTime = DebugTimeEncoder
		l, err = conf.Build()
		zap.ReplaceGlobals(l)
	}
	if err != nil {
		panic(err)
	}
	return l
}
