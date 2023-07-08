package gen4starter
import (
    p11352c8b4 "github.com/starter-go/starter/common/logs"
    p55f0853be "github.com/starter-go/vlog"
     "github.com/starter-go/application"
)

// type p11352c8b4.ConsoleWriter in package:github.com/starter-go/starter/common/logs
//
// id:com-11352c8b402dcccb-logs-ConsoleWriter
// class:class-55f0853bedbc094981acd8da904ae269-MessageFilterRegistry
// alias:
// scope:singleton
//
type p11352c8b40_logs_ConsoleWriter struct {
}

func (inst* p11352c8b40_logs_ConsoleWriter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-11352c8b402dcccb-logs-ConsoleWriter"
	r.Classes = "class-55f0853bedbc094981acd8da904ae269-MessageFilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p11352c8b40_logs_ConsoleWriter) new() any {
    return &p11352c8b4.ConsoleWriter{}
}

func (inst* p11352c8b40_logs_ConsoleWriter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p11352c8b4.ConsoleWriter)
	nop(ie, com)

	


    return nil
}



// type p11352c8b4.DefaultFormatter in package:github.com/starter-go/starter/common/logs
//
// id:com-11352c8b402dcccb-logs-DefaultFormatter
// class:class-55f0853bedbc094981acd8da904ae269-MessageFilterRegistry
// alias:
// scope:singleton
//
type p11352c8b40_logs_DefaultFormatter struct {
}

func (inst* p11352c8b40_logs_DefaultFormatter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-11352c8b402dcccb-logs-DefaultFormatter"
	r.Classes = "class-55f0853bedbc094981acd8da904ae269-MessageFilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p11352c8b40_logs_DefaultFormatter) new() any {
    return &p11352c8b4.DefaultFormatter{}
}

func (inst* p11352c8b40_logs_DefaultFormatter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p11352c8b4.DefaultFormatter)
	nop(ie, com)

	
    com.Format = inst.getFormat(ie)


    return nil
}


func (inst*p11352c8b40_logs_DefaultFormatter) getFormat(ie application.InjectionExt)string{
    return ie.GetString("${vlog.formatters.default.format}")
}



// type p11352c8b4.FileWriter in package:github.com/starter-go/starter/common/logs
//
// id:com-11352c8b402dcccb-logs-FileWriter
// class:
// alias:
// scope:singleton
//
type p11352c8b40_logs_FileWriter struct {
}

func (inst* p11352c8b40_logs_FileWriter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-11352c8b402dcccb-logs-FileWriter"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p11352c8b40_logs_FileWriter) new() any {
    return &p11352c8b4.FileWriter{}
}

func (inst* p11352c8b40_logs_FileWriter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p11352c8b4.FileWriter)
	nop(ie, com)

	


    return nil
}



// type p11352c8b4.MainGroup in package:github.com/starter-go/starter/common/logs
//
// id:com-11352c8b402dcccb-logs-MainGroup
// class:class-55f0853bedbc094981acd8da904ae269-MessageFilterRegistry
// alias:
// scope:singleton
//
type p11352c8b40_logs_MainGroup struct {
}

func (inst* p11352c8b40_logs_MainGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-11352c8b402dcccb-logs-MainGroup"
	r.Classes = "class-55f0853bedbc094981acd8da904ae269-MessageFilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p11352c8b40_logs_MainGroup) new() any {
    return &p11352c8b4.MainGroup{}
}

func (inst* p11352c8b40_logs_MainGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p11352c8b4.MainGroup)
	nop(ie, com)

	
    com.Filters = inst.getFilters(ie)
    com.FilterNameList = inst.getFilterNameList(ie)


    return nil
}


func (inst*p11352c8b40_logs_MainGroup) getFilters(ie application.InjectionExt)[]p55f0853be.MessageFilterRegistry{
    dst := make([]p55f0853be.MessageFilterRegistry, 0)
    src := ie.ListComponents(".class-55f0853bedbc094981acd8da904ae269-MessageFilterRegistry")
    for _, item1 := range src {
        item2 := item1.(p55f0853be.MessageFilterRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p11352c8b40_logs_MainGroup) getFilterNameList(ie application.InjectionExt)string{
    return ie.GetString("${vlog.groups.main.filters}")
}



// type p11352c8b4.MainLogger in package:github.com/starter-go/starter/common/logs
//
// id:com-11352c8b402dcccb-logs-MainLogger
// class:
// alias:
// scope:singleton
//
type p11352c8b40_logs_MainLogger struct {
}

func (inst* p11352c8b40_logs_MainLogger) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-11352c8b402dcccb-logs-MainLogger"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p11352c8b40_logs_MainLogger) new() any {
    return &p11352c8b4.MainLogger{}
}

func (inst* p11352c8b40_logs_MainLogger) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p11352c8b4.MainLogger)
	nop(ie, com)

	
    com.Filters = inst.getFilters(ie)
    com.Level = inst.getLevel(ie)
    com.MainGroupName = inst.getMainGroupName(ie)


    return nil
}


func (inst*p11352c8b40_logs_MainLogger) getFilters(ie application.InjectionExt)[]p55f0853be.MessageFilterRegistry{
    dst := make([]p55f0853be.MessageFilterRegistry, 0)
    src := ie.ListComponents(".class-55f0853bedbc094981acd8da904ae269-MessageFilterRegistry")
    for _, item1 := range src {
        item2 := item1.(p55f0853be.MessageFilterRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p11352c8b40_logs_MainLogger) getLevel(ie application.InjectionExt)string{
    return ie.GetString("${vlog.level}")
}


func (inst*p11352c8b40_logs_MainLogger) getMainGroupName(ie application.InjectionExt)string{
    return ie.GetString("${vlog.main}")
}


