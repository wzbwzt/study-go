package model

import (
	"github.com/micro/go-micro/v2/logger"
)

func (m *Model) Greet(msg string) string {
	logger.Debug(msg)
	return msg
}
