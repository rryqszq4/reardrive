package core

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type SignalImpl interface {
	make()
	listen()
	run()
}

type SignalT struct {
	QuitHandle func()
	SIGUSR1Handle func()
	SIGUSR2Handle func()
	UnknowHandle func()
	sChan chan os.Signal
}

func (self *SignalT) Create()  {
	self.make()

	self.listen()

	self.run()

}

func (self *SignalT) make() {
	self.sChan = make(chan os.Signal)
	GetLogger().Info("Make signal")
}

func (self *SignalT) listen() {
	signal.Notify(self.sChan, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	GetLogger().Info("Notify listen signal")
}

func (self *SignalT) run() {
	for s := range self.sChan {
		GetLogger().Info("Signal run")
		switch s {
		case os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			GetLogger().Info("Signal quit ", s)
			if self.QuitHandle != nil {
				self.QuitHandle()
			}
		case syscall.SIGUSR1:
			GetLogger().Info("usr1", s)
			//self.SIGUSR1Handle()
		case syscall.SIGUSR2:
			GetLogger().Info("usr2", s)
			if self.SIGUSR2Handle != nil {
				self.SIGUSR2Handle()
			}
		default:
			GetLogger().Info("other", s)
			//self.UnknowHandle()
		}
		fmt.Println("quited")
		os.Exit(0)
	}
}

