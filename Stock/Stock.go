package Stock

import (
	"FinanceTool/COM/cninfo"
	HSB "FinanceTool/Stock/HSB"
	"fmt"
)

/**
股票Stock
获取指定公司指定季度报表内容=股票代码+发布时间
内存级别的报表存储格式
当季数据报表的分级结构---金字塔式
	L0 总报表内容(包括 发布时间+数组个数+以L1各公司报表集合为单位的数组)-----所有公司的报表
	L1 各公司报表内容集合(包括 数组个数+以L2本公司当季度报表集合为单位的数组)-----单个公司的所有报表
	L2 本公司当季度报表集合(包括 资产负债表+利润表+现金流量表+指标表+指标行业排名+审计意见)
*/

//FSL0 总报表内容
type FSL0 struct {
	RepublicDate string //发布日期
	Count        int    //个数
	Sub          []FSL1
}

//FSL1 各公司报表内容集合
type FSL1 struct {
	RepublicDate string //发布日期
	Count        int    //个数
	Sub          []FSL2
}

//FSL2 本公司当季度报表集合
type FSL2 struct {
	RepublicDate string              //发布日期
	BS           HSB.BalanceSheet    //资产负债表
	IS           HSB.IncomeSheet     //利润表
	CS           HSB.CashFlowSheet   //现金流量表
	FR           HSB.FinancialRatios //指标表
	IMR          HSB.IMR             //指标行业排名
	AO           HSB.AuditOpinion    //审计意见
}

//股票演示代码

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

//Test1 获取A股内的所有股票，打印他们的当日最新行情
func Test1() {
	//获取所有A股股票基本信息
	stockBaseInfos := make([]HSB.BaseInfo, 0, cninfo.MaxResultLimt)
	params := map[string]string{}
	err := cninfo.GetInfoByScodeDate(HSB.APIBaseInfo, params, &stockBaseInfos, cap(stockBaseInfos))
	checkerr(err)
	if err == nil {
		fmt.Println("股票总数为", len(stockBaseInfos))
		//筛选出所有A股
		stockA := make([]HSB.BaseInfo, 0, cninfo.MaxResultLimt)
		for _, v := range stockBaseInfos {
			if v.F003V == "A股" {
				stockA = append(stockA, v)
			}
		}
		fmt.Println("A股总数为", len(stockA))
		//获取A股当日的行情信息
		detail := make([]HSB.Detail, 0, cninfo.MaxResultLimt)

		for _, v1 := range stockA {
			detail1 := make([]HSB.Detail, 0, cninfo.MaxResultLimt)

			params1 := map[string]string{
				"scode": v1.SECCODE,
			}
			err := cninfo.GetInfoByScodeDate(HSB.APIDetailNew, params1, &detail1, cap(detail1))
			checkerr(err)
			if err == nil {
				for _, v2 := range detail1 {
					detail = append(detail, v2)
				}
				//打印下信息
				for _, v := range detail {
					fmt.Printf("股票代码 %s 交易日期%s 交易所 %s 昨日收盘价%f 今日开盘价%f 成交数量%d "+
						"最高成交价%f 最低成交价%f 最近成交价%f 总笔数%d 涨跌%f 涨跌幅%f 成交金额%f "+
						"换手率 %f 振幅 %f 发行总股本%d 流通股本%d 市盈率%f",
						v.SECCODE,
						v.TRADEDATE,
						v.F001V,
						v.F002N,
						v.F003N,
						v.F004N,
						v.F005N,
						v.F006N,
						v.F007N,
						v.F008N,
						v.F009N,
						v.F010N,
						v.F011N,
						v.F012N,
						v.F013N,
						v.F020N,
						v.F021N,
						v.F026N)
				}
			}
		}

	}

}
