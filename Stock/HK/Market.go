package HK

const APIBaseInfo = "http://webapi.cninfo.com.cn/api/hk/p_hk4039"

//BaseInfo 股票基本信息 "http://webapi.cninfo.com.cn/api/hk/p_hk4039"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
type BaseInfo struct {
	SECCODE string //证券代码	varchar
	SECNAME string //证券简称	varchar

	F001V string  //拼音简称	varchar
	F002V string  //交易市场编码	varchar
	F003V string  //交易市场	varchar
	F004D string  //上市日期	date
	F005D string  //终止上市日期	date
	F006V string  //发行机构ID	varchar
	F007V string  //发行机构名称	varchar
	F008V string  //上市状态编码	varchar
	F009V string  //上市状态	varchar
	F010V string  //ISIN代码	varchar
	F011N float64 //面值	decimal
	MEMO  string  //备注	varchar
}

const APIStockDetail = "http://webapi.cninfo.com.cn/api/hk/p_hk4045"

//StockDetail 港股股本表 "http://webapi.cninfo.com.cn/api/hk/p_hk4045"
//params:	scode	股票代码	string	输入不超过300只股票代码，用逗号分隔；如： 000001,600000
//			sdate	开始变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
type StockDetail struct {
	SECCODE     string //证券代码	varchar
	SECNAME     string //香港交易证券简称 (中文简体)	varchar
	DECLAREDATE string //公告日期	date
	VARYDATE    string //变动日期	date
	F001N       string //已发行股本数量	number
	F002N       uint64 //总股本数量	number
	F003N       uint64 //自由流通股本数量	number
	F004N       uint64 //自由流通股本占总股本	number
	F005V       string //变动原因代码	varchar
}

const APIDetailDay = "http://webapi.cninfo.com.cn/api/hk/p_hk4026"

//DetailDay
//香港股票行情日报 "http://webapi.cninfo.com.cn/api/hk/p_hk4026"
//params:	scode	股票代码	string	否	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type DetailDay struct {
	SECCODE string //证券代码	varchar
	SECNAME string //证券简称	varchar

	F001N     float64 //开市价	decimal
	F002N     float64 //最高成交价	decimal
	F003N     float64 //最低成交价	decimal
	F004N     float64 //收市价	decimal
	F005N     float64 //收市价经调整	decimal
	F006N     uint64  //成交股数	decimal
	F007N     float64 //成交金额	decimal
	F008C     string  //证券状态	char
	F009N     float64 //最新市值	decimal
	F010N     float64 //涨跌额	decimal
	F011N     float64 //涨跌幅	decimal
	TRADEDATE string  //交易日期	date
	MEMO      string  //备注	varchar
}

const APIDetailWeek = "http://webapi.cninfo.com.cn/api/hk/p_hk4027"

//DetailWeek
//香港股票行情周报 "http://webapi.cninfo.com.cn/api/hk/p_hk4027"
//params:	scode	股票代码	string	否	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type DetailWeek struct {
	SECCODE string //证券代码	varchar
	SECNAME string //证券简称	varchar

	F001N     float64 //开市价	decimal
	F002N     float64 //最高成交价	decimal
	F003N     float64 //最低成交价	decimal
	F004N     float64 //收市价	decimal
	F005N     float64 //收市价经调整	decimal
	F006N     uint64  //成交股数	decimal
	F007N     float64 //成交金额	decimal
	F008C     string  //证券状态	char
	F009N     float64 //最新市值	decimal
	F010N     float64 //涨跌额	decimal
	F011N     float64 //涨跌幅	decimal
	STARTDATE string  //开始日期	date
	ENDDATE   string  //截止日期	date
	MEMO      string  //备注	varchar
}

const APIDetailMonth = "http://webapi.cninfo.com.cn/api/hk/p_hk4028"

//DetailMonth
//香港股票行情周报 "http://webapi.cninfo.com.cn/api/hk/p_hk4028"
//params:	scode	股票代码	string	否	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type DetailMonth struct {
	SECCODE string //证券代码	varchar
	SECNAME string //证券简称	varchar

	F001N     float64 //开市价	decimal
	F002N     float64 //最高成交价	decimal
	F003N     float64 //最低成交价	decimal
	F004N     float64 //收市价	decimal
	F005N     float64 //收市价经调整	decimal
	F006N     uint64  //成交股数	decimal
	F007N     float64 //成交金额	decimal
	F008C     string  //证券状态	char
	F009N     float64 //最新市值	decimal
	F010N     float64 //涨跌额	decimal
	F011N     float64 //涨跌幅	decimal
	STARTDATE string  //开始日期	date
	ENDDATE   string  //截止日期	date
	MEMO      string  //备注	varchar
}

const APIDetailQX = "http://webapi.cninfo.com.cn/api/hk/p_hk4029"

//DetailQX
//香港股票行情周报 "http://webapi.cninfo.com.cn/api/hk/p_hk4029"
//params:	scode	股票代码	string	否	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type DetailQX struct {
	SECCODE string //证券代码	varchar
	SECNAME string //证券简称	varchar

	F001N     float64 //开市价	decimal
	F002N     float64 //最高成交价	decimal
	F003N     float64 //最低成交价	decimal
	F004N     float64 //收市价	decimal
	F005N     float64 //收市价经调整	decimal
	F006N     uint64  //成交股数	decimal
	F007N     float64 //成交金额	decimal
	F008C     string  //证券状态	char
	F009N     float64 //最新市值	decimal
	F010N     float64 //涨跌额	decimal
	F011N     float64 //涨跌幅	decimal
	STARTDATE string  //开始日期	date
	ENDDATE   string  //截止日期	date
	MEMO      string  //备注	varchar
}

const APIDetailYear = "http://webapi.cninfo.com.cn/api/hk/p_hk4030"

//DetailYear
//香港股票行情周报 "http://webapi.cninfo.com.cn/api/hk/p_hk4030"
//params:	scode	股票代码	string	否	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type DetailYear struct {
	SECCODE string //证券代码	varchar
	SECNAME string //证券简称	varchar

	F001N     float64 //开市价	decimal
	F002N     float64 //最高成交价	decimal
	F003N     float64 //最低成交价	decimal
	F004N     float64 //收市价	decimal
	F005N     float64 //收市价经调整	decimal
	F006N     uint64  //成交股数	decimal
	F007N     float64 //成交金额	decimal
	F008C     string  //证券状态	char
	F009N     float64 //最新市值	decimal
	F010N     float64 //涨跌额	decimal
	F011N     float64 //涨跌幅	decimal
	STARTDATE string  //开始日期	date
	ENDDATE   string  //截止日期	date
	MEMO      string  //备注	varchar
}
