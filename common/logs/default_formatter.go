package logs

import (
	"fmt"

	"github.com/starter-go/vlog"
)

// DefaultFormatter 用于（以默认格式）格式化日志消息
type DefaultFormatter struct {
	//starter:component
	_as func(vlog.MessageFilterRegistry) //starter:as(".")

	Format string //starter:inject("${vlog.formatters.default.format}")

	formatter Formatter
}

func (inst *DefaultFormatter) _impl() (vlog.MessageFilter, vlog.MessageFilterRegistry) {
	return inst, inst
}

// ListLogFilterRegistration ...
func (inst *DefaultFormatter) ListLogFilterRegistration() []*vlog.MessageFilterRegistration {

	inst.formatter.format = inst.Format
	inst.formatter.templ = nil

	r := &vlog.MessageFilterRegistration{
		Filter: inst,
		Name:   "default-log-formatter",
	}
	return []*vlog.MessageFilterRegistration{r}
}

// DoFilter ...
func (inst *DefaultFormatter) DoFilter(msg *vlog.Message, chain vlog.MessageFilterChain) {
	msg.Text = inst.formatMessage(msg)
	chain.DoFilter(msg)
}

func (inst *DefaultFormatter) formatMessage(msg *vlog.Message) string {
	text1 := inst.formatMessageHead(msg)
	text2 := fmt.Sprintf(msg.Format, msg.Arguments...)
	return text1 + " " + text2
}

func (inst *DefaultFormatter) formatMessageHead(msg *vlog.Message) string {
	data := &FormatData{}
	data.Init(msg)
	return inst.formatter.Format(data)
}
