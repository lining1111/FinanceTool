package Bond

const urlApi = "http://webapi.cninfo.com.cn/api/bond/p_bond2801"

//BaseInfo 债券基本信息 "http://webapi.cninfo.com.cn/api/bond/p_bond2801"
//参数说明 	scode		债券代码	string		输入1个债券代码
//			platecode	交易市场	string		上交所 :012001 深交所 :012002 银行间 :012004
type BaseInfo struct {
	SECCODE  string //债券代码
	SECNAME  string //债券简称
	BONDNAME string //债券名称

	F001V string  //英文全称
	F002V string  //债券类型编码	可通过p_public0005获取
	F003V string  //债券类型
	F004N string  //偿还期限
	F005V string  //付息方式编码	通过p_public0006可获取，对应的总类编码为‘009’
	F006V string  //付息方式
	F007N float64 //付息频率
	F008V string  //债券形式编码
	F009V string  //债券形式
	F012V string  //利率类型编码
	F013V string  //利率类型
	F014V string  //计利方式
	F015D string  //起息日
	F016D string  //到期日
	F017V string  //付息日期
	F018D string  //兑付日
	F019V string  //兑付说明
	F024V string  //发行机构ID
	F025V string  //发行机构名称
	F026N float64 //系列债券关联标识
	F029V string  //交易市场编码
	F031V string  //交易市场
	F032V string  //货币名称
	F033V string  //货币名称编码
	F034V string  //年度计息基准
	F035V string  //年度计息基准编码
	F043V string  //债分类一级代码
	F044V string  //债分类一级名称
	F045V string  //债分类二级代码
	F046V string  //债分类二级名称
	F047V string  //固收分类代码
	F048V string  //固收分类名称
}
