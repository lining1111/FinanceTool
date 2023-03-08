package BSE

const apiIMR = "http://webapi.cninfo.com.cn/api/stock/p_stock2501_BSE"

//IMR 行业排名 "http://webapi.cninfo.com.cn/api/stock/p_stock2501_BSE"
type IMR struct {
	SECNAME string //证券简称
	SECCODE string //股票代码
	INDNAME string //行业名称
	INDID   string //行业ID
	F001V   string //行业级别
	F002V   string //级别说明
	F003D   string //报告期

	F004N float64 //每股收益
	F005N float64 //行业均值
	F006N float64 //行业排名
	F007N float64 //扣除后每股收益
	F008N float64 //行业均值
	F009N float64 //行业排名
	F010N float64 //每股净资产
	F011N float64 //行业均值
	F012N float64 //行业排名
	F013N float64 //净资产收益率
	F014N float64 //行业均值
	F015N float64 //行业排名
	F016N float64 //每股未分配利润
	F017N float64 //行业均值
	F018N float64 //行业排名
	F019N float64 //每股经营现金流量
	F020N float64 //行业均值
	F021N float64 //行业排名
	F022N float64 //营业收入增长率
	F023N float64 //行业均值
	F024N float64 //行业排名
	F025N float64 //净利润增长率
	F026N float64 //行业均值
	F027N float64 //行业排名
	F028N float64 //毛利率
	F029N float64 //行业均值
	F030N float64 //行业排名
	F031N float64 //资产负债率
	F032N float64 //行业均值
	F033N float64 //行业排名
	F034N float64 //应收帐款周转率
	F035N float64 //行业均值
	F036N float64 //行业排名
}
