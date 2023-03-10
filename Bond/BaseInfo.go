package Bond

const urlApi = "http://webapi.cninfo.com.cn/api/bond/p_bond2801"

//BaseInfo 债券基本信息 "http://webapi.cninfo.com.cn/api/bond/p_bond2801"
//参数说明 	scode		债券代码	string		输入1个债券代码
//			platecode	交易市场	string		上交所 :012001 深交所 :012002 银行间 :012004
type BaseInfo struct {
	SECCODE  string `excel:"债券代码"` //债券代码
	SECNAME  string `excel:"债券简称"` //债券简称
	BONDNAME string `excel:"债券名称"` //债券名称

	F001V string  `excel:"英文全称"`     //英文全称
	F002V string  `excel:"债券类型编码"`   //债券类型编码	可通过p_public0005获取
	F003V string  `excel:"债券类型"`     //债券类型
	F004N string  `excel:"偿还期限"`     //偿还期限
	F005V string  `excel:"付息方式编码"`   //付息方式编码	通过p_public0006可获取，对应的总类编码为‘009’
	F006V string  `excel:"付息方式"`     //付息方式
	F007N float64 `excel:"付息频率"`     //付息频率
	F008V string  `excel:"债券形式编码"`   //债券形式编码
	F009V string  `excel:"债券形式"`     //债券形式
	F012V string  `excel:"利率类型编码"`   //利率类型编码
	F013V string  `excel:"利率类型"`     //利率类型
	F014V string  `excel:"计利方式"`     //计利方式
	F015D string  `excel:"起息日"`      //起息日
	F016D string  `excel:"到期日"`      //到期日
	F017V string  `excel:"付息日期"`     //付息日期
	F018D string  `excel:"兑付日"`      //兑付日
	F019V string  `excel:"兑付说明"`     //兑付说明
	F024V string  `excel:"发行机构ID"`   //发行机构ID
	F025V string  `excel:"发行机构名称"`   //发行机构名称
	F026N float64 `excel:"系列债券关联标识"` //系列债券关联标识
	F029V string  `excel:"交易市场编码"`   //交易市场编码
	F031V string  `excel:"交易市场"`     //交易市场
	F032V string  `excel:"货币名称"`     //货币名称
	F033V string  `excel:"货币名称编码"`   //货币名称编码
	F034V string  `excel:"年度计息基准"`   //年度计息基准
	F035V string  `excel:"年度计息基准编码"` //年度计息基准编码
	F043V string  `excel:"债分类一级代码"`  //债分类一级代码
	F044V string  `excel:"债分类一级名称"`  //债分类一级名称
	F045V string  `excel:"债分类二级代码"`  //债分类二级代码
	F046V string  `excel:"债分类二级名称"`  //债分类二级名称
	F047V string  `excel:"固收分类代码"`   //固收分类代码
	F048V string  `excel:"固收分类名称"`   //固收分类名称
}
