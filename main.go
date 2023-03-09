package main

import (
	"FinanceTool/Stock"
	"flag"
	"github.com/golang/glog"
)

func init() {
	glog.Info("start")
}

func main() {
	var port int
	flag.IntVar(&port, "port", 10000, "本地的服务端端口")
	flag.Parse()
	defer glog.Flush()
	glog.Info("欢迎来到金融与go的世界")
	glog.Info("服务端 端口:", port)
	//HSB.BSGetFromCNINFByScode_test()
	//HSB.ISGetFromCNINFByScode_test()
	//HSB.CSGetFromCNINFByScode_test()
	//HSB.IMRGetFromCNINFByScode_test()
	//HSB.AOGetFromCNINFByScode_test()
	//HSB.FRGetFromCNINFByScode_test()
	//cninfo.PublicCodeData_test()
	Stock.Test1()
}
