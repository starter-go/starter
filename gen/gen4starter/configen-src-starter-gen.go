package gen4starter
import (
    p0ef6f2938 "github.com/starter-go/application"
    pa5abee253 "github.com/starter-go/starter/common"
    pb8e5d8aed "github.com/starter-go/starter/common/banner"
    p10812a4aa "github.com/starter-go/starter/common/debug"
    p753b8c772 "github.com/starter-go/starter/common/fs"
    p11352c8b4 "github.com/starter-go/starter/common/logs"
    p55f0853be "github.com/starter-go/vlog"
     "github.com/starter-go/application"
)

// type pa5abee253.UUIDGenService in package:github.com/starter-go/starter/common
//
// id:com-a5abee25387b059a-common-UUIDGenService
// class:
// alias:alias-3f0a0f8fe6baff1831e07bb3c7c57e6b-UUIDGenerator
// scope:singleton
//
type pa5abee2538_common_UUIDGenService struct {
}

func (inst* pa5abee2538_common_UUIDGenService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a5abee25387b059a-common-UUIDGenService"
	r.Classes = ""
	r.Aliases = "alias-3f0a0f8fe6baff1831e07bb3c7c57e6b-UUIDGenerator"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa5abee2538_common_UUIDGenService) new() any {
    return &pa5abee253.UUIDGenService{}
}

func (inst* pa5abee2538_common_UUIDGenService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa5abee253.UUIDGenService)
	nop(ie, com)

	


    return nil
}



// type pb8e5d8aed.Banner in package:github.com/starter-go/starter/common/banner
//
// id:com-b8e5d8aedf6fd10c-banner-Banner
// class:
// alias:
// scope:singleton
//
type pb8e5d8aedf_banner_Banner struct {
}

func (inst* pb8e5d8aedf_banner_Banner) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b8e5d8aedf6fd10c-banner-Banner"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb8e5d8aedf_banner_Banner) new() any {
    return &pb8e5d8aed.Banner{}
}

func (inst* pb8e5d8aedf_banner_Banner) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb8e5d8aed.Banner)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.Logger = inst.getLogger(ie)


    return nil
}


func (inst*pb8e5d8aedf_banner_Banner) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pb8e5d8aedf_banner_Banner) getLogger(ie application.InjectionExt)p55f0853be.Logger{
    return ie.GetComponent("#alias-55f0853bedbc094981acd8da904ae269-Logger").(p55f0853be.Logger)
}



// type p10812a4aa.ContextInfoLog in package:github.com/starter-go/starter/common/debug
//
// id:com-10812a4aaeee52a6-debug-ContextInfoLog
// class:
// alias:
// scope:singleton
//
type p10812a4aae_debug_ContextInfoLog struct {
}

func (inst* p10812a4aae_debug_ContextInfoLog) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-10812a4aaeee52a6-debug-ContextInfoLog"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p10812a4aae_debug_ContextInfoLog) new() any {
    return &p10812a4aa.ContextInfoLog{}
}

func (inst* p10812a4aae_debug_ContextInfoLog) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p10812a4aa.ContextInfoLog)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)
    com.Logger = inst.getLogger(ie)
    com.EnableDebug = inst.getEnableDebug(ie)
    com.LogArgs = inst.getLogArgs(ie)
    com.LogEnv = inst.getLogEnv(ie)
    com.LogProps = inst.getLogProps(ie)


    return nil
}


func (inst*p10812a4aae_debug_ContextInfoLog) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p10812a4aae_debug_ContextInfoLog) getLogger(ie application.InjectionExt)p55f0853be.Logger{
    return ie.GetComponent("#alias-55f0853bedbc094981acd8da904ae269-Logger").(p55f0853be.Logger)
}


func (inst*p10812a4aae_debug_ContextInfoLog) getEnableDebug(ie application.InjectionExt)bool{
    return ie.GetBool("${debug.enabled}")
}


func (inst*p10812a4aae_debug_ContextInfoLog) getLogArgs(ie application.InjectionExt)bool{
    return ie.GetBool("${debug.log-arguments}")
}


func (inst*p10812a4aae_debug_ContextInfoLog) getLogEnv(ie application.InjectionExt)bool{
    return ie.GetBool("${debug.log-environment}")
}


func (inst*p10812a4aae_debug_ContextInfoLog) getLogProps(ie application.InjectionExt)bool{
    return ie.GetBool("${debug.log-properties}")
}



// type p753b8c772.AFSProxy in package:github.com/starter-go/starter/common/fs
//
// id:com-753b8c77262f4ce1-fs-AFSProxy
// class:
// alias:alias-0d2a11d163e349503a64168a1cdf48a2-FS
// scope:singleton
//
type p753b8c7726_fs_AFSProxy struct {
}

func (inst* p753b8c7726_fs_AFSProxy) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-753b8c77262f4ce1-fs-AFSProxy"
	r.Classes = ""
	r.Aliases = "alias-0d2a11d163e349503a64168a1cdf48a2-FS"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p753b8c7726_fs_AFSProxy) new() any {
    return &p753b8c772.AFSProxy{}
}

func (inst* p753b8c7726_fs_AFSProxy) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p753b8c772.AFSProxy)
	nop(ie, com)

	


    return nil
}



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



// type p11352c8b4.LoggerProxy in package:github.com/starter-go/starter/common/logs
//
// id:com-11352c8b402dcccb-logs-LoggerProxy
// class:
// alias:alias-55f0853bedbc094981acd8da904ae269-Logger
// scope:singleton
//
type p11352c8b40_logs_LoggerProxy struct {
}

func (inst* p11352c8b40_logs_LoggerProxy) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-11352c8b402dcccb-logs-LoggerProxy"
	r.Classes = ""
	r.Aliases = "alias-55f0853bedbc094981acd8da904ae269-Logger"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p11352c8b40_logs_LoggerProxy) new() any {
    return &p11352c8b4.LoggerProxy{}
}

func (inst* p11352c8b40_logs_LoggerProxy) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p11352c8b4.LoggerProxy)
	nop(ie, com)

	
    com.Holder = inst.getHolder(ie)


    return nil
}


func (inst*p11352c8b40_logs_LoggerProxy) getHolder(ie application.InjectionExt)p55f0853be.LoggerHolder{
    return ie.GetComponent("#starter-main-logger-holder").(p55f0853be.LoggerHolder)
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
// alias:starter-main-logger-holder
// scope:singleton
//
type p11352c8b40_logs_MainLogger struct {
}

func (inst* p11352c8b40_logs_MainLogger) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-11352c8b402dcccb-logs-MainLogger"
	r.Classes = ""
	r.Aliases = "starter-main-logger-holder"
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
    com.Context = inst.getContext(ie)


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


func (inst*p11352c8b40_logs_MainLogger) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


