package Bond

const urlApiMarket = "http://webapi.cninfo.com.cn/api/bond/p_bond2804"

//Market 债券行情数据  "http://webapi.cninfo.com.cn/api/bond/p_bond2804"
//params : 	scode		债券代码		string	债券代码，输入1个债券， 为空时返回系统日期当天所有行情数据（不区分交易日，所以多数时间返回的是空值）
//			platecode	交易市场		string	交易市场 ,可为空 f026v 上交所 012001 深交所 012002 银行间 012004
//			sdate		开始查询时间	string	开始查询时间 ,可为空 tradedate
//			edate		结束查询时间	string	结束查询时间，可为空， tradedate 默认为当天系统日期
type Market struct {
	SECCODE   string //债券代码
	SECNAME   string //债券简称
	TRADEDATE string //交易日期

	F001V string  //交易所
	F002N float64 //昨日收盘
	F003N float64 //今日开盘
	F004N float64 //成交数量
	F005N float64 //最高成交
	F006N float64 //最低成交
	F007N float64 //最近成交
	F010N float64 //总笔数
	F011N float64 //市盈率1
	F013N float64 //升跌1
	F015N float64 //涨跌幅
	F016N float64 //成交金额(原币)
	F017N float64 //换手率
	F019N float64 //发行总股本
	F020N float64 //流通股本
	F023V string  //证券种类
	F024V string  //货币种类
	F026V string  //交易市场编码
}

const urlApiBlockTrade = "http://webapi.cninfo.com.cn/api/bond/p_bond2803"

//BlockTrade 债券行情数据---大宗交易 "http://webapi.cninfo.com.cn/api/bond/p_bond2803"
//params : 	scode		债券代码		string	债券代码，输入1个债券， 为空时返回系统日期当天所有行情数据（不区分交易日，所以多数时间返回的是空值）
//			platecode	交易市场		string	交易市场 ,可为空 f026v 上交所 012001 深交所 012002 银行间 012004
//			sdate		开始查询时间	string	开始查询时间 ,可为空 tradedate
//			edate		结束查询时间	string	结束查询时间，可为空， tradedate 默认为当天系统日期
type BlockTrade struct {
	SECCODE   string //证券代码
	SECNAME   string //证券简称
	TRADEDATE string //交易日期

	F001V string  //交易所
	F002V string  //买方营业部
	F003V string  //卖方营业部
	F004N float64 //成交价格
	F005N float64 //成交量
	F006N float64 //成交金额
	F007N float64 //序号
	F008V string  //交易所编码
}

const urlApiBankTrade = "http://webapi.cninfo.com.cn/api/bond/p_bond2814"

//BankTrade 债券行情数据---银行间汇总 "http://webapi.cninfo.com.cn/api/bond/p_bond2814"
//params:	scode	债券代码	string	输入单个债券代码，查询该债券的全部数据
//			tdate	交易日期	string	输入交易日期，查询该日期行情数据
type BankTrade struct {
	TRADEDATE string //交易日期

	F001V string  //行情时间
	F002V string  //债券代码
	F003N float64 //开盘净价
	F004N float64 //开盘收益率
	F005N float64 //收盘净价
	F006N float64 //收盘收益率
	F007N float64 //最高净价
	F008N float64 //最高收益率
	F009N float64 //最低净价
	F010N float64 //最低收益率
	F011N float64 //加权平均净价
	F012N float64 //加权平均收益率
	F013N float64 //前收盘净价
	F014N float64 //前收盘收益率
	F015N float64 //前加权平均净价
	F016N float64 //前加权平均收益率
	F017N float64 //涨跌幅
	F018N float64 //成交金额
	F019V string  //债券币种
}
