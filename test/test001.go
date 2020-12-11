package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	defer func() {
		fmt.Println("time used:", float64(time.Now().UnixNano()-start)/1e9)
	}()

	logger := NewLogger("./logs/error.log")
	defer logger.Close()

	var wg sync.WaitGroup
	defer wg.Wait()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				record := []byte("Test logging, but use a somewhat realistic message length.Test logging, but use a somewhat realistic message length.")
				record = append(record, []byte(strconv.Itoa(i*100+j))...)
				record = append(record, '\n')
				logger.Write(record)
			}
		}(i)
	}
}

type Logger struct {
	fp           *os.File
	recordBuffer chan *Buffer
	flushBuffer  chan *Buffer
	bufferCount  int
	bufferSize   int
	current      *Buffer
	wg           sync.WaitGroup
	lock         *sync.Mutex
}

func NewLogger(filename string) (l *Logger) {
	var (
		fp  *os.File
		err error
	)

	if fp, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); nil != err {
		log.Fatal(err)
	}

	l = &Logger{
		fp:          fp,
		bufferCount: 1024,
		bufferSize:  1024 * 1024,
		lock:        new(sync.Mutex),
	}

	l.recordBuffer = make(chan *Buffer, l.bufferCount)
	l.flushBuffer = make(chan *Buffer, l.bufferCount)

	for i := 0; i < l.bufferCount; i++ {
		l.recordBuffer <- NewBuffer(l.bufferSize)
	}

	l.current = <-l.recordBuffer

	l.wg.Add(1)
	go l.Flusher()

	return
}

func (l *Logger) Write(record []byte) {
	l.lock.Lock()
	defer l.lock.Unlock()

	rc := append([]byte(l.Datetime()), record...)
	for {
		ll := len(rc)
		valid := l.current.Valid()
		if ll > valid { // 放不下了
			_, _ = l.current.Write(rc[0:valid])
			rc = rc[valid:]
			l.flushBuffer <- l.current
			l.current = <-l.recordBuffer
		} else {
			_, _ = l.current.Write(rc)
			return
		}
	}
}

func (l *Logger) Datetime() string {
	return time.Now().Format("2006/01/02 15:04:05.999999 ")
}

func (l *Logger) Close() {
	defer l.fp.Close()

	close(l.recordBuffer)
	close(l.flushBuffer)

	l.wg.Wait()

	l.current.WriteTo(l.fp)
}

func (l *Logger) Flusher() {
	defer l.wg.Done()

	for b := range l.flushBuffer {
		_, _ = b.WriteTo(l.fp)
	}
}

type Buffer struct {
	*bytes.Buffer
}

func NewBuffer(size int) (b *Buffer) {
	b = &Buffer{
		Buffer: bytes.NewBuffer(make([]byte, 0, size)),
	}

	return
}

func (b *Buffer) Valid() int {
	return b.Cap() - b.Len()
}
