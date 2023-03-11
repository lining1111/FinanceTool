package HSB

import "FinanceTool/COM/cninfo"

const APIBaseInfo = "http://webapi.cninfo.com.cn/api/stock/p_stock2101"

//BaseInfo 股票基本信息 "http://webapi.cninfo.com.cn/api/stock/p_stock2101"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
type BaseInfo struct {
	ORGNAME string `excel:"机构名称"` //机构名称
	SECCODE string `excel:"证券代码"` //证券代码
	SECNAME string `excel:"证券简称"` //证券简称

	F001V string  `excel:"拼音简称"`   //拼音简称
	F002V string  `excel:"证券类别编码"` //证券类别编码
	F003V string  `excel:"证券类别"`   //证券类别
	F004V string  `excel:"交易市场编码"` //交易市场编码
	F005V string  `excel:"交易市场"`   //交易市场
	F006D string  `excel:"上市日期"`   //上市日期
	F007N uint64  `excel:"初始上市数量"` //初始上市数量 单位：股
	F008V string  `excel:"代码属性编码"` //代码属性编码
	F009V string  `excel:"代码属性"`   //代码属性
	F010V string  `excel:"上市状态编码"` //上市状态编码
	F011V string  `excel:"上市状态"`   //上市状态
	F012N float64 `excel:"面值"`     //面值 单位：元
	F013V string  `excel:"ISIN"`   //ISIN
}

//获取所有股票基本信息
func GetAllStockBaseInfo(baseInfos *[]BaseInfo) error {
	params := map[string]string{}
	return cninfo.GetInfoByScodeDate(APIBaseInfo, params, baseInfos)
}

const APIStockSectorInfo = "http://webapi.cninfo.com.cn/api/stock/p_stock0004"

//StockSectorInfo 股票所属板块 "http://webapi.cninfo.com.cn/api/stock/p_stock0004"
//scode	股票代码	string	是	输入不超过300只股票代码，用逗号分隔；如： 000001,600000
//typecode	类别代码	string	否	可以传入多个类别代码，用逗号分隔， 编码：137001 市场分类 137002 证监会行业分类 137004 申银万国行业分类 137005 新财富行业分类 137006 地区省市分类 137007 指数成份股 137008 概念板块
type StockSectorInfo struct {
	SECCODE string `excel:"证券代码"` //证券代码
	SECNAME string `excel:"证券简称"` //证券简称

	F001V string `excel:"分类标准编码"` //分类标准编码
	F002V string `excel:"分类标准"`   //分类标准
	F003V string `excel:"板块编码"`   //板块编码
	F004V string `excel:"板块一类名称"` //板块一类名称
	F005V string `excel:"板块二类名称"` //板块二类名称
	F006V string `excel:"板块三类名称"` //板块三类名称
	F007V string `excel:"板块四类名称"` //板块四类名称
	F008V string `excel:"板块五类名称"` //板块五类名称
	F009V string `excel:"板块一类编码"` //板块一类编码
	F010V string `excel:"板块二类编码"` //板块二类编码
	F011V string `excel:"板块三类编码"` //板块三类编码
	F012V string `excel:"板块四类编码"` //板块四类编码
	F013V string `excel:"板块五类编码"` //板块五类编码
}

const APIStockFinanceDetail = "http://webapi.cninfo.com.cn/api/stock/p_rzrq3104"

//StockFinanceDetail 融资融券明细数据 "http://webapi.cninfo.com.cn/api/stock/p_rzrq3104"
//params:	scode	股票代码	string	输入不超过300只股票代码，用逗号分隔；如： 000001,600000
//			sdate	开始变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
type StockFinanceDetail struct {
	TRADEDATE string `excel:"交易日期"`   //交易日期
	SECCODE   string `excel:"标的证券代码"` //标的证券代码
	SECNAME   string `excel:"标的证券简称"` //标的证券简称

	F001N float64 `excel:"本日融资余额"`  //本日融资余额 单位：元
	F002N float64 `excel:"本日融资买入额"` //本日融资买入额 单位：元
	F003N float64 `excel:"本日融资偿还额"` //本日融资偿还额 单位：元
	F004N uint64  `excel:"本日融券余量"`  //本日融券余量 单位：股
	F006N uint64  `excel:"本日融券卖出量"` //本日融券卖出量 单位：股
	F007N uint64  `excel:"本日融券偿还量"` //本日融券偿还量 单位：股
	F008N float64 `excel:"融券余量金额"`  //融券余量金额 单位：元
	F009N float64 `excel:"融资融券余额"`  //融资融券余额 单位：元
	F010V string  `excel:"数据来源"`    //数据来源
	F011V string  `excel:"证券类别编码"`  //证券类别编码
	F012V string  `excel:"证券类别"`    //证券类别
	MEMO  string  `excel:"备注"`      //备注
}

