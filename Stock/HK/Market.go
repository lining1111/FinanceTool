package HK

const APIBaseInfo = "http://webapi.cninfo.com.cn/api/hk/p_hk4039"

//BaseInfo 股票基本信息 "http://webapi.cninfo.com.cn/api/hk/p_hk4039"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
type BaseInfo struct {
	SECCODE string `excel:"证券代码"` //证券代码
	SECNAME string `excel:"证券简称"` //证券简称

	F001V string  `excel:"拼音简称"`   //拼音简称
	F002V string  `excel:"交易市场编码"` //交易市场编码
	F003V string  `excel:"交易市场"`   //交易市场
	F004D string  `excel:"上市日期"`   //上市日期
	F005D string  `excel:"终止上市日期"` //终止上市日期
	F006V string  `excel:"发行机构ID"` //发行机构ID
	F007V string  `excel:"发行机构名称"` //发行机构名称
	F008V string  `excel:"上市状态编码"` //上市状态编码
	F009V string  `excel:"上市状态"`   //上市状态
	F010V string  `excel:"ISIN代码"` //ISIN代码
	F011N float64 `excel:"面值"`     //面值
	MEMO  string  `excel:"备注"`     //备注
}

const APIStockDetail = "http://webapi.cninfo.com.cn/api/hk/p_hk4045"

//StockDetail 港股股本表 "http://webapi.cninfo.com.cn/api/hk/p_hk4045"
//params:	scode	股票代码	string	输入不超过300只股票代码，用逗号分隔；如： 000001,600000
//			sdate	开始变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束变动日期	string	否	支持格式示例：20161101 或2016-11-01 或2016/11/01
type StockDetail struct {
	SECCODE     string `excel:"证券代码"`            //证券代码
	SECNAME     string `excel:"香港交易证券简称 (中文简体)"` //香港交易证券简称 (中文简体)
	DECLAREDATE string `excel:"公告日期"`            //公告日期
	VARYDATE    string `excel:"变动日期"`            //变动日期
	F001N       string `excel:"已发行股本数量"`         //已发行股本数量
	F002N       uint64 `excel:"总股本数量"`           //总股本数量
	F003N       uint64 `excel:"自由流通股本数量"`        //自由流通股本数量
	F004N       uint64 `excel:"自由流通股本占总股本"`      //自由流通股本占总股本
	F005V       string `excel:"变动原因代码"`          //变动原因代码
}

const APIDetailDay = "http://webapi.cninfo.com.cn/api/hk/p_hk4026"

//DetailDay
//香港股票行情日报 "http://webapi.cninfo.com.cn/api/hk/p_hk4026"
//params:	scode	股票代码	string	否	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type DetailDay struct {
	SECCODE string `excel:"证券代码"` //证券代码
	SECNAME string `excel:"证券简称"` //证券简称

	F001N     float64 `excel:"开市价"`    //开市价
	F002N     float64 `excel:"最高成交价"`  //最高成交价
	F003N     float64 `excel:"最低成交价"`  //最低成交价
	F004N     float64 `excel:"收市价"`    //收市价
	F005N     float64 `excel:"收市价经调整"` //收市价经调整
	F006N     uint64  `excel:"成交股数"`   //成交股数
	F007N     float64 `excel:"成交金额"`   //成交金额
	F008C     string  `excel:"证券状态"`   //证券状态
	F009N     float64 `excel:"最新市值"`   //最新市值
	F010N     float64 `excel:"涨跌额"`    //涨跌额
	F011N     float64 `excel:"涨跌幅"`    //涨跌幅
	TRADEDATE string  `excel:"交易日期"`   //交易日期
	MEMO      string  `excel:"备注"`     //备注
}

const APIDetailWeek = "http://webapi.cninfo.com.cn/api/hk/p_hk4027"

//DetailWeek
//香港股票行情周报 "http://webapi.cninfo.com.cn/api/hk/p_hk4027"
//params:	scode	股票代码	string	否	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type DetailWeek struct {
	SECCODE string `excel:"证券代码"` //证券代码
	SECNAME string `excel:"证券简称"` //证券简称

	F001N     float64 `excel:"开市价"`    //开市价
	F002N     float64 `excel:"最高成交价"`  //最高成交价
	F003N     float64 `excel:"最低成交价"`  //最低成交价
	F004N     float64 `excel:"收市价"`    //收市价
	F005N     float64 `excel:"收市价经调整"` //收市价经调整
	F006N     uint64  `excel:"成交股数"`   //成交股数
	F007N     float64 `excel:"成交金额"`   //成交金额
	F008C     string  `excel:"证券状态"`   //证券状态
	F009N     float64 `excel:"最新市值"`   //最新市值
	F010N     float64 `excel:"涨跌额"`    //涨跌额
	F011N     float64 `excel:"涨跌幅"`    //涨跌幅
	STARTDATE string  `excel:"开始日期"`   //开始日期
	ENDDATE   string  `excel:"截止日期"`   //截止日期
	MEMO      string  `excel:"备注"`     //备注
}

