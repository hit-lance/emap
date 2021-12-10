// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type foo struct {
	a int
	b string
}

func main() {
	var f foo
	p := &f.b
	f.a = 1
	f.b = "12"
	f = foo{}
	fmt.Printf("%p\n",p)
	fmt.Printf("%p\n",&f.b)
}