const APIStockDetail = "http://webapi.cninfo.com.cn/api/stock/p_stock2215"

//StockDetail 融资融券明细数据 "http://webapi.cninfo.com.cn/api/stock/p_stock2215"
//params:	scode	股票代码	string	输入不超过300只股票代码，用逗号分隔；如： 000001,600000
//			sdate	开始变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
type StockDetail struct {
	SECCODE     string `excel:"证券代码"` //证券代码
	SECNAME     string `excel:"证券简称"` //证券简称
	ORGNAME     string `excel:"机构名称"` //机构名称
	DECLAREDATE string `excel:"公告日期"` //公告日期
	VARYDATE    string `excel:"变动日期"` //变动日期

	F001V string  `excel:"变动原因编码"`      //变动原因编码
	F002V string  `excel:"变动原因"`        //变动原因
	F003N float64 `excel:"总股本"`         //总股本 单位：万股
	F004N float64 `excel:"未流通股份"`       //未流通股份 单位：万股;
	F005N float64 `excel:"发起人股份"`       //发起人股份 单位：万股；
	F006N float64 `excel:"国家持股"`        //国家持股 单位：万股
	F007N float64 `excel:"国有法人持股"`      //国有法人持股 单位：万股
	F008N float64 `excel:"境内法人持股"`      //境内法人持股 单位：万股
	F009N float64 `excel:"境外法人持股"`      //境外法人持股 单位：万股
	F010N float64 `excel:"自然人持股"`       //自然人持股 单位：万股
	F011N float64 `excel:"募集法人股"`       //募集法人股 单位：万股
	F012N float64 `excel:"内部职工股"`       //内部职工股 单位：万股
	F013N float64 `excel:"转配股"`         //转配股 单位：万股
	F014N float64 `excel:"其他流通受限股份"`    //其他流通受限股份 单位：万股
	F015N float64 `excel:"优先股"`         //优先股 单位：万股
	F016N float64 `excel:"其他未流通股"`      //其他未流通股 单位：万股
	F021N float64 `excel:"已流通股份"`       //已流通股份 单位：万股；
	F022N float64 `excel:"人民币普通股"`      //人民币普通股 单位：万股
	F023N float64 `excel:"境内上市外资股（B股）"` //境内上市外资股（B股）单位：万股
	F024N float64 `excel:"境外上市外资股（H股）"` //境外上市外资股（H股）单位：万股
	F025N float64 `excel:"高管股"`         //高管股 单位：万股
	F026N float64 `excel:"其他流通股"`       //其他流通股 单位：万股
	F028N float64 `excel:"流通受限股份"`      //流通受限股份 单位：万股
	F017N float64 `excel:"配售法人股"`       //配售法人股 单位：万股
	F018N float64 `excel:"战略投资者持股"`     //战略投资者持股 单位：万股
	F019N float64 `excel:"证券投资基金持股"`    //证券投资基金持股 单位：万股
	F020N float64 `excel:"一般法人持股"`      //一般法人持股 单位：万股
	F029N float64 `excel:"国家持股（受限）"`    //国家持股（受限） 单位：万股
	F030N float64 `excel:"国有法人持股（受限）"`  //国有法人持股（受限） 单位：万股
	F031N float64 `excel:"其他内资持股（受限）"`  //其他内资持股（受限） 单位：万股
	F032N float64 `excel:"其中：境内法人持股"`   //其中：境内法人持股 单位：万股
	F033N float64 `excel:"其中：境内自然人持股"`  //其中：境内自然人持股	单位：万股
	F034N float64 `excel:"外资持股（受限）"`    //外资持股（受限） 单位：万股
	F035N float64 `excel:"其中：境外法人持股"`   //其中：境外法人持股	单位：万股
	F036N float64 `excel:"其中：境外自然人持股"`  //其中：境外自然人持股	单位：万股
	F037N float64 `excel:"其中：限售高管股"`    //其中：限售高管股 其中：限售高管股
	F038N float64 `excel:"其中：限售B股"`     //其中：限售B股 其中：限售B股
	F040N float64 `excel:"其中：限售H股"`     //其中：限售H股 其中：限售H股
	F027C string  `excel:"最新记录标识"`      //最新记录标识 0-否，1-是；
	F049N float64 `excel:"其他"`          //其他 单位：万股，仅适用于北交所上市公司
	F050N float64 `excel:"控股股东、实际控制人"`  //控股股东、实际控制人 单位：万股，仅适用于北交所上市公司
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
	SECCODE   string `excel:"证券代码"` //证券代码
	SECNAME   string `excel:"证券简称"` //证券简称
	TRADEDATE string `excel:"交易日期"` //交易日期

	F001V string  `excel:"交易所"`   //交易所
	F002N float64 `excel:"昨日收盘价"` //昨日收盘价 单位：元
	F003N float64 `excel:"今日开盘价"` //今日开盘价 单位：元
	F004N uint64  `excel:"成交数量"`  //成交数量 单位：股
	F005N float64 `excel:"最高成交价"` //最高成交价 单位：元
	F006N float64 `excel:"最低成交价"` //最低成交价 单位：元
	F007N float64 `excel:"最近成交价"` //最近成交价 单位：元
	F008N uint64  `excel:"总笔数"`   //总笔数 单位：笔
	F009N float64 `excel:"涨跌"`    //涨跌 单位：元
	F010N float64 `excel:"涨跌幅"`   //涨跌幅 单位：%
	F011N float64 `excel:"成交金额"`  //成交金额 单位：元
	F012N float64 `excel:"换手率"`   //换手率
	F013N float64 `excel:"振幅"`    //振幅 单位：%
	F020N uint64  `excel:"发行总股本"` //发行总股本 单位：股
	F021N uint64  `excel:"流通股本"`  //流通股本 单位：股
	F026N float64 `excel:"市盈率"`   //市盈率
}

