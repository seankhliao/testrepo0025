package main

import (
	"fmt"

	"go.seankhliao.com/testrepo0025/foo"
)

func main() {
	fmt.Println(foo.Foo)
	foo.Foo += 1
	fmt.Println(foo.Foo)
}
