package logs

import (
	"fmt"
	"strings"

	"github.com/starter-go/vlog"
)

// ChainMaker ...
type ChainMaker struct {
	table map[string]*vlog.MessageFilterRegistration
}

// Init ...
func (inst *ChainMaker) init(src []vlog.MessageFilterRegistry) error {
	inst.table = make(map[string]*vlog.MessageFilterRegistration)
	for _, registry := range src {
		registrationList := registry.ListLogFilterRegistration()
		for _, item := range registrationList {
			err := inst.add(item)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (inst *ChainMaker) add(item *vlog.MessageFilterRegistration) error {
	name := strings.TrimSpace(item.Name)
	older := inst.table[name]
	f := item.Filter
	if f == nil {
		return nil
	}
	if older != nil {
		return fmt.Errorf("the vlog.MessageFilter name is duplicate: '%s'", name)
	}
	inst.table[name] = item
	return nil
}

func (inst *ChainMaker) build(namelist []string) (vlog.MessageFilterChain, error) {
	builder := &vlog.MessageFilterChainBuilder{}
	for i := len(namelist) - 1; i >= 0; i-- {
		name := namelist[i]
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		f, err := inst.Find(name)
		if err != nil {
			return nil, err
		}
		builder.AddFilter(f)
	}
	return builder.Create(), nil
}

// Make 创建过滤器链，filterNameList 是一个以 ',' 分隔的名单字符串
func (inst *ChainMaker) Make(src []vlog.MessageFilterRegistry, filterNameList string) (vlog.MessageFilterChain, error) {
	const sep = ","
	namelist := strings.Split(filterNameList, sep)
	err := inst.init(src)
	if err != nil {
		return nil, err
	}
	return inst.build(namelist)
}

// Find ...
func (inst *ChainMaker) Find(name string) (vlog.MessageFilter, error) {

	e := fmt.Errorf("no vlog.MessageFilter with name: '%s'", name)
	t := inst.table
	if t == nil {
		return nil, e
	}

	item := t[name]
	if item == nil {
		return nil, e
	}

	f := item.Filter
	if f == nil {
		return nil, e
	}

	return f, nil
}
