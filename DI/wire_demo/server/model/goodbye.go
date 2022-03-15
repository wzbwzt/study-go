package model

import (
	"github.com/micro/go-micro/v2/logger"
)

func (m *Model) Goodbye(msg string) string {
	logger.Debug(msg)
	return msg
}
