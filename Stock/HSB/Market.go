package HSB

import "FinanceTool/COM/cninfo"

const APIBaseInfo = "http://webapi.cninfo.com.cn/api/stock/p_stock2101"

//BaseInfo 股票基本信息 "http://webapi.cninfo.com.cn/api/stock/p_stock2101"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
type BaseInfo struct {
	ORGNAME string //机构名称	varchar
	SECCODE string //证券代码	varchar
	SECNAME string //证券简称	varchar

	F001V string  //拼音简称	varchar
	F002V string  //证券类别编码	varchar
	F003V string  //证券类别	varchar
	F004V string  //交易市场编码	varchar
	F005V string  //交易市场	varchar
	F006D string  //上市日期	datetime
	F007N uint64  //初始上市数量	decimal		单位：股
	F008V string  //代码属性编码	varchar
	F009V string  //代码属性	varchar
	F010V string  //上市状态编码	varchar
	F011V string  //上市状态	varchar
	F012N float64 //面值	decimal		单位：元
	F013V string  //ISIN	varchar
}

//获取所有股票基本信息
func GetAllStockBaseInfo(baseInfos *[]BaseInfo) error {
	params := map[string]string{}
	return cninfo.GetInfoByScodeDate(APIBaseInfo, params, baseInfos, cap(*baseInfos))
}

const APIStockSectorInfo = "http://webapi.cninfo.com.cn/api/stock/p_stock0004"

//StockSectorInfo 股票所属板块 "http://webapi.cninfo.com.cn/api/stock/p_stock0004"
//scode	股票代码	string	是	输入不超过300只股票代码，用逗号分隔；如： 000001,600000
//typecode	类别代码	string	否	可以传入多个类别代码，用逗号分隔， 编码：137001 市场分类 137002 证监会行业分类 137004 申银万国行业分类 137005 新财富行业分类 137006 地区省市分类 137007 指数成份股 137008 概念板块
type StockSectorInfo struct {
	SECCODE string //证券代码	varchar
	SECNAME string //证券简称	varchar

	F001V string //分类标准编码	varchar
	F002V string //分类标准	varchar
	F003V string //板块编码	varchar
	F004V string //板块一类名称	varchar
	F005V string //板块二类名称	varchar
	F006V string //板块三类名称	varchar
	F007V string //板块四类名称	varchar
	F008V string //板块五类名称	varchar
	F009V string //板块一类编码	varchar
	F010V string //板块二类编码	varchar
	F011V string //板块三类编码	varchar
	F012V string //板块四类编码	varchar
	F013V string //板块五类编码	varchar
}

const APIStockFinanceDetail = "http://webapi.cninfo.com.cn/api/stock/p_rzrq3104"

//StockFinanceDetail 融资融券明细数据 "http://webapi.cninfo.com.cn/api/stock/p_rzrq3104"
//params:	scode	股票代码	string	输入不超过300只股票代码，用逗号分隔；如： 000001,600000
//			sdate	开始变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
type StockFinanceDetail struct {
	TRADEDATE string //交易日期	datetime
	SECCODE   string //标的证券代码	varchar
	SECNAME   string //标的证券简称	varchar

	F001N float64 //本日融资余额	decimal		单位：元
	F002N float64 //本日融资买入额	decimal		单位：元
	F003N float64 //本日融资偿还额	decimal		单位：元
	F004N uint64  //本日融券余量	decimal		单位：股
	F006N uint64  //本日融券卖出量	decimal		单位：股
	F007N uint64  //本日融券偿还量	decimal		单位：股
	F008N float64 //融券余量金额	decimal		单位：元
	F009N float64 //融资融券余额	decimal		单位：元
	F010V string  //数据来源	varchar
	F011V string  //证券类别编码	varchar
	F012V string  //证券类别	varchar
	MEMO  string  //备注	varchar
}

const APIStockDetail = "http://webapi.cninfo.com.cn/api/stock/p_stock2215"

