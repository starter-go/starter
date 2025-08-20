//starter:configen(version="4")

package main4starter

import "github.com/starter-go/application"

// ExportComponents 导出组件配置
func ExportComponents(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
