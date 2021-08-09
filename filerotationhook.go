// Original repository written by snowzach:
// https://github.com/snowzach/rotatefilehook

// MIT Licence:
// https://github.com/grumpypixel/filerotationhook/blob/master/LICENSE

package filerotationhook

import (
	"io"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	instance *FileRotationHook
	once     sync.Once

	lock = &sync.Mutex{}
)

type Config struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      logrus.Level
	Formatter  logrus.Formatter
}

type FileRotationHook struct {
	config *Config
	logger *lumberjack.Logger
	writer io.Writer
}

func NewFileRotationHook(cfg *Config) logrus.Hook {
	lock.Lock()
	defer lock.Unlock()

	once.Do(func() {
		hook := &FileRotationHook{
			config: cfg,
		}
		logger := &lumberjack.Logger{
			Filename:   cfg.Filename,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
		}
		hook.logger = logger
		hook.writer = logger
		instance = hook
	})
	return instance
}

func Instance() *FileRotationHook {
	return instance
}

func (hook *FileRotationHook) SetLevel(level logrus.Level) {
	hook.config.Level = level
}

func (hook *FileRotationHook) Rotate() {
	hook.logger.Rotate()
}

func (hook *FileRotationHook) Levels() []logrus.Level {
	return logrus.AllLevels[:hook.config.Level+1]
}

func (hook *FileRotationHook) Fire(entry *logrus.Entry) (err error) {
	b, err := hook.config.Formatter.Format(entry)
	if err != nil {
		return err
	}
	hook.writer.Write(b)
	return nil
}
