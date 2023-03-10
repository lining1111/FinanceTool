package XSB

const APIAO = "http://webapi.cninfo.com.cn/api/neeq/p_neeq6015"

//AuditOpinion 审计意见
//"http://webapi.cninfo.com.cn/api/neeq/p_neeq6015"
//
//params:	scode	股票代码		string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始变动日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
//			edate	结束变动日期	string	支持格式示例：20161101 或2016-11-01 或2016/11/01
type AuditOpinion struct {
	ORGNAME     string `excel:"机构名称"` //机构名称
	SECNAME     string `excel:"证券简称"` //证券简称
	SECCODE     string `excel:"证券代码"` //证券代码
	DECLAREDATE string `excel:"公告日期"` //公告日期

	F001D string `excel:"报告年度"`             //报告年度 1：是，2：否
	F002C string `excel:"是否经审计"`            //是否经审计
	F003V string `excel:"境内会计师事务所ID"`       //境内会计师事务所ID
	F004V string `excel:"境内会计师事务所名称"`       //境内会计师事务所名称
	F005V string `excel:"境内会计师事务所签名注册会计师"`  //境内会计师事务所签名注册会计师
	F006V string `excel:"境内会计师事务所审计意见类型编码"` //境内会计师事务所审计意见类型编码 042001无保留意见042002保留意见042003保留意见与解释性说明042004否定意见042005拒绝表示意见042006解释性说明042007无法表示意见
	F007V string `excel:"境内会计师事务所审计意见类型"`   //境内会计师事务所审计意见类型 详见境内会计师事务所审计意见类型编码备注

	MEMO string `excel:"备注"` //备注
}
