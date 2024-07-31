package fs

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/afs/files"
)

// AFSProxy ...
type AFSProxy struct {

	//starter:component
	_as func(afs.FS) //starter:as("#")

	target afs.FS
}

func (inst *AFSProxy) _impl() afs.FS {
	return inst
}

// FS ...
func (inst *AFSProxy) FS() afs.FS {
	fs := inst.target
	if fs == nil {
		fs = files.FS()
		inst.target = fs
	}
	return fs
}

// NewPath ...
func (inst *AFSProxy) NewPath(path string) afs.Path {
	return inst.FS().NewPath(path)
}

// NewURI ...
func (inst *AFSProxy) NewURI(u afs.URI) (afs.Path, error) {
	return inst.FS().NewURI(u)
}

// ListRoots ...
func (inst *AFSProxy) ListRoots() []afs.Path {
	return inst.FS().ListRoots()
}

// CreateTempFile ...
func (inst *AFSProxy) CreateTempFile(prefix, suffix string, dir afs.Path) (afs.Path, error) {
	return inst.FS().CreateTempFile(prefix, suffix, dir)
}

// PathSeparator return ';'(windows) | ':'(unix)
func (inst *AFSProxy) PathSeparator() string {
	return inst.FS().PathSeparator()
}

// Separator return '/'(unix) | '\'(windows)
func (inst *AFSProxy) Separator() string {
	return inst.FS().Separator()
}

// // SetDefaultOptionsHandler 设置一个函数，用来处理默认的I/O选项
// func (inst *AFSProxy) SetDefaultOptionsHandler(fn afs.OptionsHandlerFunc) error {
// 	return inst.FS().SetDefaultOptionsHandler(fn)
// }
