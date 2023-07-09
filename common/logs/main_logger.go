package logs

import (
	"strings"
	"text/template"
	"time"

	"github.com/starter-go/application"
	"github.com/starter-go/vlog"
)

// MainLogger 用于向文件输出日志
type MainLogger struct {
	//starter:component

	Filters       []vlog.MessageFilterRegistry //starter:inject(".")
	Level         string                       //starter:inject("${vlog.level}")
	MainGroupName string                       //starter:inject("${vlog.main}")
	Context       application.Context          //starter:inject("context")

	lc          *LoggerContext //  level vlog.Level
	startupTime time.Time
}

func (inst *MainLogger) _impl() (application.Lifecycle, vlog.LoggerFactory) {
	return inst, inst
}

// Life ...
func (inst *MainLogger) Life() *application.Life {
	return &application.Life{
		Order:       -10000,
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
	inst.printBanner()
	return nil
}

func (inst *MainLogger) onStarted() error {

	// like 'Started MyApplication in 0.906 seconds (process running for 6.514)'
	const f = "Started %s in %s"

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

func (inst *MainLogger) printBanner() error {
	const path = "/banner.txt"
	res, err := inst.Context.GetResources().GetResource(path)
	if err != nil {
		return err
	}
	textTemplate, err := res.ReadText()
	if err != nil {
		return err
	}

	data := make(map[string]string)
	data["StarterVersion"] = "6666"

	templ, err := template.New(path).Parse(textTemplate)
	if err != nil {
		return err
	}
	builder := &strings.Builder{}
	builder.WriteString("banner\n")
	err = templ.Execute(builder, data)
	if err != nil {
		return err
	}
	inst.lc.LogWithLevel(vlog.INFO, builder.String())
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
