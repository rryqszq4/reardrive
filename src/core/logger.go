package core

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type logentry_t struct {


}

type logger_t struct {

	logfile 	*logfile_t

	level 		int

	buf 		[]byte
}

type logfile_t struct {
	filename 		string

	flag 			int

	mode 			os.FileMode

	buf 			*bipbuf_t

	mutex 			sync.Mutex

	rwm 			sync.RWMutex

	fsync 			chan bool

	fwait 			sync.WaitGroup

	thread_watcher 	func()
}

const (
	TRACE = 1<<0
	DEBUG = 1<<1
	INFO = 1<<2
	WARN = 1<<3
	ERROR = 1<<4
)

var (
	levelStrings = map[int][]byte{
		TRACE:	[]byte("TRACE"),
		DEBUG: 	[]byte("DEBUG"),
		INFO:	[]byte("INFO"),
		WARN:	[]byte("WARN"),
		ERROR:	[]byte("ERROR"),
	}
)

const LOG_BUFFER_SIZE = 1024 * 32
const LOG_DELIMITER = ' '
const LOG_END = '\n'

func NewLogger(filename string, level int) *logger_t {
	self := &logger_t{
		logfile : NewLogFile(filename),
		level: level,
	}

	self.buf = self.buf[:0]
	return self
}

func (self *logger_t) Close() {
	self.logfile.Close()
}

func (self *logger_t) record(levelStr []byte, s string) {
	var buf []byte
	buf = time.Now().AppendFormat(buf,"2006-01-02 15:04:05.000000")
	buf = append(buf, LOG_DELIMITER)
	var file string
	var line int
	var ok bool
	_, file, _, ok = runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	buf = append(append(buf, levelStr...), LOG_DELIMITER)
	buf = append(append(buf, file...), ':')
	buf = strconv.AppendInt(buf, int64(line), 10)
	buf = append(buf, LOG_DELIMITER)
	buf = append(append(buf, s...), LOG_END)

	self.logfile.Write(buf)
}

func (self *logger_t) Trace(v ...interface{}) {
	if self.level & TRACE >0 {
		self.record(levelStrings[TRACE],fmt.Sprint(v...))
	}
}

func (self *logger_t) Debug(v ...interface{}) {
	if self.level & DEBUG >0 {
		self.record(levelStrings[DEBUG],fmt.Sprint(v...))
	}
}

func (self *logger_t) Info(v ...interface{}) {
	if self.level & INFO >0 {
		self.record(levelStrings[INFO],fmt.Sprint(v...))
	}
}

func (self *logger_t) Warn(v ...interface{}) {
	if self.level & WARN >0 {
		self.record(levelStrings[WARN],fmt.Sprint(v...))
	}
}

func (self *logger_t) Error(v ...interface{}) {
	if self.level & ERROR >0 {
		self.record(levelStrings[ERROR],fmt.Sprint(v...))
	}
}

func NewLogFile(filename string) *logfile_t{
	self := &logfile_t{
		filename : filename,
		flag : os.O_CREATE | os.O_WRONLY | os.O_APPEND,
		mode : 0666,
		buf : NewBipBuffer(LOG_BUFFER_SIZE),
		fsync: make(chan bool),
	}

	self.fwait.Add(1)

	timer := time.NewTicker(time.Second * 3)

	self.thread_watcher = func() {
		for {
			select {
			case <-self.fsync:
				//self.rwm.Lock()
				//self.flush()
				//self.rwm.Unlock()
				self.fwait.Done()
				//fmt.Println("==========sync==========")
			case <-timer.C:
				self.rwm.Lock()
				self.flushOnce()
				self.rwm.Unlock()
				//fmt.Println("==========async=========")
			}
		}
	}

	go self.thread_watcher()

	return self
}

func (self *logfile_t) Close() {
	self.flush()
}

func (self *logfile_t) Write(message []byte) error {
	//message = time.Now().Format("2006-01-02 15:04:05.0000000")+" [INFO] "+message+"\n"

	//self.mutex.Lock()

	//fmt.Println(self.buf.Unused())

	self.rwm.Lock()
	unused := self.buf.Unused()

	if unused < uint32(len(message)) {

		self.buf.Offer(message[0:unused])

		var tmp uint32
		for  {
			//self.waitSyncOnce()

			self.flushHalf()

			tmp = self.buf.Offer(message[unused:])

			if tmp > 0 {
				break
			}
		}

	}else {

		if self.buf.Offer(message) <= 0 {
			fmt.Println("write failed")
		}

	}
	//self.buf.Print()

	self.rwm.Unlock()

	//self.mutex.Unlock()

	return nil
}

func (self *logfile_t) openFile() (*os.File, error) {

	file, err := os.OpenFile(self.filename, self.flag, self.mode)
	if err != nil {
		return file, err
	}

	return file, nil


}

func (self *logfile_t) waitSyncOnce() {
	self.fsync <- true
	self.fwait.Wait()
	self.fwait.Add(1)
}

func (self *logfile_t) flushHalf() error {

	file, err := self.openFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//fmt.Printf("Unused=>%d, %d\n",self.buf.Unused(), self.buf.Used())
	if self.buf.Used() == LOG_BUFFER_SIZE  {
		_, err = file.Write(self.buf.Poll(LOG_BUFFER_SIZE/2))
	}

	if err != nil {
		panic(err)
	}

	return nil
}

func (self *logfile_t) flushOnce() error {

	file, err := self.openFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//fmt.Printf("Unused=>%d, %d\n",self.buf.Unused(), self.buf.Used())

	polled := self.buf.Poll(self.buf.Used())
	if polled == nil {
		_, err = file.Write(self.buf.Poll(LOG_BUFFER_SIZE/2))
	}else {
		_, err = file.Write(polled)
	}

	if err != nil {
		panic(err)
	}

	return nil
}

func (self *logfile_t) flush() error {

	file, err := self.openFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//fmt.Printf("Unused=>%d, %d\n",self.buf.Unused(), self.buf.Used())

	//self.buf.Print()
	for   {
		polled := self.buf.Poll(self.buf.Used())
		if polled == nil {
			_, err = file.Write(self.buf.Poll(LOG_BUFFER_SIZE/2))
		}else {
			_, err = file.Write(polled)
			break
		}
	}
	//self.buf.Print()

	if err != nil {
		panic(err)
	}

	return nil
}