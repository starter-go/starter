package demo4starter
import (
    p0ef6f2938 "github.com/starter-go/application"
    p9c772fefa "github.com/starter-go/starter/src/demo/parts"
     "github.com/starter-go/application"
)

// type p9c772fefa.Demo1 in package:github.com/starter-go/starter/src/demo/parts
//
// id:com-9c772fefab3c3f46-parts-Demo1
// class:
// alias:
// scope:singleton
//
type p9c772fefab_parts_Demo1 struct {
}

func (inst* p9c772fefab_parts_Demo1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9c772fefab3c3f46-parts-Demo1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9c772fefab_parts_Demo1) new() any {
    return &p9c772fefa.Demo1{}
}

func (inst* p9c772fefab_parts_Demo1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9c772fefa.Demo1)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)


    return nil
}


func (inst*p9c772fefab_parts_Demo1) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


