package core

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type logentry_t struct {


}

type logger_t struct {

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

const LOG_BUFFER_SIZE = 1024 * 16


func NewLogFile(filename string) *logfile_t{
	self := &logfile_t{
		filename : filename,
		flag : os.O_WRONLY | os.O_CREATE | os.O_APPEND,
		mode : 0666,
		buf : NewBipBuffer(LOG_BUFFER_SIZE),
		fsync: make(chan bool),
	}

	self.fwait.Add(1)

	timer := time.NewTicker(time.Millisecond * 100)

	self.thread_watcher = func() {
		for {
			select {
			case <-self.fsync:
				self.rwm.RLock()
				self.flush()
				self.rwm.RUnlock()
				self.fwait.Done()
				//fmt.Println("==========sync==========")
			case <-timer.C:
				self.rwm.RLock()
				self.flush()
				self.rwm.RUnlock()
				//fmt.Println("==========async=========")
			}
		}
	}

	go self.thread_watcher()

	return self
}

func (self *logfile_t) Write(message string) error {
	//message = time.Now().Format("2006-01-02 15:04:05.0000000")+" [INFO] "+message+"\n"
	message = message + "\n"

	self.mutex.Lock()

	//fmt.Println(self.buf.Unused())

	if self.buf.Unused() < uint32(len(message)) {
		self.fsync <- true
		self.fwait.Wait()
		self.fwait.Add(1)
	}

	self.rwm.Lock()
	if self.buf.Offer([]byte(message)) <= 0 {
		fmt.Println("write failed")
	}
	self.rwm.Unlock()

	self.mutex.Unlock()

	return nil
}

func NewLogger() {

}

func (self *logfile_t) openFile() (*os.File, error) {

	file, err := os.OpenFile(self.filename, self.flag, self.mode)
	if err != nil {
		return file, err
	}

	return file, nil


}

func (self *logfile_t) flush() error {

	file, err := self.openFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//fmt.Println("use=>" + strconv.Itoa(int(self.buf.Used())))
	_, err = file.WriteString(string(self.buf.Poll(self.buf.Used())))

	if err != nil {
		panic(err)
	}

	return nil
}