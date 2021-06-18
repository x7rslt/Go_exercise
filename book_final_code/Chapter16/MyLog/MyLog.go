package MyLog

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

const (
	output_dir = "./logs/"
	out_path   = "app.MyLog"
	err_path   = "app.myerr"
)

func init() {
	_, err := os.Stat(output_dir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(output_dir, os.ModePerm)
			if err != nil {
				fmt.Printf("创建目录失败！[%v]\n", err)
			}
		}
	}

	//设置一些基本日志格式
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "ts",
		CallerKey:     "caller",
		StacktraceKey: "trace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))

		},
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})
	infoHook_1 := os.Stdout
	infoHook_2 := getWriter(out_path)
	errorHook := getWriter(err_path)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoHook_1), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoHook_2), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorHook), warnLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	Log = logger.Sugar()
	defer logger.Sync()
}

func getWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		output_dir+filename+".%Y%m%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
func Divide(a, b int) (int, error) {
	if b != 0 {
		c := a / b
		return c, nil
	} else {
		return 0, errors.New("Func run error:b is 0")
	}
}
