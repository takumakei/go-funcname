// Package funcname provides functions getting name and package name of a function.
package funcname

import (
	"reflect"
	"runtime"
)

// Of wraps ForPC(pc) where pc is taken from reflect.ValueOf(v).Pointer().
func Of(v interface{}) string {
	return ForPC(reflect.ValueOf(v).Pointer())
}

// SplitOf wraps Split(Of(v)).
func SplitOf(v interface{}) (pkgname, name string) {
	return Split(Of(v))
}

// This calls Caller(1).
func This() string {
	return Caller(1)
}

// SplitThis wraps Split(Caller(1)).
func SplitThis() (pkgname, name string) {
	return Split(Caller(1))
}

// Caller wraps ForPC(pc) where pc is taken from runtime.Caller(skip + 1).
func Caller(skip int) string {
	pc, _, _, _ := runtime.Caller(skip + 1)
	return ForPC(pc)
}

// SplitCaller wraps Split(Caller(skip + 1)).
func SplitCaller(skip int) (pkgname, name string) {
	return Split(Caller(skip + 1))
}

// ForPC wraps runtime.FuncForPC(pc).Name().
func ForPC(pc uintptr) string {
	return runtime.FuncForPC(pc).Name()
}

// SplitForPC wraps Split(ForPC(pc)).
func SplitForPC(pc uintptr) (pkgname, name string) {
	return Split(ForPC(pc))
}
