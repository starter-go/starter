package starter

import (
	"github.com/starter-go/application"
	"github.com/starter-go/application/boot"
)

// Initializer 表示应用程序的初始化对象
type Initializer interface {
	MainModule(mod application.Module) Initializer
	WithPanic(enable bool) Initializer
	Run() error
}

// Init 返回一个新的应用初始化工具
func Init(args []string) Initializer {
	if args == nil {
		args = make([]string, 0)
	}
	i := &myInitializer{}
	i.options.Args = args
	return i
}

////////////////////////////////////////////////////////////////////////////////

type myInitializer struct {
	main        application.Module // 主模块
	enablePanic bool
	options     boot.Options
}

func (inst *myInitializer) MainModule(mod application.Module) Initializer {
	inst.main = mod
	return inst
}

func (inst *myInitializer) WithPanic(enable bool) Initializer {
	inst.enablePanic = enable
	return inst
}

func (inst *myInitializer) Run() error {
	mod := inst.main
	opt := &inst.options
	err := boot.Run(mod, opt)
	if err != nil {
		if inst.enablePanic {
			panic(err)
		}
	}
	return err
}
