package core

import (
	"reflect"
)

type Cycle struct {
	enableSignal bool

	structs []interface{}

	moduleString []string

	moduleStruct map[string]reflect.Type

	module map[string]*ModuleFrame

	logger Logger

	recorder Logger
}

var cycleInstance *Cycle

func CycleInstance(ms []string, st []interface{}, mst map[string]reflect.Type, m map[string]*ModuleFrame) *Cycle {
	if cycleInstance == nil {
		cycleInstance = &Cycle{
			enableSignal:true,
			moduleString:ms,
			structs:st,
			moduleStruct:mst,
			module:m,
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

func (self *Cycle) Cycle() {
	self.init()

	self.run()

	self.signal()
}

func (self *Cycle)signal() {
	if self.enableSignal {
		s:=&SignalT{QuitHandle:self.quit}
		s.Create()
	}
}

func (self *Cycle) bindLog() {
	self.logger = NewLoggerAdapter("./logs/error.log", DEBUG)
	//self.logger = NewLoggerAdapter("./logs/error.log", "debug+")

	self.logger.Info("Log bind over")
}

func (self *Cycle) bindRecordLog() {
	self.recorder = NewLoggerAdapter("./logs/record.log", DEBUG)

	self.logger.Info("Record log bind over")
}

func (self *Cycle) init() {

	self.bindLog()

	self.bindRecordLog()

	/*self.structs = [] interface{} {
		Startup{},
		Bootstrap{},
		Conf{},
		consume.MultiPortrait{},
	}*/

	self.logger.Info("Cycle init start")

	self.moduleStruct = make(map[string]reflect.Type)

	for _, v := range self.structs {
		moduleStr := reflect.TypeOf(v).String()
		if ArrayIndex(moduleStr, self.moduleString) > -1 {
			self.moduleStruct[moduleStr] = reflect.TypeOf(v)
		}
	}

	self.module = make(map[string]*ModuleFrame)
	for _, v := range self.moduleString {
		obj := reflect.New(self.moduleStruct[v]).Interface()
		self.module[v] = &ModuleFrame{
			Object:   obj,
		}
		self.module[v].Init_module()
	}
}

func (self *Cycle) run() {
	self.logger.Info("Cycle run start")
	for _, v:= range self.moduleString {
		self.module[v].Run_before_module()
		self.module[v].Run_module()
		self.module[v].Run_after_module()
	}
}

func (self *Cycle) quit() {
	self.logger.Info("Cycle quit start")

	// 这里需要清空HotConfig的notifyer，否则会发生死锁
	GetHotConfig().ClearObserver()

	_arr := ArrayStringRevers(self.moduleString)

	for _, v := range _arr {
		self.module[v].Close_module()
	}

	self.logger.Close()
}

