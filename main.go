package main

import (
	"FinanceTool/FinancialStatement"
	"FinanceTool/FinancialStatement/BalanceSheet"
	"FinanceTool/FinancialStatement/CashFlowSheet"
	"FinanceTool/FinancialStatement/IncomeSheet"
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
	BalanceSheet.GetFromCNINFByScode_test()
	IncomeSheet.GetFromCNINFByScode_test()
	CashFlowSheet.GetFromCNINFByScode_test()
	FinancialStatement.GetFromCNINFByScode_test()
}
