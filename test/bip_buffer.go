package main

import (
	"fmt"
	"reardrive/src/core"
	"strconv"
	"sync"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	defer func() {
		fmt.Println("time used:", float64(time.Now().UnixNano()-start)/1e9)
	}()

	//timer1 := time.NewTicker(time.Second * 1)
	//timer2 := time.NewTicker(time.Second * 4)
	//var lock sync.RWMutex

	logger := core.NewLogger("./logs/error.log",core.INFO)
	defer logger.Close()

	/*file, _ := os.OpenFile("./logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder.EncodeLevel = zapcore.CapitalLevelEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.AddSync(file),
		zapcore.DebugLevel,
	)
	logger:= zap.New(core)*/


	//f, _ := os.OpenFile("./logs/error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	//defer f.Close()
	//logger := log.New(f, "", log.LstdFlags)

	a1 := func(idx int) {
		for i := 0; i < 10000; i++ {
			a := "Test logging, but use a somewhat realistic message length.Test logging, but use a somewhat realistic message length."
			a = a + strconv.Itoa(idx * 100 + i)
			//a = append(a, '\n')
			logger.Info(a)
			//time.Sleep(time.Second)
		}

		/*time.Sleep(time.Second*5)
		for {
			select {
			case <- timer1.C:
				lock.Lock()
				ret := cb.Offer([]byte("abcd"))
				fmt.Println("Offer:")
				if ret <= 0 {
					fmt.Println("Offer failed")
					cb.Poll(cb.Used())
				}
				cb.Print()
				lock.Unlock()
			}
		}*/
	}

	var wg sync.WaitGroup
	defer wg.Wait()
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			a1(j)
		}(j)
	}

	/*go func() {
		for i := 0; i < 1000000; i++ {
			a := []byte("Test logging, but use a somewhat realistic message length.Test logging, but use a somewhat realistic message length.")
			a = append(a, strconv.Itoa(i)...)
			a = append(a, '\n')
			log.Write(a)
			//time.Sleep(time.Second)
		}

		time.Sleep(time.Second*5)
		for {
			select {
			case <- timer2.C:
				lock.Lock()
				cb.Poll(4)
				fmt.Println("Poll:")
				cb.Print()
				lock.Unlock()
			}
		}
	}()*/

	//var (
	//	name    string
	//	age     int
	//	married bool
	//)
	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
