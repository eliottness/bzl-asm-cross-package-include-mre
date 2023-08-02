package main

import (
	"fmt"
	"github.com/eliottness/purego"
)

func main() {
	_, err := purego.Dlopen("non-existant-lib.so", purego.RTLD_GLOBAL | purego.RTLD_NOW)
	if err == nil {
		panic("wtf")
	} else {
		fmt.Println("everything is fine")
	}
}
