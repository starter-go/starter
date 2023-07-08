package main

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
	"github.com/starter-go/starter/gen/gen4example"
)

//go:embed "res"
var theModuleResFS embed.FS

func theModule() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name("starter/example").Version("v1").Revision(1)
	mb.EmbedResources(theModuleResFS, "res")
	mb.Depend(starter.Module())
	mb.Components(gen4example.ExportComponents)
	return mb.Create()
}
