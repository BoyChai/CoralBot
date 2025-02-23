package log

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/BoyChai/CoralBot/config"
)

// Logger 定义日志级别常量和颜色常量
type Logger struct {
	level         int
	file          *os.File
	filePath      string
	currentDay    int
	mu            sync.RWMutex
	buffered      *bufio.Writer
	loggers       map[int]*log.Logger // 前台日志
	fileLoggers   map[int]*log.Logger // 文件日志
	colorPrefixes map[int]string      // 颜色前缀
	plainPrefixes map[int]string      // 纯文本前缀
}

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var (
	logger     *Logger
	once       sync.Once
	colorReset = "\033[0m"
	levels     = map[int]string{
		DebugLevel: "DEBUG",
		InfoLevel:  "INFO",
		WarnLevel:  "WARN",
		ErrorLevel: "ERROR",
		FatalLevel: "FATAL",
	}
	colors = map[int]string{
		DebugLevel: "\033[90m", // 灰色
		InfoLevel:  "\033[36m", // 青色
		WarnLevel:  "\033[33m", // 黄色
		ErrorLevel: "\033[31m", // 红色
		FatalLevel: "\033[31m", // 红色
	}
)

// GetLogger 获取单例Logger
func GetLogger() *Logger {
	once.Do(func() {
		logger = newLogger()
	})
	return logger
}

func newLogger() *Logger {
	l := &Logger{
		level:         InfoLevel,
		loggers:       make(map[int]*log.Logger),
		fileLoggers:   make(map[int]*log.Logger),
		colorPrefixes: make(map[int]string),
		plainPrefixes: make(map[int]string),
	}

	// 初始化前缀
	for lvl, name := range levels {
		l.colorPrefixes[lvl] = fmt.Sprintf("%s[%s]%s ", colors[lvl], name, colorReset)
		l.plainPrefixes[lvl] = fmt.Sprintf("[%s] ", name)
		l.loggers[lvl] = log.New(os.Stdout, l.colorPrefixes[lvl], log.LstdFlags)
	}

	return l
}

// SetLevel 设置日志级别
func (l *Logger) SetLevel(level int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if level >= DebugLevel && level <= FatalLevel {
		l.level = level
	}
}

// SetFile 设置日志文件
func (l *Logger) SetFile(filePath string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if err := l.closeFile(); err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		return err
	}

	l.file = file
	l.filePath = filePath
	l.currentDay = time.Now().YearDay()
	l.buffered = bufio.NewWriter(file)

	// 初始化文件logger
	for lvl := range levels {
		l.fileLoggers[lvl] = log.New(l.buffered, l.plainPrefixes[lvl], log.LstdFlags)
	}

	return nil
}

// 日志输出方法
func (l *Logger) log(level int, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	l.checkDayChange()

	// 前台输出
	l.loggers[level].Printf(format, v...)

	// 文件输出
	if config.Cfg.Log && l.file != nil {
		prefix := l.getCallerPrefix()
		l.fileLoggers[level].Printf(prefix+format, v...)
		l.buffered.Flush()
	}
}

func (l *Logger) Debug(format string, v ...interface{}) { l.log(DebugLevel, format, v...) }
func (l *Logger) Info(format string, v ...interface{})  { l.log(InfoLevel, format, v...) }
func (l *Logger) Warn(format string, v ...interface{})  { l.log(WarnLevel, format, v...) }
func (l *Logger) Error(format string, v ...interface{}) { l.log(ErrorLevel, format, v...) }
func (l *Logger) Fatal(format string, v ...interface{}) { l.log(FatalLevel, format, v...) }

// 获取调用者信息
func (l *Logger) getCallerPrefix() string {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d ", file, line)
}

// 检查日期变化并处理日志文件轮转
func (l *Logger) checkDayChange() {
	l.mu.RLock()
	day := time.Now().YearDay()
	if l.file == nil || day == l.currentDay {
		l.mu.RUnlock()
		return
	}
	l.mu.RUnlock()

	l.mu.Lock()
	defer l.mu.Unlock()

	// 再次检查防止并发修改
	if day == l.currentDay {
		return
	}

	l.currentDay = day
	l.closeFile()

	// 重命名旧文件
	postFix := time.Now().Add(-24 * time.Hour).Format("20060102")
	os.Rename(l.filePath, l.filePath+"."+postFix)

	// 创建新文件
	l.SetFile(l.filePath)
}

// 关闭文件
func (l *Logger) closeFile() error {
	if l.file != nil {
		l.buffered.Flush()
		err := l.file.Close()
		l.file = nil
		return err
	}
	return nil
}

// Init 初始化日志
func Init() {
	l := GetLogger()
	if config.Cfg.Log {
		if err := l.SetFile("logs/coralbot.log"); err != nil {
			l.Error("Failed to initialize log file: %v", err)
		}
	}
}
