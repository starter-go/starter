package logs

import (
	"fmt"

	"github.com/starter-go/vlog"
)

// ConsoleWriter 用于向默认控制台输出日志
type ConsoleWriter struct {
	//starter:component
	_as func(vlog.MessageFilterRegistry) //starter:as(".")
}

func (inst *ConsoleWriter) _impl() (vlog.MessageFilter, vlog.MessageFilterRegistry) {
	return inst, inst
}

// ListLogFilterRegistration ...
func (inst *ConsoleWriter) ListLogFilterRegistration() []*vlog.MessageFilterRegistration {
	r := &vlog.MessageFilterRegistration{
		Filter: inst,
		Name:   "console-log-writer",
	}
	return []*vlog.MessageFilterRegistration{r}
}

// DoFilter ...
func (inst *ConsoleWriter) DoFilter(msg *vlog.Message, chain vlog.MessageFilterChain) {
	fmt.Println(msg.Text)
	chain.DoFilter(msg)
}
