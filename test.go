package main

import (
	"test/logger"
)

func main() {
	var log logger.Logger
	log.Init(logger.LevelDebug, "log/debug.log")
	log.Debug("debug str")
	log.Info("info str")
	log.Warning("warn str")
	log.Error("error str")
	log.Deinit()

	log.Init(logger.LevelInfo, "log/info.log")
	log.Debug("debug str")
	log.Info("info str")
	log.Warning("warn str")
	log.Error("error str")
	log.Deinit()

	log.Init(logger.LevelWarning, "log/warning.log")
	log.Debug("debug str")
	log.Info("info str")
	log.Warning("warn str")
	log.Error("error str")
	log.Deinit()

	log.Init(logger.LevelError, "log/error.log")
	log.Debug("debug str")
	log.Info("info str")
	log.Warning("warn str")
	log.Error("error str")
	log.Deinit()
}
