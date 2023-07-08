package parts

import (
	"github.com/starter-go/application"
	"github.com/starter-go/vlog"
)

// ExampleApp 为示例应用
type ExampleApp struct {
	//starter:component

	AC application.Context //starter:inject("context")
}

// Life 返回生命周期注册信息
func (inst *ExampleApp) Life() *application.Life {
	return &application.Life{
		OnLoop: inst.loop,
	}
}

func (inst *ExampleApp) loop() error {

	vlog.Info("ExampleApp.loop()")
	vlog.Info("com-list:")

	ids := inst.AC.GetComponents().ListIDs()
	for _, id := range ids {
		vlog.Info("  com: %s", id)
	}

	return nil
}
