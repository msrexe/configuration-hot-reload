package main

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type LogConfig struct {
	Level string `json:"level"`
}

func (c *LogConfig) Update(newLogConfig LogConfig) {
	switch newLogConfig.Level {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	}
}

func (c *LogConfig) GetConfigAsJSON() []byte {
	jsonConfig, _ := json.Marshal(c)
	return jsonConfig
}
