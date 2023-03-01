package FinancialStatement

import (
	"FinanceTool/FinancialStatement/BalanceSheet"
	"FinanceTool/FinancialStatement/CashFlowSheet"
	"FinanceTool/FinancialStatement/IncomeSheet"
	"errors"
	"github.com/golang/glog"
	"time"
)

//FinanceStatement 财务报表
type FinanceStatement struct {
	//可能需要通过其他的网络手段进行查询的
	RegisteredCapital float64 //企业的注册资本
	//实际表内元素
	OrganizationName        string    //编报企业的名称
	CTime                   time.Time //资产负债表日或财务报表涵盖的会计期间
	Unit                    float64   //人民币金额单位
	IsConsolidatedStatement bool      //财务报表是合并财务报表的，应当予以标明

	BSU BalanceSheet.BSUnit
	ISU IncomeSheet.ISUnit
	CSU CashFlowSheet.CSUnit
}

//Check 检查表是否符合规定 上市公司的财务报表是需要对外发布的，肯定是经过审计的。但是非上市公司的财务报表是没有经过审计的，存在一定的可疑性
func (fs *FinanceStatement) Check() bool {

	//单表检查
	if !fs.BSU.Check() {
		glog.Error("FinanceStatement Check BSU fail,year:%v", fs.CTime.String())
		return false
	}

	if !fs.ISU.Check() {
		glog.Error("FinanceStatement Check ISU fail,year:%v", fs.CTime.String())
		return false
	}

	if !fs.CSU.Check() {
		glog.Error("FinanceStatement Check CSU fail,year:%v", fs.CTime.String())
		return false
	}

	//联合检查
	if !fs.CheckUnit() {
		glog.Error("FinanceStatement Check Unit fail,year:%v", fs.CTime.String())
		return false
	}
	checkSurplusReserves, _ := fs.CheckSurplusReserves(fs.RegisteredCapital)
	if !checkSurplusReserves {
		glog.Error("FinanceStatement Check SurplusReserves fail,year:%v", fs.CTime.String())
		return false
	}

	return true
}

//CheckUnit 查看人民币金额单位是否一致
func (fs *FinanceStatement) CheckUnit() bool {
	if fs.Unit != fs.BSU.Last.Unit || fs.Unit != fs.BSU.Now.Unit {
		glog.Error("财务报表金额单位与资产负债表金额单位不相同")
		return false
	} else if fs.Unit != fs.ISU.Last.Unit || fs.Unit != fs.ISU.Now.Unit {
		glog.Error("财务报表金额单位与利润表金额单位不相同")
		return false
	} else if fs.Unit != fs.CSU.Last.Unit || fs.Unit != fs.CSU.Now.Unit {
		glog.Error("财务报表金额单位与现金流量表金额单位不相同")
		return false
	} else {
		return true
	}
}

func (fs FinanceStatement) GetSurplusReservesRatio() float64 {
	return fs.BSU.GetSurplusReservesInc() / fs.ISU.Now.NetProfit
}

//资产负债表与利润表

//公司法 http://www.npc.gov.cn/npc/c12435/201811/68a85058b4c843d1a938420a77da14b4.shtml
//第一百六十六条　公司分配当年税后利润时，应当提取利润的百分之十列入公司法定公积金。公司法定公积金累计额为公司注册资本的百分之五十以上的，可以不再提取。

//CheckSurplusReserves 查看盈余公积数额正确性
func (fs *FinanceStatement) CheckSurplusReserves(registeredCapital float64) (bool, error) {
	if registeredCapital <= 0 {
		glog.Error("注册资本小于等于0")
		return false, errors.New("注册资本小于等于0")
	} else {
		if fs.ISU.Now.NetProfit*fs.Unit <= 0 {
			glog.Info("公司当年利益亏损,不用提取,year:%v netProfit:%v 单位%v人民币",
				fs.CTime.String(), fs.ISU.Now.NetProfit, fs.Unit)
			return true, errors.New("公司当年利益亏损,不用提取")
		} else {
			if fs.BSU.Now.LEQ.Eq.SurplusReserves*fs.Unit >= (registeredCapital * 0.5) {
				glog.Info("公司的盈余公积已达到注册资金的50%，没有法定盈余公积应为当年利润的10%限制,year:%v", fs.CTime.String())
				return true, errors.New("公司的盈余公积已达到注册资金的50%，没有法定盈余公积应为当年利润的10%限制")
			} else {
				glog.Info("公司的盈余公积未达到注册资金的50%")
				ratio := fs.GetSurplusReservesRatio()
				if ratio >= 0.1 {
					glog.Info("公司的盈余公积已达到净利润的10%,ratio:%v", ratio)
					return true, nil
				} else {
					glog.Error("公司的盈余公积未达到净利润的10%,ratio:%v", ratio)
					return false, errors.New("公司的盈余公积未达到净利润的10%")
				}
			}
		}
	}
}
