package logs

import "github.com/starter-go/vlog"

// LoggerProxy 向 application.Context 提供一个代理的 vlog.Logger 接口，
// 并将日志转发给默认的日志组件
type LoggerProxy struct {

	//starter:component
	_as func(vlog.Logger) //starter:as("#")

	Holder vlog.LoggerHolder //starter:inject("#starter-main-logger-holder")

	// logger vlog.Logger
}

func (inst *LoggerProxy) _impl() vlog.Logger {
	return inst
}

// Logger ...
func (inst *LoggerProxy) Logger() vlog.Logger {

	// l := inst.logger
	// if l == nil {
	// 	l = vlog.GetLogger()
	// 	inst.logger = l
	// }

	return inst.Holder.Logger()
}

// Trace ...
func (inst *LoggerProxy) Trace(fmt string, args ...any) {
	inst.Logger().Trace(fmt, args...)
}

// Debug ...
func (inst *LoggerProxy) Debug(fmt string, args ...any) {
	inst.Logger().Debug(fmt, args...)
}

// Info ...
func (inst *LoggerProxy) Info(fmt string, args ...any) {
	inst.Logger().Info(fmt, args...)
}

// Warn ...
func (inst *LoggerProxy) Warn(fmt string, args ...any) {
	inst.Logger().Warn(fmt, args...)
}

// Error ...
func (inst *LoggerProxy) Error(fmt string, args ...any) {
	inst.Logger().Error(fmt, args...)
}

// Fatal ...
func (inst *LoggerProxy) Fatal(fmt string, args ...any) {
	inst.Logger().Fatal(fmt, args...)
}

// IsTraceEnabled ...
func (inst *LoggerProxy) IsTraceEnabled() bool {
	return inst.Logger().IsTraceEnabled()
}

// IsDebugEnabled ...
func (inst *LoggerProxy) IsDebugEnabled() bool {
	return inst.Logger().IsDebugEnabled()
}

// IsInfoEnabled ...
func (inst *LoggerProxy) IsInfoEnabled() bool {
	return inst.Logger().IsInfoEnabled()
}

// IsWarnEnabled ...
func (inst *LoggerProxy) IsWarnEnabled() bool {
	return inst.Logger().IsWarnEnabled()
}

// IsErrorEnabled ...
func (inst *LoggerProxy) IsErrorEnabled() bool {
	return inst.Logger().IsErrorEnabled()
}

// IsFatalEnabled ..
func (inst *LoggerProxy) IsFatalEnabled() bool {
	return inst.Logger().IsFatalEnabled()
}

// ForSender ...
func (inst *LoggerProxy) ForSender(sender any) vlog.Logger {
	return inst.Logger().ForSender(sender)
}

// ForTag ...
func (inst *LoggerProxy) ForTag(tag string) vlog.Logger {
	return inst.Logger().ForTag(tag)
}
