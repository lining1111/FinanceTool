package main

import (
	"flag"
	"github.com/golang/glog"
)

func main() {
	var port int
	var isPrintStdout bool;
	flag.IntVar(&port,"port",10000,"本地的服务端端口")
	flag.BoolVar(&isPrintStdout,"isPrintStdout",true,"日志是否打印到控制台")
	flag.Parse()
	defer glog.Flush()
	glog.Info("欢迎来到金融与go的世界")
	glog.Info("服务端 端口:",port)

}