//StockDetail 融资融券明细数据 "http://webapi.cninfo.com.cn/api/stock/p_stock2215"
//params:	scode	股票代码	string	输入不超过300只股票代码，用逗号分隔；如： 000001,600000
//			sdate	开始变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
type StockDetail struct {
	SECCODE     string //证券代码	varchar
	SECNAME     string //证券简称	varchar
	ORGNAME     string //机构名称	varchar
	DECLAREDATE string //公告日期	date
	VARYDATE    string //变动日期	date

	F001V string  //变动原因编码	varchar
	F002V string  //变动原因	varchar
	F003N float64 //总股本	decimal		单位：万股
	F004N float64 //未流通股份	decimal		单位：万股;
	F005N float64 //发起人股份	decimal		单位：万股；
	F006N float64 //国家持股	decimal		单位：万股
	F007N float64 //国有法人持股	decimal		单位：万股
	F008N float64 //境内法人持股	decimal		单位：万股
	F009N float64 //境外法人持股	decimal		单位：万股
	F010N float64 //自然人持股	decimal		单位：万股
	F011N float64 //募集法人股	decimal		单位：万股
	F012N float64 //内部职工股	decimal		单位：万股
	F013N float64 //转配股	decimal		单位：万股
	F014N float64 //其他流通受限股份	decimal		单位：万股
	F015N float64 //优先股	decimal		单位：万股
	F016N float64 //其他未流通股	decimal		单位：万股
	F021N float64 //已流通股份	decimal		单位：万股；
	F022N float64 //人民币普通股	decimal		单位：万股
	F023N float64 //境内上市外资股（B股）	decimal		单位：万股
	F024N float64 //境外上市外资股（H股）	decimal		单位：万股
	F025N float64 //高管股	decimal		单位：万股
	F026N float64 //其他流通股	decimal		单位：万股
	F028N float64 //流通受限股份	decimal		单位：万股
	F017N float64 //配售法人股	decimal		单位：万股
	F018N float64 //战略投资者持股	decimal		单位：万股
	F019N float64 //证券投资基金持股	decimal		单位：万股
	F020N float64 //一般法人持股	decimal		单位：万股
	F029N float64 //国家持股（受限）	decimal		单位：万股
	F030N float64 //国有法人持股（受限）	decimal		单位：万股
	F031N float64 //其他内资持股（受限）	decimal		单位：万股
	F032N float64 //其中：境内法人持股	decimal		单位：万股
	F033N float64 //其中：境内自然人持股	decimal		单位：万股
	F034N float64 //外资持股（受限）	decimal		单位：万股
	F035N float64 //其中：境外法人持股	decimal		单位：万股
	F036N float64 //其中：境外自然人持股	decimal		单位：万股
	F037N float64 //其中：限售高管股	decimal		其中：限售高管股
	F038N float64 //其中：限售B股	decimal		其中：限售B股
	F040N float64 //其中：限售H股	decimal		其中：限售H股
	F027C string  //最新记录标识	char		0-否，1-是；
	F049N float64 //其他	DECIMAL(12,4)		单位：万股，仅适用于北交所上市公司
	F050N float64 //控股股东、实际控制人	DECIMAL(12,4)		单位：万股，仅适用于北交所上市公司
}

const APIDetailNew = "http://webapi.cninfo.com.cn/api/stock/p_stock2401"
const APIDetailHistory = "http://webapi.cninfo.com.cn/api/stock/p_stock2402"

//Detail
//股票最新日行情 "http://webapi.cninfo.com.cn/api/stock/p_stock2401"
//params:	scode	股票代码	string	否	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
//股票历史日行情 "http://webapi.cninfo.com.cn/api/stock/p_stock2402"
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type Detail struct {
	SECCODE   string //证券代码	varchar
	SECNAME   string //证券简称	varchar
	TRADEDATE string //交易日期	date

	F001V string  //交易所	varchar
	F002N float64 //昨日收盘价	decimal		单位：元
	F003N float64 //今日开盘价	decimal		单位：元
	F004N uint64  //成交数量	decimal		单位：股
	F005N float64 //最高成交价	decimal		单位：元
	F006N float64 //最低成交价	decimal		单位：元
	F007N float64 //最近成交价	decimal		单位：元
	F008N uint64  //总笔数	decimal		单位：笔
	F009N float64 //涨跌	decimal		单位：元
	F010N float64 //涨跌幅	decimal		单位：%
	F011N float64 //成交金额	decimal		单位：元
	F012N float64 //换手率	decimal
	F013N float64 //振幅	decimal		单位：%
	F020N uint64  //发行总股本	decimal		单位：股
	F021N uint64  //流通股本	decimal		单位：股
	F026N float64 //市盈率	decimal
}

const APIBlockTrade = "http://webapi.cninfo.com.cn/api/stock/p_stock2416"

//BlockTrade 大宗交易数据 "http://webapi.cninfo.com.cn/api/stock/p_stock2416"
//params: 	scode	股票代码	string	股票代码不为空，交易日期为空时，则查询股票的所有交易数据
//			tdate	交易日期	string	交易日期与股票代码不能同时为空，输入一个交易日期，选择该日期所有数据
type BlockTrade struct {
	SECCODE   string //证券代码	VARCHAR
	SECNAME   string //证券简称	VARCHAR
	TRADEDATE string //交易日期	DATE

	F001V string  //交易所	VARCHAR
	F007N uint64  //序号	NUMBER
	F002V string  //买方营业部	VARCHAR
	F003V string  //卖方营业部	VARCHAR
	F004N float64 //成交价格	NUMBER
	F005N uint64  //成交量	NUMBER
	F006N float64 //成交金额	NUMBER
}
