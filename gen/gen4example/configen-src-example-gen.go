package gen4example
import (
    p0ef6f2938 "github.com/starter-go/application"
    p744545790 "github.com/starter-go/starter/src/example/parts"
     "github.com/starter-go/application"
)

// type p744545790.ExampleApp in package:github.com/starter-go/starter/src/example/parts
//
// id:com-7445457906390f52-parts-ExampleApp
// class:
// alias:
// scope:singleton
//
type p7445457906_parts_ExampleApp struct {
}

func (inst* p7445457906_parts_ExampleApp) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7445457906390f52-parts-ExampleApp"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7445457906_parts_ExampleApp) new() any {
    return &p744545790.ExampleApp{}
}

func (inst* p7445457906_parts_ExampleApp) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p744545790.ExampleApp)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)


    return nil
}


func (inst*p7445457906_parts_ExampleApp) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


