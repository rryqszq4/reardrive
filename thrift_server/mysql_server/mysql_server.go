package mysql_server

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reardrive/core"
	"reardrive/core/modules"
	"reardrive/thrift_server/library/mysql_service"
	"github.com/apache/thrift/lib/go/thrift"
	"time"
)

type MysqlServer struct {
	modules.ThriftServerModule

	mysqlPool *sql.DB

	mysqlPoolMap map[string]*sql.DB

}

func (self *MysqlServer) Run_before_module() {
	self.initMysqlPool()

	self.initThriftServer()
}


func (self *MysqlServer) initMysqlPool() {
	var err error
	_conf := core.GetInjectModule(core.HotConfig{}).(*core.HotConfig).GetMap("mysql")
	self.mysqlPool, err = sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		_conf["user"],
		_conf["password"],
		_conf["host"],
		_conf["port"],
		_conf["database"],
		_conf["charset"],
	))
	if (err != nil) {
		self.Logger.Error(err)
	}

	self.mysqlPool.SetConnMaxLifetime(30 * time.Second)
	self.mysqlPool.SetMaxOpenConns(2)
	self.mysqlPool.SetMaxIdleConns(2)

	_conf_mysql_service := core.GetInjectModule(core.HotConfig{}).(*core.HotConfig).GetMap("mysql_service")
	self.mysqlPoolMap = make(map[string]*sql.DB, len(_conf_mysql_service))
	for _,v := range _conf_mysql_service {
		self.mysqlPoolMap[fmt.Sprintf("%s@%s:%d/%s",
			v.(map[interface{}]interface{})["user"],
			v.(map[interface{}]interface{})["host"],
			v.(map[interface{}]interface{})["port"],
			v.(map[interface{}]interface{})["database"],
		)], err = sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
			v.(map[interface{}]interface{})["user"],
			v.(map[interface{}]interface{})["password"],
			v.(map[interface{}]interface{})["host"],
			v.(map[interface{}]interface{})["port"],
			v.(map[interface{}]interface{})["database"],
			v.(map[interface{}]interface{})["charset"],
		))
		if (err != nil) {
			self.Logger.Error(err)
		}
	}
	for _, v := range self.mysqlPoolMap {
		v.SetConnMaxLifetime(30 * time.Second)
		v.SetMaxOpenConns(2)
		v.SetMaxIdleConns(2)
	}

}

func (self *MysqlServer) Run_after_module() {
}


func (self *MysqlServer) initThriftServer() {

	_thrift := core.GetInjectModule(core.HotConfig{}).(*core.HotConfig).GetMap("thrift_server")

	_addr := _thrift["addr"]
	self.Addr = _addr.(string)

	_protocol := _thrift["protocol"]
	self.Protocol = _protocol.(string)

	self.Hooker.GetProcessor = self.getProcessor
}

func (self *MysqlServer) getProcessor() thrift.TProcessor {
	handler := NewMysqlHandler(self.mysqlPoolMap)

	return mysql_service.NewMysqlServiceProcessor(handler)
}

func (self *MysqlServer) Close_module() {

	for _, v := range self.mysqlPoolMap {
		v.Close()
	}

	self.mysqlPool.Close()

	self.Server.Stop()
	self.Stop()
}