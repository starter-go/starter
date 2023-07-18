package banner

import (
	"strconv"
	"strings"
	"text/template"

	"github.com/starter-go/application"
	"github.com/starter-go/starter/common"
	"github.com/starter-go/vlog"
)

// Banner 用于打印启动条幅
type Banner struct {

	//starter:component

	Context application.Context //starter:inject("context")
	Logger  vlog.Logger         //starter:inject("#")
}

// Life ...
func (inst *Banner) Life() *application.Life {
	return &application.Life{
		Order:    common.StartupOrderBanner,
		OnCreate: inst.init,
	}
}

func (inst *Banner) init() error {
	return inst.printBanner()
}

func (inst *Banner) loadText() (string, error) {

	const path = "/banner.txt"
	text, err := inst.Context.GetResources().ReadText(path)
	if err != nil {
		return "", err
	}

	templ, err := template.New(path).Parse(text)
	if err != nil {
		return text, err
	}

	data := inst.Context.GetProperties().Export(nil)
	data = inst.rewriteMapKeys(data)
	data["starter_version"] = inst.getStarterVersion()

	builder := &strings.Builder{}
	builder.WriteString("banner\n")
	err = templ.Execute(builder, data)
	if err != nil {
		return text, err
	}

	text = builder.String()
	return text, nil
}

func (inst *Banner) rewriteMapKeys(src map[string]string) map[string]string {
	dst := make(map[string]string)
	for k, v := range src {
		k2 := inst.rewriteMapKey(k)
		dst[k2] = v
	}
	return dst
}

func (inst *Banner) rewriteMapKey(key string) string {
	key = strings.ReplaceAll(key, ".", "_")
	key = strings.ReplaceAll(key, "-", "_")
	return key
}

func (inst *Banner) printBanner() error {
	text, err := inst.loadText()
	if err != nil {
		inst.Logger.Warn("%v", err)
	}
	if text != "" {
		inst.Logger.Info("%s", text)
	}
	return nil
}

func (inst *Banner) getStarterVersion() string {
	const moduleName = "github.com/starter-go/starter"
	mods := inst.Context.GetModules()
	for _, m := range mods {
		if m.Name() == moduleName {
			ver := m.Version()
			rev := strconv.Itoa(m.Revision())
			return ver + "-r" + rev
		}
	}
	return "<no value>"
}
