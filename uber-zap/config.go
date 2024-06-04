package uberzap

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func ZapConfig() {
	hook := lumberjack.Logger{
		Filename:   "./logs/spikeProxy1.log", // 日志文件路径
		MaxSize:    128,                      // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                       // 日志文件最多保存多少个备份
		MaxAge:     7,                        // 文件最多保存多少天
		Compress:   true,                     // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     getEncodeTime,                  // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	// switch logLevel {
	// case "debug":
	// 	atomicLevel.SetLevel(zap.DebugLevel)
	// case "info":
	// 	atomicLevel.SetLevel(zap.InfoLevel)
	// case "warn":
	// 	atomicLevel.SetLevel(zap.WarnLevel)
	// case "error":
	// 	atomicLevel.SetLevel(zap.ErrorLevel)
	// case "dpanic":
	// 	atomicLevel.SetLevel(zap.DPanicLevel)
	// case "panic":
	// 	atomicLevel.SetLevel(zap.PanicLevel)
	// case "fatal":
	// 	atomicLevel.SetLevel(zap.FatalLevel)
	// default:
	// 	atomicLevel.SetLevel(zap.InfoLevel) // 默认级别
	// }
	//

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&hook),
		), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 构造日志
	logger := zap.New(core, caller, development)
	Logger = logger
}

func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02-15:04:05.000"))
}
