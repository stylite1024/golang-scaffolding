package logger

import (
	"go-app/pkg/config"
	"go-app/pkg/tools"
	"os"

	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var lg *zap.Logger

// 初始化zap logger
func Setup(conf *config.LogConfig, mode string) (err error) {
	var core zapcore.Core

	encoder := getEncoder()

	// // 实现两个判断日志等级的interface
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	// 按大小切割
	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := getLogWriterBySize(conf.InfoFilename, conf.MaxSize, conf.MaxBackups, conf.MaxAge, conf.Compress)
	errorWriter := getLogWriterBySize(conf.ErrorFilename, conf.MaxSize, conf.MaxBackups, conf.MaxAge, conf.Compress)

	// 按时间切割
	// infoWriter := getLogWriterByTime(conf.InfoFilename, conf.MaxAge)
	// errorWriter := getLogWriterByTime(conf.ErrorFilename, conf.MaxAge)

	if mode == "dev" {
		// 进入开发模式，日志输出到终端
		// console不同级别日志颜色方案：https://github.com/uber-go/zap/pull/307
		config := zap.NewDevelopmentEncoderConfig()
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(encoder, infoWriter, infoLevel),
			zapcore.NewCore(encoder, errorWriter, errorLevel),
		)
	} else {
		// 创建具体的Logger
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, infoWriter, zapcore.InfoLevel),
			zapcore.NewCore(encoder, errorWriter, zap.ErrorLevel),
		)
	}
	lg = zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	zap.ReplaceGlobals(lg)
	zap.L().Info(tools.Green("logger init success !"))
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "ts"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 按大小切割
func getLogWriterBySize(filename string, maxSize, maxBackup, maxAge int, compress bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 按时间切割
// 根据时间切割，请看： https://www.jianshu.com/p/d729c7ec9c85
func getLogWriterByTime(filename string, maxAge int) zapcore.WriteSyncer {
	hook, err := rotatelogs.New(
		filename+".%Y%m%d",
		// strings.Replace(filename, ".log", "", -1)+"-%Y%m%d.log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(maxAge)),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(hook)
}
