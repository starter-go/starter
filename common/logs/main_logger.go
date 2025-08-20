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

	logger  vlog.Logger
	adapter vlog.LoggerAdapter

	startupTime time.Time
}

func (inst *MainLogger) _impl() (application.Lifecycle, vlog.LoggerFactory, vlog.LoggerHolder) {
	return inst, inst, inst
}

func (inst *MainLogger) init() vlog.Logger {
	return initLoggerAdapter(&inst.adapter)
}

// Life ...
func (inst *MainLogger) Life() *application.Life {

	inst.init()

	return &application.Life{
		Order:       common.StartupOrderLogger,
		OnCreate:    inst.onInit1,
		OnStartPost: inst.onInit2,
	}
}

func (inst *MainLogger) onInit1() error {

	steps := make([]func() error, 0)
	steps = append(steps, inst.initThisLogger)
	steps = append(steps, inst.applyThisLogger)
	steps = append(steps, inst.flushStartupLogs)

	for _, step := range steps {
		err := step()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *MainLogger) onInit2() error {
	inst.logStartingInfo()
	return nil
}

func (inst *MainLogger) innerCreateLogger() (vlog.Logger, error) {

	tag := "MainLogger"

	// level
	level, err := vlog.ParseLevel(inst.Level)
	if err != nil {
		return nil, err
	}

	// chain
	cm := &ChainMaker{}
	chain, err := cm.Make(inst.Filters, inst.MainGroupName)
	if err != nil {
		return nil, err
	}

	ada := &inst.adapter
	ada.SetSender(inst)
	ada.SetLevelAccepted(level)
	ada.SetTag(tag)
	ada.SetTargetChain(chain)

	return ada, nil
}

func (inst *MainLogger) innerGetLogger() (vlog.Logger, error) {
	l := inst.logger
	if l != nil {
		return l, nil
	}
	l, err := inst.innerCreateLogger()
	if err != nil {
		return l, err
	}
	inst.logger = l
	return l, nil
}

func (inst *MainLogger) initThisLogger() error {

	_, err := inst.innerGetLogger()

	// if err != nil {
	// 	vlog.Error("starter:MainLogger:error: %v", err.Error())
	// }

	return err
}

func (inst *MainLogger) applyThisLogger() error {
	vlog.SetLoggerFactory(inst)
	return nil
}

func (inst *MainLogger) logStartingInfo() {

	// like 'Started MyApplication in 0.906 seconds (process running for 6.514)'
	const f = "Started app(%s) in %s"

	lv := vlog.INFO
	t0 := inst.startupTime
	t1 := time.Now()
	diff := t1.Sub(t0)
	appName := inst.Context.GetMainModule().Name()

	inst.logger.Log(lv, f, appName, diff.String())
}

func (inst *MainLogger) flushStartupLogs() error {

	const name = attrNameForStartupBuffer
	atts := inst.Context.GetAttributes()
	object := atts.GetAttribute(name)
	atts.SetAttribute(name, nil)

	buffer := object.(*myStartupBuffer)
	src := buffer.logs
	buffer.logs = nil
	inst.startupTime = buffer.startupTime
	for _, msg := range src {
		inst.logger.HandleMessage(msg)
	}

	return nil
}

// Logger ...
func (inst *MainLogger) Logger() vlog.Logger {
	return inst.Create()
}

// Create ...
func (inst *MainLogger) Create() vlog.Logger {
	l, _ := inst.innerGetLogger()
	if l == nil {
		l = &inst.adapter
	}
	return l
}
