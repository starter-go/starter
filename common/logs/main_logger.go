package logs

import (
	"time"

	"github.com/starter-go/application"
	"github.com/starter-go/starter/common"
	"github.com/starter-go/vlog"
)

// MainLogger 用于向文件输出日志
type MainLogger struct {
	//starter:component(alias="starter-main-logger-holder")

	Filters       []vlog.MessageFilterRegistry //starter:inject(".")
	Level         string                       //starter:inject("${vlog.level}")
	MainGroupName string                       //starter:inject("${vlog.main}")
	Context       application.Context          //starter:inject("context")

	logger      vlog.Logger
	lc          *LoggerContext //  level vlog.Level
	startupTime time.Time
}

func (inst *MainLogger) _impl() (application.Lifecycle, vlog.LoggerFactory, vlog.LoggerHolder) {
	return inst, inst, inst
}

// Life ...
func (inst *MainLogger) Life() *application.Life {
	return &application.Life{
		Order:       common.StartupOrderLogger,
		OnCreate:    inst.onInit,
		OnStartPost: inst.onStarted,
	}
}

func (inst *MainLogger) onInit() error {
	err := inst.initLoggerContext()
	if err != nil {
		return err
	}
	vlog.SetLoggerFactory(inst)
	inst.flushStartupLogs()
	logger := vlog.GetLogger()
	inst.logger = logger
	return nil
}

func (inst *MainLogger) onStarted() error {

	// like 'Started MyApplication in 0.906 seconds (process running for 6.514)'
	const f = "Started app(%s) in %s"

	t0 := inst.startupTime
	t1 := time.Now()
	diff := t1.Sub(t0)
	appName := inst.Context.GetMainModule().Name()

	inst.lc.LogWithLevel(vlog.INFO, f, appName, diff.String())
	return nil
}

func (inst *MainLogger) flushStartupLogs() {

	const name = attrNameForStartupBuffer
	atts := inst.Context.GetAttributes()
	object := atts.GetAttribute(name)
	atts.SetAttribute(name, nil)

	buffer := object.(*myStartupBuffer)
	src := buffer.logs
	buffer.logs = nil
	inst.startupTime = buffer.startupTime
	for _, msg := range src {
		inst.lc.LogWithMessage(msg)
	}
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

// Logger ...
func (inst *MainLogger) Logger() vlog.Logger {
	l := inst.logger
	if l == nil {
		l = vlog.GetLogger()
	}
	return l
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
