package HSB

const urlApiBaseInfo = "http://webapi.cninfo.com.cn/api/stock/p_stock2101"

//BaseInfo 股票基本信息 "http://webapi.cninfo.com.cn/api/stock/p_stock2101"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
type BaseInfo struct {
	ORGNAME string //机构名称
	SECCODE string //证券代码
	SECNAME string //证券简称

	F001V string  //拼音简称
	F002V string  //证券类别编码
	F003V string  //证券类别
	F004V string  //交易市场编码
	F005V string  //交易市场
	F006D string  //上市日期
	F007N float64 //初始上市数量
	F008V string  //代码属性编码
	F009V string  //代码属性
	F010V string  //上市状态编码
	F011V string  //上市状态
	F012N float64 //面值
	F013V string  //ISIN
}

const urlApiStockFinanceDetail = "http://webapi.cninfo.com.cn/api/stock/p_rzrq3104"

//StockFinanceDetail 融资融券明细数据 "http://webapi.cninfo.com.cn/api/stock/p_rzrq3104"
//params:	scode	股票代码	string	输入不超过300只股票代码，用逗号分隔；如： 000001,600000
//			sdate	开始变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
type StockFinanceDetail struct {
	TRADEDATE string //交易日期
	SECCODE   string //标的证券代码
	SECNAME   string //标的证券简称

	F001N float64 //本日融资余额 单位：元
	F002N float64 //本日融资买入额 单位：元
	F003N float64 //本日融资偿还额 单位：元
	F004N float64 //本日融券余量 单位：股
	F006N float64 //本日融券卖出量 单位：股
	F007N float64 //本日融券偿还量 单位：股
	F008N float64 //融券余量金额 单位：元
	F009N float64 //融资融券余额	单位：元
	F010V string  //数据来源
	F011V string  //证券类别编码
	F012V string  //证券类别
	MEMO  string  //备注
}

const apiDetailNew = "http://webapi.cninfo.com.cn/api/stock/p_stock2401"
const apiDetailHistory = "http://webapi.cninfo.com.cn/api/stock/p_stock2402"

//Detail
//股票最新日行情 "http://webapi.cninfo.com.cn/api/stock/p_stock2401"
//params:	scode	股票代码	string	否	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
//股票历史日行情 "http://webapi.cninfo.com.cn/api/stock/p_stock2402"
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type Detail struct {
	SECCODE   string //证券代码
	SECNAME   string //证券简称
	TRADEDATE string //交易日期

	F001V string  //交易所
	F002N float64 //昨收盘 单位：元
	F003N float64 //开盘盘 单位：元
	F004N float64 //成交数量 单位：股
	F005N float64 //最高价 单位：元
	F006N float64 //最低价 单位：元
	F007N float64 //最近成交价 单位：元
	F008N float64 //总笔数 单位：笔
	F009N float64 //涨跌 单位：元
	F010N float64 //涨跌幅 单位：%
	F011N float64 //成交金额 单位：元
	F012N float64 //换手率
	F013N float64 //振幅 单位：%
	F020N float64 //发行总股本 单位：股
	F021N float64 //流通股本 单位：股
	F026N float64 //市盈率
}

const apiBlockTrade = "http://webapi.cninfo.com.cn/api/stock/p_stock2416"

//BlockTrade 大宗交易数据 "http://webapi.cninfo.com.cn/api/stock/p_stock2416"
//params: 	scode	股票代码	string	股票代码不为空，交易日期为空时，则查询股票的所有交易数据
//			tdate	交易日期	string	交易日期与股票代码不能同时为空，输入一个交易日期，选择该日期所有数据
type BlockTrade struct {
	SECCODE   string //证券代码
	SECNAME   string //证券简称
	TRADEDATE string //交易日期

	F001V string  //交易所
	F007N float64 //序号
	F002V string  //买方营业部
	F003V string  //卖方营业部
	F004N float64 //成交价格
	F005N float64 //成交量
	F006N float64 //成交金额
}
