package main

import (
	"fmt"

	"github.com/rahmaniali-ir/shorter/shorter"
)

func main()  {
	short := shorter.MakeShorter("ali.ir")
	println(short)
	long, _ := shorter.MakeLonger(short)
	println(long)
	println("__________")
	
	short2 := shorter.MakeShorter("ali.ir")
	println(short2)
	long2, _ := shorter.MakeLonger(short2)
	println(long2)
	println("__________")
	
	short3 := shorter.MakeShorter("ali.ir")
	println(short3)
	long3, _ := shorter.MakeLonger(short3)
	println(long3)
	println("__________")
	
	short4 := shorter.MakeShorter("google.com")
	println(short4)
	long4, _ := shorter.MakeLonger(short4)
	println(long4)
	println("__________")

	fmt.Printf("%v", shorter.Records)
}
