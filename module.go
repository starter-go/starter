package starter

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter/gen/gen4starter"
)

const (
	theModuleName     = "github.com/starter-go/starter"
	theModuleVersion  = "v1.0.0"
	theModuleRevision = 1
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// Module 导出模块 [github.com/starter-go/starter]
func Module() application.Module {
	mb := application.ModuleBuilder{}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(gen4starter.ExportComponents)
	mb.Depend(nil)
	return mb.Create()
}
