package main

import (
	"flag"
	"fmt"

	"gostudy/sub"
)

var version = "1.0.0"

func main__() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	var (
		i = flag.Int("iflag", 0, "int flag")
		s = flag.String("sflag", "default", "string flag")
		b = flag.Bool("bflag", false, "bool flag")
	)

	flag.Parse()
	if showVersion {
		fmt.Println("version:", version)
		return
	}

	fmt.Println(*i, *s, *b)
}

// ------------------------------------------------------------------
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {

	var result int = sub.Add(3, sub.Subhoge1)
	fmt.Println("result", result)

	var arr []string = []string{"Golang", "Java", "php", "perl"}
	s := sub.Join(arr, ".")
	fmt.Println(s)

	var i I = T{"hello"}
	i.M()

}
