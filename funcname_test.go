package funcname_test

import (
	"fmt"
	"net/http/httputil"

	"github.com/takumakei/go-funcname"
)

func Example() {
	name, pkgname := funcname.Of(httputil.DumpRequest)
	fmt.Printf("%q,%q\n", pkgname, name)
	name, pkgname = funcname.This()
	fmt.Printf("%q,%q\n", pkgname, name)
	// Output:
	// "net/http/httputil","DumpRequest"
	// "github.com/takumakei/go-funcname_test","Example"
}
