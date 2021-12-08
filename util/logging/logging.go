package logging

import (
	"os"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

var lock = &sync.Mutex{}
var logging *logrus.Logger

func GetLogger(filename string) *logrus.Logger {
	lock.Lock()
	defer lock.Unlock()
	file, err := os.OpenFile("/log_serotonin/"+filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("%s", err)
	}

	if logging == nil {
		logging = logrus.New()
	}
	logging.SetOutput(file)
	logging.SetFormatter(&logrus.JSONFormatter{})
	return logging
}
