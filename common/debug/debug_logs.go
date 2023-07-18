package debug

import (
	"sort"

	"github.com/starter-go/application"
	"github.com/starter-go/starter/common"
	"github.com/starter-go/vlog"
)

// ContextInfoLog ...
type ContextInfoLog struct {

	//starter:component

	Context application.Context //starter:inject("context")
	Logger  vlog.Logger         //starter:inject("#")

	EnableDebug bool //starter:inject("${debug.enabled}")
	LogArgs     bool //starter:inject("${debug.log-arguments}")
	LogEnv      bool //starter:inject("${debug.log-environment}")
	LogProps    bool //starter:inject("${debug.log-properties}")
}

func (inst *ContextInfoLog) _impl() {
}

// Life ...
func (inst *ContextInfoLog) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.init,
		Order:    common.StartupOrderDebug,
	}
}

func (inst *ContextInfoLog) init() error {
	inst.log()
	return nil
}

func (inst *ContextInfoLog) log() {

	if !inst.EnableDebug {
		return
	}

	if inst.LogArgs {
		args := inst.Context.GetArguments().Raw()
		inst.logList("arguments:", args)
	}

	if inst.LogEnv {
		env := inst.Context.GetEnvironment().Export(nil)
		inst.logTable("environment:", env)
	}

	if inst.LogProps {
		props := inst.Context.GetProperties().Export(nil)
		inst.logTable("properties:", props)
	}
}

func (inst *ContextInfoLog) logTable(tag string, table map[string]string) {
	src := table
	dst := make([]string, 0)
	for k, v := range src {
		item := k + " = " + v
		dst = append(dst, item)
	}
	sort.Strings(dst)
	inst.logList(tag, dst)
}

func (inst *ContextInfoLog) logList(tag string, list []string) {
	lg := inst.Logger
	lg.Info(tag)
	for _, item := range list {
		lg.Info("    %s", item)
	}
}
