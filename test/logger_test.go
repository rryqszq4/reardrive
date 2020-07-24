package main

import (
	"../src/core"
	"strconv"
	"testing"
)

/*
func ExampleLogger1() {

	log := core.NewLogFile("../logs/error.log")

	//fmt.Println(123)

	for i := 0; i < 10000000; i++ {
		log.Write("Test logging, but use a somewhat realistic message length.Test logging, but use a somewhat realistic message length." + strconv.Itoa(i))
		//fmt.Println("iiiiiiiiiii=>"+strconv.Itoa(i))
	}


	// output:
	// test

	time.Sleep(time.Second*5)
}
*/

func BenchmarkLogger1(b *testing.B) {
	log := core.NewLogFile("../logs/error.log")

	fakeMessage := "Test logging, but use a somewhat realistic message length.Test logging, but use a somewhat realistic message length."

	for i := 0; i < b.N; i++ {
		log.Write(fakeMessage + strconv.Itoa(i))
	}
}