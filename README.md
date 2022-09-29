funcname
======================================================================

[![Go Reference](https://pkg.go.dev/badge/github.com/takumakei/go-funcname/v2.svg)](https://pkg.go.dev/github.com/takumakei/go-funcname/v2)

Package funcname provides functions getting name and package name of a function.

```text
$ go doc -all
package funcname // import "github.com/takumakei/go-funcname/v2"

Package funcname provides functions getting name and package name of a function.

FUNCTIONS

func Caller(skip int) string
    Caller wraps ForPC(pc) where pc is taken from runtime.Caller(skip + 1).

func ForPC(pc uintptr) string
    ForPC wraps runtime.FuncForPC(pc).Name().

func Of(v interface{}) string
    Of wraps ForPC(pc) where pc is taken from reflect.ValueOf(v).Pointer().

func Split(s string) (pkgname, name string)
    Split splits s into two strings by the first '.' after the last '/'. In case
    of finding no '/', Split splits s by the first '.'. In case of finding no
    '.', Split returns pkgname = s and name = "".

func SplitCaller(skip int) (pkgname, name string)
    SplitCaller wraps Split(Caller(skip + 1)).

func SplitForPC(pc uintptr) (pkgname, name string)
    SplitForPC wraps Split(ForPC(pc)).

func SplitOf(v interface{}) (pkgname, name string)
    SplitOf wraps Split(Of(v)).

func SplitThis() (pkgname, name string)
    SplitThis wraps Split(Caller(1)).

func This() string
    This calls Caller(1).

$
```
