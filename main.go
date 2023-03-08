package main

import (
	AuditOpinion_stock "FinanceTool/Stock/AuditOpinion"
	BalanceSheet_stock "FinanceTool/Stock/BalanceSheet"
	CashFlowSheet_stock "FinanceTool/Stock/CashFlowSheet"
	FinancialRatios_stock "FinanceTool/Stock/FinancialRatios"
	IMR_stock "FinanceTool/Stock/IMR"
	IncomeSheet_stock "FinanceTool/Stock/IncomeSheet"
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
	BalanceSheet_stock.GetFromCNINFByScode_test()
	IncomeSheet_stock.GetFromCNINFByScode_test()
	CashFlowSheet_stock.GetFromCNINFByScode_test()
	FinancialRatios_stock.GetFromCNINFByScode_test()
	AuditOpinion_stock.GetFromCNINFByScode_test()
	IMR_stock.GetFromCNINFByScode_test()
}
