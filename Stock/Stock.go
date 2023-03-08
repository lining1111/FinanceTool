package Stock

import (
	HSB "FinanceTool/Stock/HSB"
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
