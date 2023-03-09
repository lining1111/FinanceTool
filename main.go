package main

import (
	"FinanceTool/COM/cninfo"
	"FinanceTool/Stock/HSB"
	"flag"
	"github.com/golang/glog"
)

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
	baseInfo := make([]HSB.BaseInfo, 1, 2000)
	params := map[string]string{
		"scode": "000001",
		"sdate": cninfo.Getrdate("2021", cninfo.Q1),
		"edate": cninfo.Getrdate("2021", cninfo.Q1),
		"type":  "071001",
	}

	err := cninfo.GetInfoByScodeDate(HSB.APIBaseInfo, params, &baseInfo)
	if err != nil {
		glog.Info(err)
	}

}
