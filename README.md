funcname
======================================================================

[![Go Reference](https://pkg.go.dev/badge/github.com/takumakei/go-funcname.svg)](https://pkg.go.dev/github.com/takumakei/go-funcname)

Package funcname provides functions getting name and package name of a function.

```text
$ go doc -all
package funcname // import "github.com/takumakei/go-funcname"

Package funcname provides functions getting name and package name of a
function.

FUNCTIONS

func ForPC(pc uintptr) (name, pkgname string)
    ForPC returns the name and the package name of the function from the pc.

func Of(v interface{}) (name, pkgname string)
    Of returns the name and the package name of the function v.

func Split(s string) (lhs, rhs string)
    Split splits s into two strings by the first '.' after the last '/'.

func This() (name, pkgname string)
    This returns the name and the package name of the function calling This.

$
```
