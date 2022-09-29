package test

import "github.com/takumakei/go-funcname/v3"

func This() string {
	return funcname.This()
}

func SplitThis() (pkgname, name string) {
	return funcname.SplitThis()
}

func Caller(skip int) string {
	return funcname.Caller(skip)
}

func SplitCaller(skip int) (pkgname, name string) {
	return funcname.SplitCaller(skip)
}
