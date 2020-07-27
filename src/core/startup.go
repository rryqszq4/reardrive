package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
)

type Startup struct {
	Name string

	pidFile *os.File

	pidFileName string
}

func NewStartup() ModuleImpl {
	return &Startup{}
}

func (self *Startup) Init_module() {
	GetLogger().Info("startup init")
	var err error
	pid := os.Getpid()
	self.pidFileName = "./pid"

	GetLogger().Infof("Processing PID: %d", pid)

	self.pidFile, err = os.OpenFile(self.pidFileName, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		GetLogger().Error("Open pid file error.")
	}

	if syscall.Flock(int(self.pidFile.Fd()), syscall.LOCK_EX|syscall.LOCK_NB) != nil {
		GetLogger().Error("Pid file locked.")
	}

	if ioutil.WriteFile(self.pidFileName, []byte(fmt.Sprintln(pid)), os.ModePerm) != nil {
		GetLogger().Error("Write pid failed.")
	}

}

func (self *Startup) Run_before_module() {}

func (self *Startup) Run_module() {
	GetLogger().Info("startup run")
}

func (self *Startup) Run_after_module() {}

func (self *Startup) Close_module() {
	if err := self.pidFile.Close(); err != nil {
		GetLogger().Error("Close pid file failed.")
	}
	if os.Remove(self.pidFileName) != nil {
		GetLogger().Error("Remove file failed.")
	}
	GetLogger().Info("startup close")
}
