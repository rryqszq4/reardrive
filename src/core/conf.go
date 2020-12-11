package core

import (
	"reardrive/conf_struct"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfImpl interface {
	GetAll() conf_struct.CoreConf
}

type Conf struct {
	ModuleFrame
	FileName string
	Opt conf_struct.CoreConf
}

func (self *Conf) Init_module(cycle *Cycle) {
	self.Name = "Conf"
	self.FileName = cycle.fileconf
	self.Opt = conf_struct.CoreConf{}
	buffer , err := ioutil.ReadFile(self.FileName)
	err = yaml.Unmarshal(buffer, &self.Opt)
	if err != nil {
		GetLogger().Error(err)
	}
	GetLogger().Infof("%v", self.Opt)
}

func (self *Conf) Run_before_module() {}

func (self *Conf) Run_module() {
	GetLogger().Infof("conf run")
}

func (self *Conf) Run_after_module() {}

func (self *Conf) Close_module() {
	GetLogger().Infof("conf close")
}

func (self *Conf) GetAll() conf_struct.CoreConf {
	return self.Opt
}
