package parts

import (
	"github.com/starter-go/application"
	"github.com/starter-go/vlog"
)

// Demo1 为示例应用
type Demo1 struct {
	//starter:component

	AC application.Context //starter:inject("context")
}

// Life 返回生命周期注册信息
func (inst *Demo1) Life() *application.Life {
	return &application.Life{
		OnCreate:  inst.create,
		OnStart:   inst.start,
		OnLoop:    inst.runLoop,
		OnStop:    inst.stop,
		OnDestroy: inst.destroy,
	}
}

func (inst *Demo1) logMyself(fn string) {
	vlog.Info("Demo1." + fn)
}

func (inst *Demo1) create() error {
	inst.logMyself("OnCreate()")
	return nil
}

func (inst *Demo1) start() error {
	inst.logMyself("OnStart()")
	return nil
}

func (inst *Demo1) runLoop() error {
	inst.logMyself("OnLoop()")

	vlog.Info("com-list:")
	ids := inst.AC.GetComponents().ListIDs()
	for _, id := range ids {
		vlog.Info("  com: %s", id)
	}
	return nil
}

func (inst *Demo1) stop() error {
	inst.logMyself("OnStop()")
	return nil
}

func (inst *Demo1) destroy() error {
	inst.logMyself("OnDestroy()")
	return nil
}
