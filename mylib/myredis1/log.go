package myredis

import (
	"context"

	"github.com/phial3/mylib/mylog"
)

type logger struct {
}

func (l *logger) Printf(ctx context.Context, format string, v ...interface{}) {
	mylog.Infof(format, v...)
}
