// core/init_logrus.go
// 初始化日志格式及日志分割方式
package core

import (
	"blogx_backend/global"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

/**********************************************************************************************************/
/**********************************为logrus添加HOOK进行日志拆分***********************************************/
/**********************************************************************************************************/

// logrus在记录Levels()返回的日志级别的消息时会触发HOOK，
// 按照Fire方法定义的内容修改logrus.Entry。
type FileDateHook struct {
	file     *os.File
	logPath  string
	fileDate string // 判断日期切换目录
	appName  string
}

// 定义哪些level生效
func (hook FileDateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook FileDateHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format(time.DateOnly) // "2006-01-02", 返回值是对应格式的当前时间
	line, _ := entry.String()                 // 读取当前日志缓冲区中的日志

	// 如果文件时间是对应的时间，则追加写入
	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		return nil
	}

	// 时间不相等，说明日期不对，该切换目录了
	hook.file.Close()
	os.MkdirAll(fmt.Sprintf("%s/%s", hook.logPath, timer), os.ModePerm) // os.MkdirAll就是路径上所有不存在的目录都一并创建

	fileName := fmt.Sprintf("%s/%s/%s.log", hook.logPath, timer, hook.appName)

	// open的文件描述符权限是只写、并且是追加写而不是覆盖写、及文件不存在时创建
	// 如果文件不存在要创建，那么创建权限是0600对应d rwx rwx rwx 中的 - r-- --- ---
	hook.file, _ = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	hook.file.Write([]byte(line))
	return nil
}

func initFile(logPath, appName string) {
	fileDate := time.Now().Format(time.DateOnly)

	// 创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件
	filename := fmt.Sprintf("%s/%s/%s.log", logPath, fileDate, appName)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		logrus.Error(err)
		return
	}
	fileHook := FileDateHook{file, logPath, fileDate, appName}
	logrus.AddHook(&fileHook)
}

/***********************************end of hook*******************************************************/

/**********************************************************************************************************/
/*******************************初始化logrus格式及正式添加hook进行文件拆分***********************************/
// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

// logrus需要自己接口 (logFormatter struct)::Formatter(*logrus.Entry) []byte, error

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同的日志级别（level）去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	// 初始化一个写日志的缓冲区
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 自定义日志格式
	timestamp := entry.Time.Format(time.DateTime)
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)

		// 自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n",
			timestamp, levelColor, entry.Level.String(), fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level.String(), entry.Message)
	}

	return b.Bytes(), nil
}

// 下面是新建一个logger对象，这样的话后面打印就需要使用这个logrus函数进行操作
//var log *logrus.Logger
//
//func init() {
//	log = NewLog()
//}
//
//func NewLog() *logrus.Logger {
//	mLog := logrus.New()               // 新建一个实例
//	mLog.SetOutput(os.Stdout)          // 设置输出类型
//	mLog.SetReportCaller(true)         // 开启返回函数名和行号
//	mLog.SetFormatter(&LogFormatter{}) // 设置自定义的Formatter
//	mLog.SetLevel(logrus.DebugLevel)   // 设置最低的Level
//	return mLog
//}

// 也可以直接操纵原始的logrus
// 该函数是初始化logrus的入口，在主程序中执行
func InitLogrus() {
	logrus.SetOutput(os.Stdout)          // 设置输出类型
	logrus.SetReportCaller(true)         // 开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{}) // 设置自定义的Formatter
	logrus.SetLevel(logrus.DebugLevel)   // 设置最低的Level
	initFile(global.Config.Log.Dir, global.Config.Log.App)
}

/*****************************************end of initLogrus***************************************************/
