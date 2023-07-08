package starter

import (
	"github.com/starter-go/application"
	"github.com/starter-go/application/arguments"
	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/boot"
	"github.com/starter-go/application/environment"
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/safe"
	"github.com/starter-go/starter/common/logs"
)

// Initializer 表示应用程序的初始化对象
type Initializer interface {
	MainModule(mod application.Module) Initializer
	WithPanic(enable bool) Initializer
	SetArguments(args []string) Initializer

	GetAttributes() attributes.Table
	GetEnvironment() environment.Table
	GetParameters() parameters.Table
	GetProperties() properties.Table

	Run() error
}

// Init 返回一个新的应用初始化工具
func Init(args []string) Initializer {
	i := &myInitializer{}
	i.SetArguments(args)
	i.collections.Complete(nil)

	logbuf := logs.NewStartupBuffer()
	logbuf.Init(i.GetAttributes())

	return i
}

////////////////////////////////////////////////////////////////////////////////

type myInitializer struct {
	main        application.Module // 主模块
	enablePanic bool
	mode        safe.Mode
	collections application.Collections
}

func (inst *myInitializer) GetArguments() arguments.Table {
	return inst.collections.Arguments
}

func (inst *myInitializer) GetAttributes() attributes.Table {
	return inst.collections.Attributes
}

func (inst *myInitializer) GetEnvironment() environment.Table {
	return inst.collections.Environment
}

func (inst *myInitializer) GetParameters() parameters.Table {
	return inst.collections.Parameters
}

func (inst *myInitializer) GetProperties() properties.Table {
	return inst.collections.Properties
}

func (inst *myInitializer) SetArguments(args []string) Initializer {
	inst.collections.Arguments = arguments.NewTable(args, nil)
	return inst
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

	col := &inst.collections
	col.Complete(nil)

	opt := &boot.Options{}
	opt.Mode = inst.mode
	opt.Args = col.Arguments.Raw()
	opt.Attributes = col.Attributes
	opt.Environment = col.Environment
	opt.Properties = col.Properties
	opt.Parameters = col.Parameters

	err := boot.Run(mod, opt)
	if err != nil {
		if inst.enablePanic {
			panic(err)
		}
	}
	return err
}
