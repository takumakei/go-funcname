package funcname_test

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"reflect"

	"github.com/takumakei/go-funcname/v2"
	"github.com/takumakei/go-funcname/v2/internal/test"
)

func Example() {
	var pkgname, name string
	name = funcname.Of(httputil.DumpRequest)
	fmt.Printf("%q\n", name)
	pkgname, name = funcname.Split(name)
	fmt.Printf("%q, %q\n", pkgname, name)
	pkgname, name = funcname.SplitOf(httputil.DumpRequest)
	fmt.Printf("%q, %q\n", pkgname, name)
	pkgname, name = test.SplitThis()
	fmt.Printf("%q, %q\n", pkgname, name)
	// Output:
	// "net/http/httputil.DumpRequest"
	// "net/http/httputil", "DumpRequest"
	// "net/http/httputil", "DumpRequest"
	// "github.com/takumakei/go-funcname/v2/internal/test", "SplitThis"
}

func ExampleOf() {
	name := funcname.Of(http.ListenAndServe)
	fmt.Printf("%q", name)
	// Output:
	// "net/http.ListenAndServe"
}

func ExampleSplitOf() {
	pkgname, name := funcname.SplitOf(http.ListenAndServe)
	fmt.Printf("%q, %q", pkgname, name)
	// Output:
	// "net/http", "ListenAndServe"
}

func ExampleThis() {
	name := test.This()
	fmt.Printf("%q", name)
	// Output:
	// "github.com/takumakei/go-funcname/v2/internal/test.This"
}

func ExampleSplitThis() {
	pkgname, name := test.SplitThis()
	fmt.Printf("%q, %q", pkgname, name)
	// Output:
	// "github.com/takumakei/go-funcname/v2/internal/test", "SplitThis"
}

func ExampleCaller() {
	name := test.Caller(0)
	fmt.Printf("%q", name)
	// Output:
	// "github.com/takumakei/go-funcname/v2/internal/test.Caller"
}

func ExampleSplitCaller() {
	pkgname, name := test.SplitCaller(0)
	fmt.Printf("%q, %q", pkgname, name)
	// Output:
	// "github.com/takumakei/go-funcname/v2/internal/test", "SplitCaller"
}

func ExampleForPC() {
	pc := reflect.ValueOf(http.ListenAndServe).Pointer()
	name := funcname.ForPC(pc)
	fmt.Printf("%q", name)
	// Output:
	// "net/http.ListenAndServe"
}

func ExampleSplitForPC() {
	pc := reflect.ValueOf(http.ListenAndServe).Pointer()
	pkgname, name := funcname.SplitForPC(pc)
	fmt.Printf("%q, %q", pkgname, name)
	// Output:
	// "net/http", "ListenAndServe"
}
