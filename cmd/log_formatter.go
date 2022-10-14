package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)


// MyFromatter custom, human readable log format
type MyFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var fields string
	keys := make([]string, 0, len(entry.Data))
	for k, v := range entry.Data {
		if k == "tag" {
			continue
		}

		if k == "source" {
			continue
		}
		keys = append(keys, fmt.Sprintf("%s: %v", k, v))
	}

	if len(keys) > 0 {
		fields = fmt.Sprintf("[%s] ", strings.Join(keys, ", "))
	}

	if _, ok := entry.Data["tag"]; !ok {
		entry.Data["tag"] = ""
	}

	return []byte(fmt.Sprintf("%s %s [%s] %s%s: %s\n", entry.Time.Format(f.TimestampFormat), f.LevelDesc[entry.Level], entry.Data["tag"], fields, entry.Data["source"], entry.Message)), nil
}

type ContextHook struct{}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook ContextHook) Fire(entry *logrus.Entry) error {
	_, pkg := path.Split(path.Dir(entry.Caller.File))
	file := path.Base(entry.Caller.File)
	entry.Data["source"] = fmt.Sprintf("%s/%s:%v", pkg, file, entry.Caller.Line)
	return nil
}
