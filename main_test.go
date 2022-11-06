package main

import (
	"os"
	"test/logger"
	"testing"
)

func TestInitInvalidArgument(t *testing.T) {
	var log logger.Logger
	if err := log.Init(4, "", logger.OutputTypeNormal); err != logger.ErrorInvalidArgument {
		t.Errorf("Test case[Init() Invalid argument] err:%d want:%d", err, logger.ErrorInvalidArgument)
	}

	if err := log.Init(logger.LevelDebug, "", 2); err != logger.ErrorInvalidArgument {
		t.Errorf("Test case[Init() Invalid argument] err:%d want:%d", err, logger.ErrorInvalidArgument)
	}
}

func TestInitOpenError(t *testing.T) {
	var log logger.Logger
	if err := log.Init(logger.LevelDebug, "/log/debug.log", logger.OutputTypeNormal); err != logger.ErrorLogFileOpenFailed {
		t.Errorf("Test case[Init() Open Error] err:%d want:%d", err, logger.ErrorLogFileOpenFailed)
	}
}

func TestInitDefault(t *testing.T) {
	os.Remove("./info.log")

	var log logger.Logger
	if err := log.Init(logger.LevelDebug, "", logger.OutputTypeNormal); err != logger.ErrorNone {
		t.Errorf("Test case[Init() Default] err:%d want:%d", err, logger.ErrorNone)
	}

	if _, err := os.Stat("./info.log"); err != nil {
		t.Errorf("Test case[Init() Default] info.log not found.")
	}

	log.Deinit()
	os.Remove("./info.log")
}

func TestInitSpecify(t *testing.T) {
	os.Remove("log/test.log")

	var log logger.Logger
	if err := log.Init(logger.LevelDebug, "log/test.log", logger.OutputTypeNormal); err != logger.ErrorNone {
		t.Errorf("Test case[Init() Specify] err:%d want:%d", err, logger.ErrorNone)
	}

	if _, err := os.Stat("log/test.log"); err != nil {
		t.Errorf("Test case[Init() Specify] log/test.log not found.")
	}

	log.Deinit()
	os.Remove("log/test.log")
}

func TestDebug(t *testing.T) {
	os.Remove("log/debug.log")

	var log logger.Logger
	log.Init(logger.LevelDebug, "log/debug.log", logger.OutputTypeNormal)

	if err := log.Debug("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Debug() Debug] err:%d want:%d", err, logger.ErrorNone)
	}

	if err := log.Info("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Debug() Info] err:%d want:%d", err, logger.ErrorNone)
	}

	if err := log.Warning("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Debug() Warning] err:%d want:%d", err, logger.ErrorNone)
	}

	if err := log.Error("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Debug() Error] err:%d want:%d", err, logger.ErrorNone)
	}

	log.Deinit()
	os.Remove("log/debug.log")
}

func TestDebugJson(t *testing.T) {
	os.Remove("log/debug.log")

	var log logger.Logger
	log.Init(logger.LevelDebug, "log/debug.log", logger.OutputTypeJson)

	if err := log.Debug("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Debug() Type:Json Debug] err:%d want:%d", err, logger.ErrorNone)
	}

	if err := log.Info("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Debug() Type:Json Info] err:%d want:%d", err, logger.ErrorNone)
	}

	if err := log.Warning("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Debug() Type:Json Warning] err:%d want:%d", err, logger.ErrorNone)
	}

	if err := log.Error("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Debug() Type:Json Error] err:%d want:%d", err, logger.ErrorNone)
	}

	log.Deinit()
	os.Remove("log/debug.log")
}

func TestInfo(t *testing.T) {
	os.Remove("log/info.log")

	var log logger.Logger
	log.Init(logger.LevelInfo, "log/info.log", logger.OutputTypeNormal)

	if err := log.Debug("str"); err != logger.ErrorLevelNotEnough {
		t.Errorf("Test case[Info() Debug] err:%d want:%d", err, logger.ErrorLevelNotEnough)
	}

	if err := log.Info("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Info() Info] err:%d want:%d", err, logger.ErrorNone)
	}

	if err := log.Warning("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Info() Warning] err:%d want:%d", err, logger.ErrorNone)
	}

	if err := log.Error("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Info() Error] err:%d want:%d", err, logger.ErrorNone)
	}

	log.Deinit()
	os.Remove("log/info.log")
}

func TestWarning(t *testing.T) {
	os.Remove("log/warning.log")

	var log logger.Logger
	log.Init(logger.LevelWarning, "log/warning.log", logger.OutputTypeNormal)

	if err := log.Debug("str"); err != logger.ErrorLevelNotEnough {
		t.Errorf("Test case[Warning() Debug] err:%d want:%d", err, logger.ErrorLevelNotEnough)
	}

	if err := log.Info("str"); err != logger.ErrorLevelNotEnough {
		t.Errorf("Test case[Warning() Info] err:%d want:%d", err, logger.ErrorLevelNotEnough)
	}

	if err := log.Warning("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Warning() Warning] err:%d want:%d", err, logger.ErrorNone)
	}

	if err := log.Error("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Warning() Error] err:%d want:%d", err, logger.ErrorNone)
	}

	log.Deinit()
	os.Remove("log/warning.log")
}

func TestError(t *testing.T) {
	os.Remove("log/error.log")

	var log logger.Logger
	log.Init(logger.LevelError, "log/error.log", logger.OutputTypeNormal)

	if err := log.Debug("str"); err != logger.ErrorLevelNotEnough {
		t.Errorf("Test case[Error() Debug] err:%d want:%d", err, logger.ErrorLevelNotEnough)
	}

	if err := log.Info("str"); err != logger.ErrorLevelNotEnough {
		t.Errorf("Test case[Error() Info] err:%d want:%d", err, logger.ErrorLevelNotEnough)
	}

	if err := log.Warning("str"); err != logger.ErrorLevelNotEnough {
		t.Errorf("Test case[Error() Warning] err:%d want:%d", err, logger.ErrorLevelNotEnough)
	}

	if err := log.Error("str"); err != logger.ErrorNone {
		t.Errorf("Test case[Error() Error] err:%d want:%d", err, logger.ErrorNone)
	}

	log.Deinit()
	os.Remove("log/error.log")
}

func TestDeinit(t *testing.T) {
	var log logger.Logger
	log.Init(logger.LevelError, "log/test.log", logger.OutputTypeNormal)

	if err := log.Deinit(); err != logger.ErrorNone {
		t.Errorf("Test case[Deinit() normal] err:%d want:%d", err, logger.ErrorNone)
	}

	if err := log.Deinit(); err != logger.ErrorLogFileCloseFailed {
		t.Errorf("Test case[Deinit() error] err:%d want:%d", err, logger.ErrorLogFileCloseFailed)
	}

	os.Remove("log/test.log")
}
