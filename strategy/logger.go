package strategy

import "fmt"

// 一个日志管理器
type LogManager struct {
	Logger
}

func NewLogManager(logger Logger) *LogManager {
	return &LogManager{
		logger,
	}
}

// 抽象的日志
type Logger interface {
	Info()
	Erorr()
}

// 文件保存日志
type FileLogger struct {
}

func (fl *FileLogger) Info() {
	fmt.Println("file logger info...")
}

func (fl *FileLogger) Erorr() {
	fmt.Println("file logger error...")
}

// 数据库保存日志
type DbLogger struct {
}

func (dl *DbLogger) Info() {
	fmt.Println("db logger info...")
}

func (dl *DbLogger) Erorr() {
	fmt.Println("db logger error...")
}
