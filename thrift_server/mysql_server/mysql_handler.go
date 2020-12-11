package mysql_server

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reardrive/core"
	"reardrive/thrift_server/library/mysql_service"
)

type MysqlHandler struct {
	mysql_service.MysqlService

	mysqlPoolMap map[string]*sql.DB

	Logger core.Logger
}

func NewMysqlHandler(mysqlPoolMap map[string]*sql.DB) *MysqlHandler {
	return &MysqlHandler{
		mysqlPoolMap:mysqlPoolMap,
		Logger: core.GetLogger(),
	}
}

func (self *MysqlHandler) Query(ctx context.Context, mysql *mysql_service.MysqlStruct, query *mysql_service.QueryStruct) (r *mysql_service.QueryResultStruct, err error) {
	result := make([]*mysql_service.ItemStruct,0)

	// get mysql pool
	_mysqlPool, ok := self.mysqlPoolMap[fmt.Sprintf("%s@%s:%d/%s",
		mysql.GetUser(),
		mysql.GetHost(),
		mysql.GetPort(),
		mysql.GetDatabase(),
	)]
	if !ok {
		self.Logger.Errorf("mysql pool(%s) is not exist",fmt.Sprintf("%s@%s:%d/%s",
			mysql.GetUser(),
			mysql.GetHost(),
			mysql.GetPort(),
			mysql.GetDatabase(),
		))
		return
	}

	// query run
	self.Logger.Info("SQL: " + query.Sql)
	rows,err := _mysqlPool.Query(query.Sql)

	if err != nil {
		self.Logger.Error(err)
		return
	}

	columns,_ := rows.Columns()

	// column handle
	columnData := make(map[string]string,len(columns))
	columnTypes,_ := rows.ColumnTypes()
	for _,value := range columnTypes {
		columnData[value.Name()] = value.DatabaseTypeName()
	}

	// result handle
	rawResult := make([][]byte, len(columns))
	dest := make([]interface{}, len(columns))
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		_ = rows.Scan(dest...)
		item := &mysql_service.ItemStruct{make(map[string]string,len(columns))}

		for i, raw := range rawResult {
			if raw == nil {
				item.Data[columns[i]] = ""
			}else {
				item.Data[columns[i]] = string(raw)
			}
		}
		result = append(result, item)
	}

	var errStr = ""
	if err != nil {
		errStr = err.Error()
	}
	return &mysql_service.QueryResultStruct{
		result,
		errStr,
		query,
		columnData,
	},err
}

func (self *MysqlHandler) Execute(ctx context.Context, mysql *mysql_service.MysqlStruct, query *mysql_service.QueryStruct) (r *mysql_service.ExecuteResultStruct, err error) {
	// get mysql pool
	_mysqlPool, ok := self.mysqlPoolMap[fmt.Sprintf("%s@%s:%d/%s",
		mysql.GetUser(),
		mysql.GetHost(),
		mysql.GetPort(),
		mysql.GetDatabase(),
	)]
	if !ok {
		self.Logger.Errorf("mysql pool(%s) is not exist",fmt.Sprintf("%s@%s:%d/%s",
			mysql.GetUser(),
			mysql.GetHost(),
			mysql.GetPort(),
			mysql.GetDatabase(),
		))
		return
	}

	self.Logger.Info("SQL: " + query.Sql)
	res, err := _mysqlPool.Exec(query.Sql)

	if err != nil {
		self.Logger.Error(err)
		return
	}

	lastId ,err := res.LastInsertId()
	if err != nil {
		self.Logger.Error(err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		self.Logger.Error(err)
		return
	}

	var errStr = ""
	if err != nil {
		errStr = err.Error()
	}
	return &mysql_service.ExecuteResultStruct{
		"",
		errStr,
		lastId,
		row,
		query,
	},err
}
