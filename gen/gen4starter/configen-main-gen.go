package gen4starter

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p10812a4aae_debug_ContextInfoLog{})
    inst.register(&p11352c8b40_logs_ConsoleWriter{})
    inst.register(&p11352c8b40_logs_DefaultFormatter{})
    inst.register(&p11352c8b40_logs_FileWriter{})
    inst.register(&p11352c8b40_logs_LoggerProxy{})
    inst.register(&p11352c8b40_logs_MainGroup{})
    inst.register(&p11352c8b40_logs_MainLogger{})
    inst.register(&p753b8c7726_fs_AFSProxy{})
    inst.register(&pa5abee2538_common_UUIDGenService{})
    inst.register(&pb8e5d8aedf_banner_Banner{})


    return nil
}
