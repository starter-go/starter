package logs

import "github.com/starter-go/vlog"

// MainGroup  这是 vlog 的主要过滤器组
type MainGroup struct {
	//starter:component
	_as func(vlog.MessageFilterRegistry) //starter:as(".")

	Filters        []vlog.MessageFilterRegistry //starter:inject(".")
	FilterNameList string                       //starter:inject("${vlog.groups.main.filters}")

	chain vlog.MessageFilterChain
}

func (inst *MainGroup) _impl() (vlog.MessageFilterGroup, vlog.MessageFilterRegistry) {
	return inst, inst
}

// ListLogFilterRegistration ...
func (inst *MainGroup) ListLogFilterRegistration() []*vlog.MessageFilterRegistration {
	filter := inst
	r1 := &vlog.MessageFilterRegistration{
		Name:   "main",
		Order:  0,
		Filter: filter,
	}
	return []*vlog.MessageFilterRegistration{r1}
}

// GetFilterChain ...
func (inst *MainGroup) GetFilterChain() vlog.MessageFilterChain {
	chain := inst.chain
	if chain == nil {
		c, err := inst.loadChain()
		if err != nil {
			panic(err)
		}
		chain = c
		inst.chain = c
	}
	return chain
}

func (inst *MainGroup) dispatch(msg *vlog.Message) {
	chain := inst.GetFilterChain()
	chain.DoFilter(msg)
}

func (inst *MainGroup) loadChain() (vlog.MessageFilterChain, error) {
	cm := &ChainMaker{}
	return cm.Make(inst.Filters, inst.FilterNameList)
}

// DoFilter ...
func (inst *MainGroup) DoFilter(msg *vlog.Message, chain vlog.MessageFilterChain) {
	inst.dispatch(msg)
	chain.DoFilter(msg)
}
