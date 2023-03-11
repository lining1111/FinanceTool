package XSB

const APIBaseInfo = "http://webapi.cninfo.com.cn/api/neeq/p_neeq6003"

//BaseInfo 新三板证券信息表 "http://webapi.cninfo.com.cn/api/neeq/p_neeq6003"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
type BaseInfo struct {
	SECCODE string `excel:"证券代码"` //证券代码
	SECNAME string `excel:"证券简称"` //证券简称

	F001V string  `excel:"拼音简称"`   //拼音简称
	F002V string  `excel:"证券类别编码"` //证券类别编码 001004股份报价001010代办转让
	F003V string  `excel:"证券类别"`   //证券类别 详见证券类别编码备注
	F005V string  `excel:"交易市场编码"` //交易市场编码 012005代办转让012008股份报价系统
	F006V string  `excel:"交易市场"`   //交易市场 详见交易市场编码备注
	F007D string  `excel:"挂牌日期"`   //挂牌日期	date
	F008D string  `excel:"摘牌日期"`   //摘牌日期	date
	F009N uint64  `excel:"初始挂牌数量"` //初始挂牌数量 单位：股
	F015V string  `excel:"发行机构名称"` //发行机构名称
	F019N float64 `excel:"面值"`     //面值 单位：元
	F017V string  `excel:"挂牌状态"`   //挂牌状态 基础层挂牌、终止挂牌、创新层挂牌
	F018V string  `excel:"转让方式"`   //转让方式 协议、做市
	F020V string  `excel:"行业编码"`   //行业编码 008001证监会行业分类标准
	F021V string  `excel:"一级行业"`   //一级行业
	F022V string  `excel:"二级行业"`   //二级行业
	MEMO  string  `excel:"备注"`     //备注
}

const APIDetail = "http://webapi.cninfo.com.cn/api/neeq/p_neeq6028"

//Detail
//新三板股份报价日行情信息 "http://webapi.cninfo.com.cn/api/neeq/p_neeq6028"
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type Detail struct {
	TRADEDATE string `excel:"交易日期"` //交易日期
	SECCODE   string `excel:"证券代码"` //证券代码
	SECNAME   string `excel:"证券简称"` //证券简称

	F001N float64 `excel:"昨日收盘价"`    //昨日收盘价 单位：元
	F002N float64 `excel:"今日开盘价"`    //今日开盘价 单位：元
	F003N float64 `excel:"最近成交价"`    //最近成交价 单位：元
	F004N uint64  `excel:"成交数量"`     //成交数量 单位：股
	F005N float64 `excel:"成交金额(原币)"` //成交金额(原币)	decimal
	F007N float64 `excel:"最高成交价"`    //最高成交价 单位：元
	F008N float64 `excel:"最低成交价"`    //最低成交价 单位：元
	F012N float64 `excel:"涨跌幅"`      //涨跌幅 单位：%
	MEMO  string  `excel:"备注"`       //备注
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
