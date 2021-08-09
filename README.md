# logrus FileHookRotation

This is a simple hook for logrus to write log files using https://github.com/natefinch/lumberjack

```golang
hook, err := filerotationhook.NewFileRotationHook(filerotationhook.Config{
    Filename: "logfile.log",
    MaxSize: 5,
    MaxBackups: 7,
    MaxAge: 7,
    Level: logrus.LevelDebug,
    Formatter: logrus.TextFormatter,
})
```
