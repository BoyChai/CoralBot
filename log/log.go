package log

import (
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger

	logOut     *os.File
	logLevel   int
	currentDay int
	logFile    string
	fileLock   sync.RWMutex
)

const (
	DebugLevel = iota // 0
	InfoLevel         // 1
	WarnLevel         // 2
	ErrorLevel        // 3
	FatalLevel        // 4
)

func init() {
	fileLock = sync.RWMutex{}
}

func SetLevel(level int) {
	logLevel = level
}

func SetFile(file string) {
	var err error
	logOut, err = os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	} else {
		// 获取今天是当年的第几天
		currentDay = time.Now().YearDay()
		initLog(logOut)
		logFile = file
	}
}

func Debug(format string, v ...any) {
	checkIfDayChange()
	if logLevel <= DebugLevel {
		debugLogger.Printf(getPrefix()+format, v...)
	}
}
func Info(format string, v ...any) {
	checkIfDayChange()
	if logLevel <= InfoLevel {
		infoLogger.Printf(format, v...)
		// infoLogger.Printf(getPrefix()+format, v...)
	}
}
func Warn(format string, v ...any) {
	checkIfDayChange()
	if logLevel <= WarnLevel {
		warnLogger.Printf(format, v...)
	}
}
func Error(format string, v ...any) {
	checkIfDayChange()
	if logLevel <= ErrorLevel {
		errorLogger.Printf(format, v...)
	}
}
func Fatal(format string, v ...any) {
	checkIfDayChange()
	if logLevel <= FatalLevel {
		fatalLogger.Printf(getPrefix()+format, v...)
	}
}

func getCallTrace() (string, int) {
	// 函数名 文件 行号 是否出现异常
	//pc, file, line, ok := runtime.Caller(0)

	_, file, line, ok := runtime.Caller(3)
	if ok {
		return file, line
	}
	return "", 0
}

func getPrefix() string {
	file, line := getCallTrace()

	return file + ":" + strconv.Itoa(line) + " "
}

func checkIfDayChange() {
	// 锁
	fileLock.Lock()
	defer fileLock.Unlock()
	day := time.Now().YearDay()
	if day == currentDay {
		return
	} else {
		currentDay = day
		logOut.Close()
		postFix := time.Now().Add(-24 * time.Hour).Format("20060102")
		os.Rename(logFile, logFile+"."+postFix)
		logOut, _ = os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
		initLog(logOut)
	}
}

func initLog(logOut *os.File) {
	infoLogger = log.New(io.MultiWriter(logOut, os.Stdout), "[INFO] ", log.LstdFlags)
	debugLogger = log.New(io.MultiWriter(logOut, os.Stdout), "[DEBUG]", log.LstdFlags)
	warnLogger = log.New(io.MultiWriter(logOut, os.Stdout), "[WARN]", log.LstdFlags)
	errorLogger = log.New(io.MultiWriter(logOut, os.Stdout), "[ERROR]", log.LstdFlags)
	fatalLogger = log.New(io.MultiWriter(logOut, os.Stdout), "[FATAL]", log.LstdFlags)
}