const APIDetailMonth = "http://webapi.cninfo.com.cn/api/hk/p_hk4028"

//DetailMonth
//香港股票行情周报 "http://webapi.cninfo.com.cn/api/hk/p_hk4028"
//params:	scode	股票代码	string	否	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type DetailMonth struct {
	SECCODE string `excel:"证券代码"` //证券代码
	SECNAME string `excel:"证券简称"` //证券简称

	F001N     float64 `excel:"开市价"`    //开市价
	F002N     float64 `excel:"最高成交价"`  //最高成交价
	F003N     float64 `excel:"最低成交价"`  //最低成交价
	F004N     float64 `excel:"收市价"`    //收市价
	F005N     float64 `excel:"收市价经调整"` //收市价经调整
	F006N     uint64  `excel:"成交股数"`   //成交股数
	F007N     float64 `excel:"成交金额"`   //成交金额
	F008C     string  `excel:"证券状态"`   //证券状态
	F009N     float64 `excel:"最新市值"`   //最新市值
	F010N     float64 `excel:"涨跌额"`    //涨跌额
	F011N     float64 `excel:"涨跌幅"`    //涨跌幅
	STARTDATE string  `excel:"开始日期"`   //开始日期
	ENDDATE   string  `excel:"截止日期"`   //截止日期
	MEMO      string  `excel:"备注"`     //备注
}

const APIDetailQX = "http://webapi.cninfo.com.cn/api/hk/p_hk4029"

//DetailQX
//香港股票行情周报 "http://webapi.cninfo.com.cn/api/hk/p_hk4029"
//params:	scode	股票代码	string	否	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type DetailQX struct {
	SECCODE string `excel:"证券代码"` //证券代码
	SECNAME string `excel:"证券简称"` //证券简称

	F001N     float64 `excel:"开市价"`    //开市价
	F002N     float64 `excel:"最高成交价"`  //最高成交价
	F003N     float64 `excel:"最低成交价"`  //最低成交价
	F004N     float64 `excel:"收市价"`    //收市价
	F005N     float64 `excel:"收市价经调整"` //收市价经调整
	F006N     uint64  `excel:"成交股数"`   //成交股数
	F007N     float64 `excel:"成交金额"`   //成交金额
	F008C     string  `excel:"证券状态"`   //证券状态
	F009N     float64 `excel:"最新市值"`   //最新市值
	F010N     float64 `excel:"涨跌额"`    //涨跌额
	F011N     float64 `excel:"涨跌幅"`    //涨跌幅
	STARTDATE string  `excel:"开始日期"`   //开始日期
	ENDDATE   string  `excel:"截止日期"`   //截止日期
	MEMO      string  `excel:"备注"`     //备注
}

const APIDetailYear = "http://webapi.cninfo.com.cn/api/hk/p_hk4030"

//DetailYear
//香港股票行情周报 "http://webapi.cninfo.com.cn/api/hk/p_hk4030"
//params:	scode	股票代码	string	否	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type DetailYear struct {
	SECCODE string `excel:"证券代码"` //证券代码
	SECNAME string `excel:"证券简称"` //证券简称

	F001N     float64 `excel:"开市价"`    //开市价
	F002N     float64 `excel:"最高成交价"`  //最高成交价
	F003N     float64 `excel:"最低成交价"`  //最低成交价
	F004N     float64 `excel:"收市价"`    //收市价
	F005N     float64 `excel:"收市价经调整"` //收市价经调整
	F006N     uint64  `excel:"成交股数"`   //成交股数
	F007N     float64 `excel:"成交金额"`   //成交金额
	F008C     string  `excel:"证券状态"`   //证券状态
	F009N     float64 `excel:"最新市值"`   //最新市值
	F010N     float64 `excel:"涨跌额"`    //涨跌额
	F011N     float64 `excel:"涨跌幅"`    //涨跌幅
	STARTDATE string  `excel:"开始日期"`   //开始日期
	ENDDATE   string  `excel:"截止日期"`   //截止日期
	MEMO      string  `excel:"备注"`     //备注
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
