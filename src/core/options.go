package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

var defaultExecName string

func get_options(cycle *Cycle) int {

	ret := 0
	defaultConfFile := "conf/dev.yaml"
	defaultErrorFile := "logs/error.log"
	defaultPidFile := ".pid"

	var ok bool
	if defaultExecName == "" {
		_, defaultExecName, _, ok = runtime.Caller(3)
		if !ok {
			defaultPidFile = "???"
		}
		defaultExecName = strings.Split(defaultExecName, ".go")[0][strings.LastIndexByte(strings.Split(defaultExecName, ".go")[0], '/')+1:]
	}

	usage := []string{
		"[Welcome to use "+ defaultExecName +"]\n",
		"Options:\n",
		"    -h    : Show how to use. \n",
		"    -s    : start, stop, restart server. \n",
		"    -conf : set config file (default: "+ defaultConfFile +"). \n",
		"    -log  : set error log (default: "+ defaultErrorFile +"). \n",
		"    -pid  : set pid file (default: "+ defaultExecName+defaultPidFile +"). \n",
	}

	if len(os.Args) <= 1  {
		fmt.Println(strings.Join(usage,""))
		goto end
	}

	cycle.fileconf = defaultConfFile
	cycle.filelog = defaultErrorFile
	cycle.filepid = "logs/"+defaultExecName + defaultPidFile

	for i:=1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-h":
			fmt.Println(strings.Join(usage,""))
		case "-s":
			i++
			if i >= len(os.Args) {
				fmt.Println("Missing parameter at \"-s\"")
				goto end
			}
			if os.Args[i] == "start" {
				// 不退出程序
				ret = 1
				fmt.Println(defaultExecName +" "+ os.Args[i])
			}
			if os.Args[i] == "stop" {
				// 根据pid终止进程
				buffer, err := ioutil.ReadFile(cycle.filepid)
				if err != nil {
					fmt.Errorf(err.Error())
				}
				_pid ,_ := strconv.Atoi(strings.Trim(string(buffer),"\n"))
				if _pid != 0 {
					syscall.Kill(_pid, syscall.SIGQUIT)
					fmt.Println(defaultExecName + " " + os.Args[i])
				}
			}
			if os.Args[i] == "restart" {
				// 根据pid重启进程
				buffer, err := ioutil.ReadFile(cycle.filepid)
				if err != nil {
					fmt.Errorf(err.Error())
				}
				_pid, _ := strconv.Atoi(strings.Trim(string(buffer), "\n"))
				if _pid != 0 {
					syscall.Kill(_pid, syscall.SIGUSR2)
					fmt.Println(defaultExecName + " " + os.Args[i])
				}
			}
		case "-conf":
			i++
			if cycle.fileconf != "" {
				cycle.fileconf = os.Args[i]
			}
		case "-log":
			i++
			if cycle.filelog != "" {
				cycle.filelog = os.Args[i]
			}
		case "-pid":
			i++
			if cycle.filepid != "" {
				cycle.filepid = os.Args[i]
			}
		default:
			//
		}
	}

end:

	if ret == 0 {
		os.Exit(0)
	}
	return ret
}