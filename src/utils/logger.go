package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func SetLoggerOptions(useColor bool, verbose bool) {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:          !useColor,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	log.SetOutput(os.Stderr)

	if verbose == true {
		log.SetLevel(log.DebugLevel)
	}
}

var (
	LogDebug   = log.Debugf
	LogWarning = log.Warnf
	LogInfo    = log.Infof
	LogError   = log.Errorf
	LogPanic   = log.Panicf
	LogTrace   = log.Tracef
)
