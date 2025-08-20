package starter

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter/gen/demo4starter"
	"github.com/starter-go/starter/gen/main4starter"
	"github.com/starter-go/starter/gen/test4starter"
)

const (
	theModuleName     = "github.com/starter-go/starter"
	theModuleVersion  = "v1.0.13"
	theModuleRevision = 15
)

////////////////////////////////////////////////////////////////////////////////

const (
	theSrcMainModuleResPath = "src/main/resources"
	theSrcTestModuleResPath = "src/test/resources"
	theSrcDemoModuleResPath = "src/demo/resources"
)

//go:embed "src/main/resources"
var theSrcMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theSrcTestModuleResFS embed.FS

//go:embed "src/demo/resources"
var theSrcDemoModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// Module 导出模块 (#main) [github.com/starter-go/starter]
func Module() application.Module {
	mb := application.ModuleBuilder{}
	mb.Name(theModuleName + "#main")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theSrcMainModuleResFS, theSrcMainModuleResPath)
	mb.Components(main4starter.ExportComponents)
	return mb.Create()
}

// Module 导出模块 (#test) [github.com/starter-go/starter]
func ModuleForTest() application.Module {
	mb := application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theSrcTestModuleResFS, theSrcTestModuleResPath)
	mb.Components(test4starter.ExportComponents)

	mb.Depend(Module())

	return mb.Create()
}

// Module 导出模块 (#demo) [github.com/starter-go/starter]
func ModuleForDemo() application.Module {
	mb := application.ModuleBuilder{}
	mb.Name(theModuleName + "#demo")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theSrcDemoModuleResFS, theSrcDemoModuleResPath)
	mb.Components(demo4starter.ExportComponents)

	mb.Depend(Module())

	return mb.Create()
}
