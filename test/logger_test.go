package main

import (
	"log"
	"os"
	"strconv"
	"testing"
	"../src/core"
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

	for i := 0; i < b.N; i++ {
		fakeMessage := []byte("Test logging, but use a somewhat realistic message length.Test logging, but use a somewhat realistic message length")

		fakeMessage = append(fakeMessage, strconv.Itoa(i)...)
		fakeMessage = append(fakeMessage, '\n')
		log.Write(fakeMessage)

	}
}

func BenchmarkLogger2(b *testing.B) {
	log := core.NewLogger("../logs/error.log", core.INFO)

	fakeMessage := "Test logging, but use a somewhat realistic message length.Test logging, but use a somewhat realistic message length."

	for i :=0; i < b.N; i++ {
		log.Info(fakeMessage + strconv.Itoa(i))
	}
}

func BenchmarkStdLogger1(b *testing.B) {
	logFile, _ := os.OpenFile("../logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	log.SetOutput(logFile)

	fakeMessage := "Test logging, but use a somewhat realistic message length.Test logging, but use a somewhat realistic message length."

	for i:=0; i < b.N; i++ {
		log.Println(fakeMessage + strconv.Itoa(i))
	}
}