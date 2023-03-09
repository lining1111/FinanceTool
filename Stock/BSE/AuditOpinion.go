package BSE

const APIAO = "http://webapi.cninfo.com.cn/api/stock/p_stock2239_BSE"

//AuditOpinion 审计意见
//"http://webapi.cninfo.com.cn/api/stock/p_stock2239_BSE"
//
//params:	scode	股票代码		string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始变动日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束变动日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type AuditOpinion struct {
	SECNAME     string //证券简称
	SECCODE     string //证券代码
	DECLAREDATE string //公告日期

	F001D string //报告年度
	F002C string //是否经审计
	F003V string //境内会计师事务所ID
	F004V string //境内会计师事务所名称
	F005V string //境内会计师事务所签名注册会计师
	F006V string //境内会计师事务所审计意见类型编码
	F007V string //境内会计师事务所审计意见类型
	F008V string //非标准审计意见的事项内容

	F009V string //境外会计师事务所ID
	F010V string //境外会计师事务所名称
	F011V string //境外会计师事务所签名注册会计师
	F012V string //境外审计报告会计师事务所审计意见类型编码
	F013V string //境外审计报告会计师事务所审计意见类型
	F014V string //境外审计意见非标准的事项说明

	F015V string //董事会对会计师事务所出具的有保留意见书、解释说明
	F016V string //监事会对董事会就上述事项的说明所表示的意见
}
