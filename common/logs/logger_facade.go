package logs

import "github.com/starter-go/vlog"

// LoggerFacade 实现 Logger 接口
type LoggerFacade struct {
	vlog.LoggerAdapter
}

func (inst *LoggerFacade) init() vlog.Logger {
	return initLoggerAdapter(&inst.LoggerAdapter)
}

// // ForSender ...
// func (inst *LoggerFacade) ForSender(sender any) vlog.Logger {
// 	lc := inst.context.Clone()
// 	lc.sender = sender
// 	return &LoggerFacade{context: lc}
// }

// // ForTag ...
// func (inst *LoggerFacade) ForTag(tag string) vlog.Logger {
// 	lc := inst.context.Clone()
// 	lc.tag = tag
// 	return &LoggerFacade{context: lc}
// }

// // IsTraceEnabled 判断是否输出等级为（trace）的日志
// func (inst *LoggerFacade) IsTraceEnabled() bool {
// 	return inst.context.IsLevelEnabled(vlog.TRACE)
// }

// // IsDebugEnabled 判断是否输出等级为（debug）的日志
// func (inst *LoggerFacade) IsDebugEnabled() bool {
// 	return inst.context.IsLevelEnabled(vlog.DEBUG)
// }

// // IsInfoEnabled 判断是否输出等级为（info）的日志
// func (inst *LoggerFacade) IsInfoEnabled() bool {
// 	return inst.context.IsLevelEnabled(vlog.INFO)
// }

// // IsWarnEnabled 判断是否输出等级为（warn）的日志
// func (inst *LoggerFacade) IsWarnEnabled() bool {
// 	return inst.context.IsLevelEnabled(vlog.WARN)
// }

// // IsErrorEnabled 判断是否输出等级为（error）的日志
// func (inst *LoggerFacade) IsErrorEnabled() bool {
// 	return inst.context.IsLevelEnabled(vlog.ERROR)
// }

// // IsFatalEnabled 判断是否输出等级为（fatal）的日志
// func (inst *LoggerFacade) IsFatalEnabled() bool {
// 	return inst.context.IsLevelEnabled(vlog.FATAL)
// }

// // Trace 输出等级为(Trace)的日志
// func (inst *LoggerFacade) Trace(fmt string, args ...any) {
// 	inst.context.LogWithLevel(vlog.TRACE, fmt, args...)
// }

// // Debug 输出等级为(Debug)的日志
// func (inst *LoggerFacade) Debug(fmt string, args ...any) {
// 	inst.context.LogWithLevel(vlog.DEBUG, fmt, args...)
// }

// // Info 输出等级为(Info)的日志
// func (inst *LoggerFacade) Info(fmt string, args ...any) {
// 	inst.context.LogWithLevel(vlog.INFO, fmt, args...)
// }

// // Warn 输出等级为(Warn)的日志
// func (inst *LoggerFacade) Warn(fmt string, args ...any) {
// 	inst.context.LogWithLevel(vlog.WARN, fmt, args...)
// }

// // Error 输出等级为(Error)的日志
// func (inst *LoggerFacade) Error(fmt string, args ...any) {
// 	inst.context.LogWithLevel(vlog.ERROR, fmt, args...)
// }

// // Fatal 输出等级为(Fatal)的日志
// func (inst *LoggerFacade) Fatal(fmt string, args ...any) {
// 	inst.context.LogWithLevel(vlog.FATAL, fmt, args...)
// }
