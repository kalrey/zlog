package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Logger *zap.Logger

func init() {
	Logger = newZapLogger("./logs")
}



func newZapLogger(path string) *zap.Logger {

	hook := lumberjack.Logger{
		Filename:   path + "/data.log", // 日志文件路径
		MaxSize:    2048,               // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 300,                // 日志文件最多保存多少个备份
		MaxAge:     7,                  // 文件最多保存多少天
		Compress:   false,              // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	//// 设置初始化字段
	//filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	logger := zap.New(core, caller, development)

	//alevel := zap.NewAtomicLevel()
	//http.HandleFunc("/handle/level", alevel.ServeHTTP)
	//go func() {
	//	if err := http.ListenAndServe(":29090", nil); err != nil {
	//		panic(err)
	//	}
	//}()

	logger.Info("DefaultLogger init success")
	return logger
}