const APIBlockTrade = "http://webapi.cninfo.com.cn/api/stock/p_stock2416"

//BlockTrade 大宗交易数据 "http://webapi.cninfo.com.cn/api/stock/p_stock2416"
//params: 	scode	股票代码	string	股票代码不为空，交易日期为空时，则查询股票的所有交易数据
//			tdate	交易日期	string	交易日期与股票代码不能同时为空，输入一个交易日期，选择该日期所有数据
type BlockTrade struct {
	SECCODE   string `excel:"证券代码"` //证券代码
	SECNAME   string `excel:"证券简称"` //证券简称
	TRADEDATE string `excel:"交易日期"` //交易日期

	F001V string  `excel:"交易所"`   //交易所
	F007N uint64  `excel:"序号"`    //序号
	F002V string  `excel:"买方营业部"` //买方营业部
	F003V string  `excel:"卖方营业部"` //卖方营业部
	F004N float64 `excel:"成交价格"`  //成交价格
	F005N uint64  `excel:"成交量"`   //成交量
	F006N float64 `excel:"成交金额"`  //成交金额
}

const APIDataCur = "http://webapi.cninfo.com.cn/api-data/cube/dailyLine"

//股票当时的分时数据 "http://webapi.cninfo.com.cn/api-data/cube/dailyLine"
//stockCode
//_
type DataCurItem struct {
	TIME        uint64  `excel:"时间"`
	OPEN        float64 `excel:"开盘价"`
	CLOSE       float64 `excel:"收盘价"`
	HIGH        float64 `excel:"最高价"`
	LOW         float64 `excel:"最低价"`
	MONEY       uint64  `excel:"交易额"`
	VOL         uint64  `excel:"交易量"`
	KZHANGDIEFU float64 `excel:"涨跌幅"`
	KZHANGDIE   float64 `excel:"涨跌"`
}

type DataCur struct {
	Valuetype []string
	Code      string
	Line      []DataCurItem
}
