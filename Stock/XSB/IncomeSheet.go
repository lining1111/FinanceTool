package XSB

const APIIS = "http://webapi.cninfo.com.cn/api/stock/p_stock2301"

//IncomeSheet 利润表 "http://webapi.cninfo.com.cn/api/stock/p_stock2301"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			sdate	开始查询日期	string	否	开始查询日期 可为空，F001D
//			edate	结束查询日期	string	否	结束查询日期 可为空，F001D
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type IncomeSheet struct {
	SECCODE     string `excel:" 证券代码 "` //证券代码
	SECNAME     string `excel:" 证券简称 "` //证券简称
	DECLAREDATE string `excel:" 公告日期 "` //公告日期
	STARTDATE   string `excel:" 开始日期 "` //开始日期
	ENDDATE     string `excel:" 截止日期 "` //截止日期

	F001D string  `excel:"报告年度"`               //报告年度
	F002V string  `excel:"合并类型编码"`             //合并类型编码 071001合并本期071002合并上期071003母公司本期071004母公司上期
	F003V string  `excel:"合并类型"`               //合并类型 详见合并类型编码备注
	F004V string  `excel:"报表来源编码"`             //报表来源编码 033003定期报告033006转让说明书033013其他
	F005V string  `excel:"报表来源"`               //报表来源 详见报表来源编码备注
	F035N float64 `excel:"一、营业总收入"`            //一、营业总收入 单位：元
	F006N float64 `excel:"其中：营业收入"`            //其中：营业收入 单位：元
	F033N float64 `excel:"利息收入"`               //利息收入 单位：元
	F034N float64 `excel:"已赚保费"`               //已赚保费 单位：元
	F042N float64 `excel:"手续费及佣金收入"`           //手续费及佣金收入 单位：元
	F036N float64 `excel:"二、营业总成本"`            //二、营业总成本 单位：元
	F007N float64 `excel:"其中：营业成本"`            //其中：营业成本 单位：元
	F043N float64 `excel:"利息支出"`               //利息支出 单位：元
	F044N float64 `excel:"手续费及佣金支出"`           //手续费及佣金支出 单位：元
	F045N float64 `excel:"退保金"`                //退保金 单位：元
	F046N float64 `excel:"赔付支出净额"`             //赔付支出净额 单位：元
	F047N float64 `excel:"提取保险合同准备金净额"`        //提取保险合同准备金净额 单位：元
	F048N float64 `excel:"保单红利支出"`             //保单红利支出 单位：元
	F049N float64 `excel:"分保费用"`               //分保费用 单位：元
	F008N float64 `excel:"营业税金及附加"`            //营业税金及附加 单位：元
	F009N float64 `excel:"销售费用"`               //销售费用 单位：元
	F010N float64 `excel:"管理费用"`               //管理费用 单位：元
	F011N float64 `excel:"堪探费用"`               //堪探费用 单位：元
	F012N float64 `excel:"财务费用"`               //财务费用 单位：元
	F013N float64 `excel:"资产减值损失"`             //资产减值损失 单位：元
	F014N float64 `excel:"加:公允价值变动净收益"`        //加:公允价值变动净收益 单位：元
	F015N float64 `excel:"投资收益"`               //投资收益 单位：元
	F016N float64 `excel:"其中：对联营企业和合营企业的投资收益"` //其中：对联营企业和合营企业的投资收益 单位：元
	F037N float64 `excel:"汇兑收益"`               //汇兑收益 单位：元
	F017N float64 `excel:"影响营业利润的其他科目"`        //影响营业利润的其他科目 单位：元
	F018N float64 `excel:"三、营业利润"`             //三、营业利润 单位：元
	F019N float64 `excel:"加:补贴收入"`             //加：补贴收入 单位：元
	F020N float64 `excel:"营业外收入"`              //营业外收入 单位：元
	F050N float64 `excel:"其中:非流动资产处置利得"`       //其中：非流动资产处置利得 单位：元
	F021N float64 `excel:"减:营业外支出"`            //减：营业外支出 单位：元
	F022N float64 `excel:"其中:非流动资产处置损失"`       //其中：非流动资产处置损失 单位：元
	F023N float64 `excel:"加:影响利润总额的其他科目"`      //加：影响利润总额的其他科目 单位：元
	F024N float64 `excel:"四、利润总额"`             //四、利润总额 单位：元
	F025N float64 `excel:"减:所得税"`              //减：所得税 单位：元
	F026N float64 `excel:"加:影响净利润的其他科目"`       //加:影响净利润的其他科目 单位：元
	F027N float64 `excel:"五、净利润"`              //五、净利润 单位：元
	F028N float64 `excel:"归属于母公司所有者的净利润"`      //归属于母公司所有者的净利润 单位：元
	F029N float64 `excel:"少数股东损益"`             //少数股东损益 单位：元
	F030N float64 `excel:"六、每股收益"`             //六、每股收益 单位：元
	F031N float64 `excel:"(一)基本每股收益"`          //（一）基本每股收益 单位：元
	F032N float64 `excel:"(二)稀释每股收益"`          //（二）稀释每股收益 单位：元
	F038N float64 `excel:"七、其他综合收益"`           //七、其他综合收益 单位：元
	F039N float64 `excel:"八、综合收益总额"`           //八、综合收益总额 单位：元
	F040N float64 `excel:"其中:归属于母公司"`          //其中：归属于母公司 单位：元
	F041N float64 `excel:"其中:归属于少数股东"`         //其中：归属于少数股东 单位：元
	MEMO  string  `excel:"备注"`                 //备注
	F051N float64 `excel:"其他收益"`               //其他收益 元	20210826新增
	F052N float64 `excel:"资产处置收益"`             //资产处置收益 元	20210826新增
	F053N float64 `excel:"持续经营净利润"`            //持续经营净利润 元	20210826新增
	F054N float64 `excel:"终止经营净利润"`            //终止经营净利润 元	20210826新增
	F055N float64 `excel:"研发费用"`               //研发费用 元	20210826新增
	F056N float64 `excel:"财务费用-利息费用"`          //财务费用-利息费用 元	20210826新增
	F057N float64 `excel:"财务费用-利息收入"`          //财务费用-利息收入 元	20210826新增
	F058N float64 `excel:"信用减值损失"`             //信用减值损失 元	20210826新增
	F059N float64 `excel:"净敞口套期收益"`            //净敞口套期收益 元	20210826新增
	F060N float64 `excel:"营业收入-利息净收入"`         //营业收入-利息净收入 元	20210826新增
	F061N float64 `excel:"营业收入-手续费及佣金净收收入"`    //营业收入-手续费及佣金净收收入 元	20210826新增
	F062N float64 `excel:"营业收入-其他业务收入"`        //营业收入-其他业务收入 元	20210826新增
	F063N float64 `excel:"营业收入-其他业务成本"`        //营业收入-其他业务成本 元	20210826新增
	F064N float64 `excel:"以摊余成本计量的金融资产终止确认收益"` //以摊余成本计量的金融资产终止确认收益 元	20210826新增
}
