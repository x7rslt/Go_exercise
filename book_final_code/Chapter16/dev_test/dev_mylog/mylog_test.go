package testmylog

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var Log *zap.SugaredLogger

const(
	output_dir ="./logs/"
	out_path = "app.MyLog"
	err_path = "app.myerr"

)

func init(){
	_,err := os.Stat(output_dir)
	if err != nil{
		if os.IsNotExist(err){
			err := os.Mkdir(output_dir,os.ModePerm)
			if err:=nil{
				fmt.Println("创建目录失败！[%v]\n",err)
			}
		}
	}

	//设置一些基本日志格式
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:"msg",
		LevelKey:"level",
		TimeKey:"ts",
		CallerKey:"caller",
		StacktrackKey:"trace",
		LineEnding:zapcore.DefaultLineEnding,
		EncodeLevel:zapcore.LowercaseLevelEncoder,
		EncodeCaller:zapcore.ShortCallerEncoder,
		EncodeTime:func(t time.Time,enc zapcore.PrimitiveArrayEncoder){
			enc.AppendInt64(int64(d)/1000000)

		},

	})
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level)bool{
		return true
	})
	warnLevel:= zap.LevelEnablerFunc(func(lvl zapcore.Level)bool{
		return lvl >= zapcore.WarnLevel
	})
	infoHook_1 := os.Stdout
	infoHook_2:=getWrite(out_path)
	errorHook:= getWrite(err_path)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder,zapcore.AddSync(infoHook_1),infoLevel),
		zapcore.NewCore(encoder,zapcore.AddSync(infoHook_2),infoLevel),
		zapcore.NewCore(encoder,zapcore.AddSync(errorHook),warnLevel),
	)
	logger := zap.New(core,zap.AddCaller(),zap.AddStacktrace(zap.ErrorLevel))
	Log = logger.Sugar()
	defer logger.Sync()
}

func getWriter(filename string)io.Writer{
	hook,err := rotatelogs.New(
		output_dir+filename+".%Y%m%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7)
		rotatelogs.WithRotatimeTime(time.Hour*24)
	)
	if err !=nil{
		panic(err)
	}
	return hook
}