package XSB

const APIBaseInfo = "http://webapi.cninfo.com.cn/api/neeq/p_neeq6003"

//BaseInfo 新三板证券信息表 "http://webapi.cninfo.com.cn/api/neeq/p_neeq6003"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；如： 000001,600000
type BaseInfo struct {
	SECCODE string //证券代码	varchar
	SECNAME string //证券简称	varchar

	F001V string  //拼音简称	varchar
	F002V string  //证券类别编码	varchar		001004股份报价001010代办转让
	F003V string  //证券类别	varchar		详见证券类别编码备注
	F005V string  //交易市场编码	varchar		012005代办转让012008股份报价系统
	F006V string  //交易市场	varchar		详见交易市场编码备注
	F007D string  //挂牌日期	date
	F008D string  //摘牌日期	date
	F009N uint64  //初始挂牌数量	varchar		单位：股
	F015V string  //发行机构名称	varchar
	F019N float64 //面值	decimal		单位：元
	F017V string  //挂牌状态	varchar		基础层挂牌、终止挂牌、创新层挂牌
	F018V string  //转让方式	varchar		协议、做市
	F020V string  //行业编码	varchar		008001证监会行业分类标准
	F021V string  //一级行业	varchar
	F022V string  //二级行业	varchar
	MEMO  string  //备注	varchar
}

const APIDetail = "http://webapi.cninfo.com.cn/api/neeq/p_neeq6028"

//Detail
//新三板股份报价日行情信息 "http://webapi.cninfo.com.cn/api/neeq/p_neeq6028"
//params:	scode	股票代码	string	输入1只股票代码，用逗号分隔；如： 000001
//			sdate	开始查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束查询日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type Detail struct {
	TRADEDATE string //交易日期	date
	SECCODE   string //证券代码	varchar
	SECNAME   string //证券简称	varchar

	F001N float64 //昨日收盘价	decimal		单位：元
	F002N float64 //今日开盘价	decimal		单位：元
	F003N float64 //最近成交价	decimal		单位：元
	F004N uint64  //成交数量	decimal		单位：股
	F005N float64 //成交金额(原币)	decimal
	F007N float64 //最高成交价	decimal		单位：元
	F008N float64 //最低成交价	decimal		单位：元
	F012N float64 //涨跌幅	decimal		单位：%
	MEMO  string  //备注	varchar
}
