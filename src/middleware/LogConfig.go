package middleware

import "git.sxxfuture.net/filfi/letsfil/letsfil-job/src/configs"

type LogConfig struct {
	Level string `json:"level"`
}

func NewLogConfig() *LogConfig {
	return &LogConfig{Level: configs.Get().Log.Level}
}
