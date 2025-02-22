package log

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/BoyChai/CoralBot/config"
)

var (

	// 前台Logger
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger

	// 文件Logger
	debugFileLogger *log.Logger
	infoFileLogger  *log.Logger
	warnFileLogger  *log.Logger
	errorFileLogger *log.Logger
	fatalFileLogger *log.Logger

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

const (
	// 颜色重置
	colorReset = "\033[0m"
	// 红色
	colorRed = "\033[31m"
	// 黄色
	colorYellow = "\033[33m"
	// 青色
	colorCyan = "\033[36m"
	// 灰色
	colorGray = "\033[90m"
)

func init() {
	fileLock = sync.RWMutex{}
	if config.Cfg.Log {
		SetFile("logs/coralbot.log")
	}
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
		debugLogger.Printf(format, v...)
		if config.Cfg.Log {
			debugFileLogger.Printf(format, v...)
		}
	}
}
func Info(format string, v ...any) {
	checkIfDayChange()
	if logLevel <= InfoLevel {
		infoLogger.Printf(format, v...)
		if config.Cfg.Log {
			infoFileLogger.Printf(format, v...)
		}
	}
}
func Warn(format string, v ...any) {
	checkIfDayChange()
	if logLevel <= WarnLevel {
		warnLogger.Printf(format, v...)
		if config.Cfg.Log {
			warnFileLogger.Printf(format, v...)
		}
	}
}
func Error(format string, v ...any) {
	checkIfDayChange()
	if logLevel <= ErrorLevel {
		errorLogger.Printf(getPrefix()+format, v...)
		if config.Cfg.Log {
			errorFileLogger.Printf(getPrefix()+format, v...)
		}
	}
}
func Fatal(format string, v ...any) {
	checkIfDayChange()
	if logLevel <= FatalLevel {
		fatalLogger.Printf(getPrefix()+format, v...)
		if config.Cfg.Log {
			fatalFileLogger.Printf(getPrefix()+format, v...)
		}
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
	// 前台Logger
	infoLogger = log.New(os.Stdout, colorCyan+"[INFO] "+colorReset, log.LstdFlags)
	debugLogger = log.New(os.Stdout, colorGray+"[DEBUG]"+colorReset, log.LstdFlags)
	warnLogger = log.New(os.Stdout, colorYellow+"[WARN] "+colorReset, log.LstdFlags)
	errorLogger = log.New(os.Stdout, colorRed+"[ERROR]"+colorReset, log.LstdFlags)
	fatalLogger = log.New(os.Stdout, colorRed+"[FATAL]"+colorReset, log.LstdFlags)

	// 文件Logger
	if config.Cfg.Log {
		infoFileLogger = log.New(logOut, "[INFO] ", log.LstdFlags)
		debugFileLogger = log.New(logOut, "[DEBUG]", log.LstdFlags)
		warnFileLogger = log.New(logOut, "[WARN] ", log.LstdFlags)
		errorFileLogger = log.New(logOut, "[ERROR]", log.LstdFlags)
		fatalFileLogger = log.New(logOut, "[FATAL]", log.LstdFlags)
	}

}
