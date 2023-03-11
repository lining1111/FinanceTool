package BSE

const APIBaseInfo = "http://webapi.cninfo.com.cn/api/stock/p_stock2101_BSE"

//BaseInfo 股票基本信息 "http://webapi.cninfo.com.cn/api/stock/p_stock2101_BSE"
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
	F007N float64 `excel:"初始上市数量"` //初始上市数量
	F008V string  `excel:"代码属性编码"` //代码属性编码
	F009V string  `excel:"代码属性"`   //代码属性
	F010V string  `excel:"上市状态编码"` //上市状态编码
	F011V string  `excel:"上市状态"`   //上市状态
	F012N float64 `excel:"面值"`     //面值
	F013V string  `excel:"ISIN"`   //ISIN
}

const APIStockFinanceDetail = "http://webapi.cninfo.com.cn/api/stock/p_rzrq3104_BSE"

//StockFinanceDetail 融资融券明细数据 "http://webapi.cninfo.com.cn/api/stock/p_rzrq3104_BSE"
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
	F009N float64 `excel:"融资融券余额"`  //融资融券余额	单位：元
	F010V string  `excel:"数据来源"`    //数据来源
	F011V string  `excel:"证券类别编码"`  //证券类别编码
	F012V string  `excel:"证券类别"`    //证券类别
	MEMO  string  `excel:"备注"`      //备注
}

const APIDetailNew = "http://webapi.cninfo.com.cn/api/stock/p_stock2401_BSE"
const APIDetailHistory = "http://webapi.cninfo.com.cn/api/stock/p_stock2402_BSE"

//Detail
//股票最新日行情 "http://webapi.cninfo.com.cn/api/stock/p_stock2401_BSE"
//params:	scode	股票代码	string	否	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
//股票历史日行情 "http://webapi.cninfo.com.cn/api/stock/p_stock2402_BSE"
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type Detail struct {
	SECCODE   string `excel:"证券代码"` //证券代码
	SECNAME   string `excel:"证券简称"` //证券简称
	TRADEDATE string `excel:"交易日期"` //交易日期

	F001V string  `excel:"交易所"`   //交易所
	F002N float64 `excel:"昨收盘"`   //昨收盘 单位：元
	F003N float64 `excel:"开盘价"`   //开盘价 单位：元
	F004N uint64  `excel:"成交数量"`  //成交数量 单位：股
	F005N float64 `excel:"最高价"`   //最高价 单位：元
	F006N float64 `excel:"最低价"`   //最低价 单位：元
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

const APIBlockTrade = "http://webapi.cninfo.com.cn/api/stock/p_stock2416_BSE"

//BlockTrade 大宗交易数据 "http://webapi.cninfo.com.cn/api/stock/p_stock2416_BSE"
//params: 	scode	股票代码	string	股票代码不为空，交易日期为空时，则查询股票的所有交易数据
//			tdate	交易日期	string	交易日期与股票代码不能同时为空，输入一个交易日期，选择该日期所有数据
type BlockTrade struct {
	SECCODE   string `excel:"证券代码"` //证券代码
	SECNAME   string `excel:"证券简称"` //证券简称
	TRADEDATE string `excel:"交易日期"` //交易日期

	F001V string  `excel:"交易所"`   //交易所
	F007N float64 `excel:"序号"`    //序号
	F002V string  `excel:"买方营业部"` //买方营业部
	F003V string  `excel:"卖方营业部"` //卖方营业部
	F004N float64 `excel:"成交价格"`  //成交价格
	F005N float64 `excel:"成交量"`   //成交量
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
