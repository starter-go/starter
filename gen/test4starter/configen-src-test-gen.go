package test4starter
import (
    pbb8612058 "github.com/starter-go/starter/src/test/units"
     "github.com/starter-go/application"
)

// type pbb8612058.Unit1 in package:github.com/starter-go/starter/src/test/units
//
// id:com-bb861205875d89f2-units-Unit1
// class:
// alias:
// scope:singleton
//
type pbb86120587_units_Unit1 struct {
}

func (inst* pbb86120587_units_Unit1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-bb861205875d89f2-units-Unit1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pbb86120587_units_Unit1) new() any {
    return &pbb8612058.Unit1{}
}

func (inst* pbb86120587_units_Unit1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pbb8612058.Unit1)
	nop(ie, com)

	


    return nil
}


