package main

import (
	hellomodV1 "github.com/donvito/hellomod"
	"github.com/donvito/hellomod/v2"
)

func main() {
	hellomodV1.SayHello()
	hellomod.SayHello("manu")
}
