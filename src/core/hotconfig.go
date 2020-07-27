package core

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

const (
	Default = 0
	Yaml	= 1
	Ini		= 2
	Json	= 3
)

type HotConfNotifyer interface {
	HCN_callback(config *HotConfig)
}

type HotConfig struct {
	ModuleFrame
	thread struct{
		GoThread
	}
	filename string
	filetype int
	data map[interface{}]interface{}
	lastModifyTime int64
	rwLock sync.RWMutex
	notifyList []HotConfNotifyer
}

func (self *HotConfig) Init_module() {
	self.load()
}

func (self *HotConfig) Run_before_module() {}

func (self *HotConfig) Run_module() {
	self.thread.Start()
}

func (self *HotConfig) Run_after_module() {}

func (self *HotConfig) Close_module() {
	self.thread.Stop()
}

func (self *HotConfig) AddObserver(n HotConfNotifyer) {
	self.notifyList = append(self.notifyList, n)
}

func (self *HotConfig) ClearObserver() {
	self.notifyList = self.notifyList[0:0]
}

func (self *HotConfig) GetData() map[interface{}]interface{}{
	self.rwLock.RLock()
	defer self.rwLock.RUnlock()
	return self.data
}

func (self *HotConfig) GetInt(key string) (value int) {
	self.rwLock.RLock()
	defer self.rwLock.RUnlock()

	return self.data[key].(int)
}

func (self *HotConfig) GetBool(key string) (value bool) {
	self.rwLock.RLock()
	defer self.rwLock.RUnlock()

	return self.data[key].(bool)
}

func (self *HotConfig) GetString(key string) (value string) {
	self.rwLock.RLock()
	defer self.rwLock.RUnlock()

	return self.data[key].(string)
}

func (self *HotConfig) GetMap(key string) (value map[interface{}]interface{}) {
	self.rwLock.RLock()
	defer self.rwLock.RUnlock()

	return self.data[key].(map[interface{}]interface{})
}

func (self *HotConfig) GetList(key string) (value []interface{}) {
	self.rwLock.RLock()
	defer self.rwLock.RUnlock()

	return self.data[key].([]interface{})
}

func (self *HotConfig) load() {
	self.Name = "HotConfig"
	self.filename = "./conf/dev.yaml"
	self.filetype = Yaml

	self.data = make(map[interface{}]interface{})
	buffer, err := ioutil.ReadFile(self.filename)
	if err != nil {
		//todo
		return
	}

	self.rwLock.Lock()
	err = yaml.Unmarshal(buffer, &self.data)
	self.lastModifyTime, err = self.modTime()
	self.rwLock.Unlock()

	self.thread.Routine = self.reload
	go self.thread.Create()
	return

}

func (self *HotConfig) reload() {
	// Timer
	time.Sleep(time.Second * 5)
	currentModifyTime, err := self.modTime()
	if err != nil {
		//todo
		return
	}

	//GetLogger().Info(currentModifyTime)
	if currentModifyTime > self.lastModifyTime {
		buffer, err := ioutil.ReadFile(self.filename)
		if err != nil {
			//todo
			return
		}

		self.rwLock.Lock()
		err = yaml.Unmarshal(buffer, &self.data)
		self.lastModifyTime = currentModifyTime
		self.rwLock.Unlock()

		//GetLogger().Info(self.data)

		// append notifyList
		for _, n := range self.notifyList {
			n.HCN_callback(self)
		}
	}
}

func (self *HotConfig) modTime() (int64, error) {
	f, err := os.Open(self.filename)
	if err != nil {

	}
	defer f.Close()

	fi, err := f.Stat()

	return fi.ModTime().Unix(), err
}



