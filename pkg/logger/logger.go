package logger

import (
	"fmt"
	"log"
	"os"
)

type CustomLogger struct {
	logger *log.Logger
}

type LogMessage struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

func New() *CustomLogger {
	return &CustomLogger{
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}
}

func (c *CustomLogger) Debugln(message string) {
	c.logger.Printf(fmt.Sprintf("DEBUG: %s", message))
}

func (c *CustomLogger) Infoln(message string) {
	c.logger.Printf(fmt.Sprintf("INFO: %s", message))
}

func (c *CustomLogger) Errorln(message string) {
	c.logger.Printf(fmt.Sprintf("ERROR: %s", message))
}
