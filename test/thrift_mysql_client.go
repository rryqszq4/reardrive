package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"reardrive/thrift_server/library/mysql_service"
)

var defaultMysqlCtx = context.Background()

func handleMysqlClient(client *mysql_service.MysqlServiceClient) (err error) {
	sql := "SELECT * FROM table limit 10;"
	result, err := client.Query(defaultMysqlCtx,&mysql_service.MysqlStruct{
		"127.0.0.1",
		3306,
		"user",
		"password",
		"databases",
		"utf8"},
		&mysql_service.QueryStruct{sql})
	if err != nil {
		fmt.Println(err)
	}else {
		if result != nil {
			for _, item := range result.Rows {
				fmt.Println(item)
			}
			fmt.Println(result.Query)
			fmt.Println(result.Columns)
			fmt.Println(result.Error)
		}else {
			fmt.Println("error: result is nil")
		}
	}

	sql = "select version();"
	result1, err := client.Execute(defaultMysqlCtx,&mysql_service.MysqlStruct{
		"127.0.0.1",
		3306,
		"user",
		"password",
		"databases",
		"utf8"},
		&mysql_service.QueryStruct{sql})
	if err != nil {
		fmt.Println(err)
	}else {
		if result1 != nil {
			fmt.Println(result.Query)
			fmt.Println(result1.GetLastInsertId())
			fmt.Println(result1.GetRowsAffected())
			fmt.Println(result1.GetError())
		}else {
			fmt.Println("error: result1 is nil")
		}
	}

	return err
}

func runMysqlClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return handleMysqlClient(mysql_service.NewMysqlServiceClient(thrift.NewTStandardClient(iprot, oprot)))
}

func main() {

	var transportFactory thrift.TTransportFactory
	var protocolFactory thrift.TProtocolFactory

	transportFactory = thrift.NewTTransportFactory()
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()


	if err := runMysqlClient(transportFactory, protocolFactory, "localhost:9090", false); err != nil {
		fmt.Println("error running client: ", err)
	}

}