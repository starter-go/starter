package logs

import (
	"github.com/starter-go/application"
	"github.com/starter-go/vlog"
)

// MainLogger 用于向文件输出日志
type MainLogger struct {
	//starter:component

	Filters       []vlog.MessageFilterRegistry //starter:inject(".")
	Level         string                       //starter:inject("${vlog.level}")
	MainGroupName string                       //starter:inject("${vlog.main}")

	lc *LoggerContext //  level vlog.Level
}

func (inst *MainLogger) _impl() (application.Lifecycle, vlog.LoggerFactory) {
	return inst, inst
}

// Life ...
func (inst *MainLogger) Life() *application.Life {
	return &application.Life{
		Order:    -10000,
		OnCreate: inst.init,
	}
}

func (inst *MainLogger) init() error {
	err := inst.initLoggerContext()
	if err != nil {
		return err
	}
	vlog.SetLoggerFactory(inst)
	return nil
}

func (inst *MainLogger) initLoggerContext() error {

	level, err := vlog.ParseLevel(inst.Level)
	if err != nil {
		return err
	}

	cm := &ChainMaker{}
	chain, err := cm.Make(inst.Filters, inst.MainGroupName)
	if err != nil {
		return err
	}

	lc := &LoggerContext{}
	lc.chain = chain
	lc.acceptLevels = level
	lc.sender = inst
	lc.tag = "MainLogger"

	inst.lc = lc
	return nil
}

// Create ...
func (inst *MainLogger) Create() vlog.Logger {
	lc := inst.lc
	if lc == nil {
		panic("LoggerContext is nil")
	}
	facade := &LoggerFacade{}
	facade.context = lc
	return facade
}
