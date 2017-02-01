package nodejs

import (
	"path/filepath"

	"github.com/gopherjs/gopherjs/js"
)

func Require(name string) *js.Object {
	return js.Global.Call("require", name)
}

// FileName returns the file path of the compiled `JS` file.
func FileName() string {
	return js.Module.Get("filename").String()
}

// DirName returns the directory containing the compiled `JS` file.
func DirName() string {
	fn := FileName()
	return filepath.Dir(fn)
}
