package main

import (
	"reardrive/src/core"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main()  {

	cb := core.NewBipBuffer(16)
	fmt.Println(cb.Size())

	fmt.Println(cb.IsEmpty())

	timer1 := time.NewTicker(time.Second * 1)
	timer2 := time.NewTicker(time.Second * 4)
	var lock sync.RWMutex

	log := core.NewLogFile("./logs/error.log")

	go func() {
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
		}
	}()

	go func() {
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
	}()


	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}