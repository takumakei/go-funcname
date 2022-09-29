package funcname_test

import (
	"fmt"
	"testing"

	"github.com/takumakei/go-funcname/v3"
)

func ExampleSplit() {
	for _, v := range []string{
		"main.hello",
		"net.Dial",
		"net/http.ListenAndServe",
		"net/http/httputil.DumpRequest",
		"example/only",
	} {
		pkgname, name := funcname.Split(v)
		fmt.Printf("%q,%q\n", pkgname, name)
	}
	// Output:
	// "main","hello"
	// "net","Dial"
	// "net/http","ListenAndServe"
	// "net/http/httputil","DumpRequest"
	// "example/only",""
}

func TestSplit(t *testing.T) {
	cases := []struct {
		In, Lhs, Rhs string
	}{
		{In: "", Lhs: "", Rhs: ""},
		{In: ".", Lhs: "", Rhs: ""},
		{In: "/", Lhs: "/", Rhs: ""},
		{In: "/.", Lhs: "/", Rhs: ""},
		{In: "./", Lhs: "./", Rhs: ""},
		{In: "./.", Lhs: "./", Rhs: ""},
		{In: "hello", Lhs: "hello", Rhs: ""},
		{In: "main.hello", Lhs: "main", Rhs: "hello"},
		{In: "net.Dial", Lhs: "net", Rhs: "Dial"},
		{In: "net/http.ListenAndServe", Lhs: "net/http", Rhs: "ListenAndServe"},
		{In: "net/http/httputil.DumpRequest", Lhs: "net/http/httputil", Rhs: "DumpRequest"},
	}
	for i, c := range cases {
		lhs, rhs := funcname.Split(c.In)
		if lhs != c.Lhs || rhs != c.Rhs {
			t.Errorf(
				"[%d] funcname.Split(%q) -> %q,%q, want %q,%q",
				i, c.In, lhs, rhs, c.Lhs, c.Rhs,
			)
		} else if testing.Verbose() {
			t.Logf(
				"[%d] funcname.Split(%q) -> %q,%q",
				i, c.In, lhs, rhs,
			)
		}
	}
}
