package main

import (
	"reardrive/src/core"
	"github.com/kinone/sakura/mlog"
	"log"
	"os"
	"strconv"
	"testing"
)

/*func ExampleLogger1() {

	log := core.NewLogFile("../logs/error.log")

	fmt.Println(123)

	for i := 0; i < 1000000; i++ {
		a := []byte("Test logging, but use a somewhat realistic message length.Test logging, but use a somewhat realistic message length.")
		a = append(a, strconv.Itoa(i)...)
		a = append(a, '\n')
		log.Write(a)
		//fmt.Println("iiiiiiiiiii=>"+strconv.Itoa(i))
	}

	log.Close()

	// output:
	// test

	time.Sleep(time.Second*5)
}*/


func BenchmarkLogger1(b *testing.B) {
	log := core.NewLogFile("../logs/error.log")

	for i := 0; i < b.N; i++ {
		fakeMessage := []byte("Test logging, but use a somewhat realistic message length.Test logging, but use a somewhat realistic message length")

		fakeMessage = append(fakeMessage, strconv.Itoa(i)...)
		fakeMessage = append(fakeMessage, '\n')
		log.Write(fakeMessage)

	}

	defer log.Close()
}

func BenchmarkLogger2(b *testing.B) {
	log := core.NewLogger("../logs/error.log", core.INFO)

	fakeMessage := "Test logging, but use a somewhat realistic message length.Test logging, but use a somewhat realistic message length."

	for i :=0; i < b.N; i++ {
		log.Info(fakeMessage + strconv.Itoa(i))
	}

	defer log.Close()
}

func BenchmarkStdLogger1(b *testing.B) {
	logFile, _ := os.OpenFile("../logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	log.SetOutput(logFile)

	fakeMessage := "Test logging, but use a somewhat realistic message length.Test logging, but use a somewhat realistic message length."

	for i:=0; i < b.N; i++ {
		log.Println(fakeMessage + strconv.Itoa(i))
	}
}

func BenchmarkMlog1(b *testing.B) {
	var logger mlog.LevelLogger
	logger = mlog.NewLogger(&mlog.Option{
		File : "../logs/error.log",
		Levels: []string{"info+"},
	})

	fakeMessage := "Test logging, but use a somewhat realistic message length.Test logging, but use a somewhat realistic message length."

	for i:=0; i < b.N; i++ {
		logger.Info(fakeMessage + strconv.Itoa(i))
	}
}