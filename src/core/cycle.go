package core

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

const (
	DeveEnv string = "deve"
	ProdEnv string = "prod"
	TestEnv string = "test"
)

type Cycle struct {
	env string

	injector *injector

	moduleArr []interface{}

	enableSignal bool

	structs []interface{}

	moduleString []string

	moduleStruct map[string]reflect.Type

	module map[string]*ModuleFrame

	fileconf string

	filelog string

	filepid string

	logger Logger

	recorder Logger

	cycleLock sync.Mutex
}

var cycleInstance *Cycle

func CycleInstance(st []interface{}) *Cycle {
	if cycleInstance == nil {
		cycleInstance = &Cycle{
			env:DeveEnv,
			injector:NewInjector(),
			enableSignal:true,
			structs:st,
		}
	}
	return cycleInstance
}

func GetCycleInstance() *Cycle {
	return cycleInstance
}

func GetModules() map[string]*ModuleFrame {
	return cycleInstance.module
}

func GetModule(s string) interface{} {
	return cycleInstance.module[s].Object
}

func GetConfModule() ConfImpl {
	return cycleInstance.module["core.Conf"].Object.(ConfImpl)
}

func GetHotConfig() *HotConfig{
	return cycleInstance.module["core.HotConfig"].Object.(*HotConfig)
}

func GetLogger() Logger{
	return cycleInstance.logger
}

func GetRecorder() Logger{
	return cycleInstance.recorder
}

func GetInjectModule(v interface{}) interface{} {
	return cycleInstance.injector.Get(reflect.TypeOf(v)).Elem().Interface().(ModuleFrame).Object
}

func (self *Cycle) Cycle() *Cycle {
	self.init()

	self.run()

	self.signal()

	return self
}

func (self *Cycle) SetEnv(env string) *Cycle {
	self.env = env
	return self
}

func (self *Cycle)signal() {
	if self.enableSignal {
		s:=&SignalT{
			QuitHandle:self.quit,
			SIGUSR2Handle:self.restart,
		}
		s.Create()
	}
}

func (self *Cycle) bindLog() {
	self.logger = NewLoggerAdapter(self.filelog, DEBUG)
	//self.logger = NewLoggerAdapter("./logs/error.log", "debug+")

	self.logger.Info("Log bind over")
}

func (self *Cycle) bindRecordLog() {
	self.recorder = NewLoggerAdapter("./logs/record.log", DEBUG)

	self.logger.Info("Record log bind over")
}

func (self *Cycle) init() {

	get_options(self)

	self.bindLog()

	self.bindRecordLog()

	self.logger.Info("Cycle init start")

	for _, v :=range self.structs {
		//fmt.Println(self.injector.Get(reflect.TypeOf(v)).Elem())
		m := &ModuleFrame{
			Object:   reflect.New(reflect.TypeOf(v)).Interface(),
		}
		m.Init_module(self)
		self.injector.Set(reflect.TypeOf(v), reflect.ValueOf(m))

		self.moduleArr = append(self.moduleArr, v)
	}

}

func (self *Cycle) run() {
	self.logger.Info("Cycle run start")

	for _, v:=range self.structs {
		if self.injector.Get(reflect.TypeOf(v)).Type() == reflect.TypeOf(&ModuleFrame{}) {

			m := self.injector.Get(reflect.TypeOf(v)).Elem().Interface().(ModuleFrame)
			m.Run_before_module()
			m.Run_module()
			m.Run_after_module()
		}
	}

}

func (self *Cycle) quit() {
	self.cycleLock.Lock()
	defer self.cycleLock.Unlock()

	go func (delay time.Duration) {
		for {
			for _,r :=range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(delay)
			}
		}
	}(100 * time.Millisecond)

	self.logger.Info("Cycle quit start")

	// 这里需要清空HotConfig的notifyer，否则会发生死锁
	GetInjectModule(HotConfig{}).(*HotConfig).ClearObserver()

	_arr := ArrayRevers(self.moduleArr)

	for _, v := range _arr {
		m := self.injector.Get(reflect.TypeOf(v)).Elem().Interface().(ModuleFrame)
		m.Close_module()
	}

	self.moduleArr = append([]interface{}{})

	for _, v := range self.structs {
		delete(self.injector.values, reflect.TypeOf(v))
	}

	self.logger.Close()
}

func (self *Cycle) restart() {
	self.quit()

	self.Cycle()
}

