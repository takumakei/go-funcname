// Package funcname provides functions getting name and package name of a function.
package funcname

import (
	"reflect"
	"runtime"
)

// Of returns the name and the package name of the function v.
func Of(v interface{}) (name, pkgname string) {
	return ForPC(reflect.ValueOf(v).Pointer())
}

// This returns the name and the package name of the function calling This.
func This() (name, pkgname string) {
	pc, _, _, _ := runtime.Caller(1)
	return ForPC(pc)
}

// ForPC returns the name and the package name of the function from the pc.
func ForPC(pc uintptr) (name, pkgname string) {
	pkgname, name = Split(runtime.FuncForPC(pc).Name())
	return
}
