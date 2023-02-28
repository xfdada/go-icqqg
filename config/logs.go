package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

func parseLeve(s string) logrus.Level {
	types := strings.ToUpper(s)
	switch types {
	case "FATAL":
		return logrus.FatalLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "WARN":
		return logrus.WarnLevel
	case "INFO":
		return logrus.InfoLevel
	case "DEBUG":
		return logrus.DebugLevel
	case "TRACE":
		return logrus.TraceLevel
	}
	return logrus.DebugLevel
}

// ErrorLog 记录错误日志
func ErrorLog(msg string) *logrus.Logger {
	src, err := fileObj(AppConfig.Logger.ErrorPath)
	if err != nil {
		fmt.Println("err", err)
	}
	logger := logrus.New()
	//设置输出
	logger.Out = src
	//设置日志级别
	logger.SetLevel(parseLeve(AppConfig.Logger.ErrorLevel))
	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	errDetail := getCallerInfo(2)
	logger.WithFields(logrus.Fields{
		"err_detail": errDetail,
		"msg":        msg,
	}).Info()
	return logger
}

// AppLog 记录请求数据
func AppLog(statusCode int, latencyTime time.Duration, clientIP, reqMethod, reqUri string) *logrus.Logger {
	src, err := fileObj(AppConfig.Logger.AppPath)
	if err != nil {
		fmt.Println("err", err)
	}
	//实例化
	logger := logrus.New()
	//设置输出
	logger.Out = src
	//设置日志级别
	logger.SetLevel(parseLeve(AppConfig.Logger.AppLevel))
	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.WithFields(logrus.Fields{
		"status_code":  statusCode,
		"latency_time": latencyTime,
		"client_ip":    clientIP,
		"req_method":   reqMethod,
		"req_uri":      reqUri,
	}).Info()
	return logger
}

func SqlLog(s string) *logrus.Logger {
	src, err := fileObj(AppConfig.Logger.SqlPath)
	if err != nil {
		fmt.Println("err", err)
	}
	//实例化
	logger := logrus.New()
	//设置输出
	logger.Out = src
	//设置日志级别
	logger.SetLevel(parseLeve(AppConfig.Logger.SqlLevel))
	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.WithFields(logrus.Fields{}).Info(s)
	return logger
}

//

func fileObj(configPath string) (*os.File, error) {
	now := time.Now()
	logFilePath := configPath
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	return src, err
}

//getCallerInfo 获取运行时信息
func getCallerInfo(skip int) (info string) {

	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		info = "runtime.Caller() failed"
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // Base函数返回路径的最后一个元素
	return fmt.Sprintf("FuncName:%s, file:%s, line:%d ", funcName, fileName, lineNo)
}
