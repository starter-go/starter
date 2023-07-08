package logs

import (
	"time"

	"github.com/starter-go/application/attributes"
	"github.com/starter-go/vlog"
)

// StartupBuffer 是 starter 启动初期的日志缓冲区
type StartupBuffer interface {
	Init(atts attributes.Table)
}

////////////////////////////////////////////////////////////////////////////////

const attrNameForStartupBuffer = "github.com/starter-go/starter/common/logs/StartupBuffer#binding"

type myStartupBuffer struct {
	logs []*vlog.Message
}

// Init ...
func (inst *myStartupBuffer) Init(atts attributes.Table) {
	const name = attrNameForStartupBuffer
	atts.SetAttribute(name, inst)
	vlog.SetLoggerFactory(inst.getLoggerFactory())
}

func (inst *myStartupBuffer) push(msg *vlog.Message) {
	if msg == nil {
		return
	}
	inst.logs = append(inst.logs, msg)
}

func (inst *myStartupBuffer) getLoggerFactory() vlog.LoggerFactory {
	return &myStartupBufferLoggerFactory{buffer: inst}
}

////////////////////////////////////////////////////////////////////////////////

// NewStartupBuffer 创建 starter 启动初期的日志缓冲区
func NewStartupBuffer() StartupBuffer {
	return &myStartupBuffer{}
}

////////////////////////////////////////////////////////////////////////////////

type myStartupBufferLoggerFactory struct {
	buffer *myStartupBuffer
}

func (inst *myStartupBufferLoggerFactory) _impl() vlog.LoggerFactory {
	return inst
}

func (inst *myStartupBufferLoggerFactory) Create() vlog.Logger {
	return &myStartupBufferLogger{buffer: inst.buffer}
}

////////////////////////////////////////////////////////////////////////////////

type myStartupBufferLogger struct {
	buffer *myStartupBuffer
}

func (inst *myStartupBufferLogger) _impl() vlog.Logger {
	return inst
}

func (inst *myStartupBufferLogger) log(level vlog.Level, fmt string, args ...any) {
	msg := &vlog.Message{}
	msg.Timestamp = time.Now()
	msg.Level = level
	msg.Format = fmt
	msg.Arguments = args
	msg.Sender = inst.buffer
	msg.Tag = "starter-startup"
	inst.buffer.push(msg)
}

func (inst *myStartupBufferLogger) Trace(fmt string, args ...any) {
	inst.log(vlog.TRACE, fmt, args...)
}

func (inst *myStartupBufferLogger) Debug(fmt string, args ...any) {
	inst.log(vlog.DEBUG, fmt, args...)
}

func (inst *myStartupBufferLogger) Info(fmt string, args ...any) {
	inst.log(vlog.INFO, fmt, args...)
}

func (inst *myStartupBufferLogger) Warn(fmt string, args ...any) {
	inst.log(vlog.WARN, fmt, args...)
}

func (inst *myStartupBufferLogger) Error(fmt string, args ...any) {
	inst.log(vlog.ERROR, fmt, args...)
}

func (inst *myStartupBufferLogger) Fatal(fmt string, args ...any) {
	inst.log(vlog.FATAL, fmt, args...)
}

func (inst *myStartupBufferLogger) IsTraceEnabled() bool {
	return true
}

func (inst *myStartupBufferLogger) IsDebugEnabled() bool {
	return true
}

func (inst *myStartupBufferLogger) IsInfoEnabled() bool {
	return true
}

func (inst *myStartupBufferLogger) IsWarnEnabled() bool {
	return true
}

func (inst *myStartupBufferLogger) IsErrorEnabled() bool {
	return true
}

func (inst *myStartupBufferLogger) IsFatalEnabled() bool {
	return true
}

func (inst *myStartupBufferLogger) ForSender(sender any) vlog.Logger {
	return inst // 不支持 sender
}

func (inst *myStartupBufferLogger) ForTag(tag string) vlog.Logger {
	return inst // 不支持 tag
}

////////////////////////////////////////////////////////////////////////////////
