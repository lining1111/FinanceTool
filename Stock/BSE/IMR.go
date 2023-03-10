package BSE

const apiIMR = "http://webapi.cninfo.com.cn/api/stock/p_stock2501_BSE"

//IMR 行业排名 "http://webapi.cninfo.com.cn/api/stock/p_stock2501_BSE"
type IMR struct {
	SECNAME string `excel:"证券简称"` //证券简称
	SECCODE string `excel:"股票代码"` //股票代码
	INDNAME string `excel:"行业名称"` //行业名称
	INDID   string `excel:"行业ID"` //行业ID

	F001V string  `excel:"行业级别"`     //行业级别
	F002V string  `excel:"级别说明"`     //级别说明
	F003D string  `excel:"报告期"`      //报告期
	F004N float64 `excel:"每股收益"`     //每股收益
	F005N float64 `excel:"行业均值"`     //行业均值
	F006N uint64  `excel:"行业排名"`     //行业排名
	F007N float64 `excel:"扣除后每股收益"`  //扣除后每股收益
	F008N float64 `excel:"行业均值"`     //行业均值
	F009N uint64  `excel:"行业排名"`     //行业排名
	F010N float64 `excel:"每股净资产"`    //每股净资产
	F011N float64 `excel:"行业均值"`     //行业均值
	F012N uint64  `excel:"行业排名"`     //行业排名
	F013N float64 `excel:"净资产收益率"`   //净资产收益率
	F014N float64 `excel:"行业均值"`     //行业均值
	F015N uint64  `excel:"行业排名"`     //行业排名
	F016N float64 `excel:"每股未分配利润"`  //每股未分配利润
	F017N float64 `excel:"行业均值"`     //行业均值
	F018N uint64  `excel:"行业排名"`     //行业排名
	F019N float64 `excel:"每股经营现金流量"` //每股经营现金流量
	F020N float64 `excel:"行业均值"`     //行业均值
	F021N uint64  `excel:"行业排名"`     //行业排名
	F022N float64 `excel:"营业收入增长率"`  //营业收入增长率
	F023N float64 `excel:"行业均值"`     //行业均值
	F024N uint64  `excel:"行业排名"`     //行业排名
	F025N float64 `excel:"净利润增长率"`   //净利润增长率
	F026N float64 `excel:"行业均值"`     //行业均值
	F027N uint64  `excel:"行业排名"`     //行业排名
	F028N float64 `excel:"毛利率"`      //毛利率
	F029N float64 `excel:"行业均值"`     //行业均值
	F030N uint64  `excel:"行业排名"`     //行业排名
	F031N float64 `excel:"资产负债率"`    //资产负债率
	F032N float64 `excel:"行业均值"`     //行业均值
	F033N uint64  `excel:"行业排名"`     //行业排名
	F034N float64 `excel:"应收帐款周转率"`  //应收帐款周转率
	F035N float64 `excel:"行业均值"`     //行业均值
	F036N uint64  `excel:"行业排名"`     //行业排名
}
