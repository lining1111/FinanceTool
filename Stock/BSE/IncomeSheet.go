package BSE

const APIIS = "http://webapi.cninfo.com.cn/api/stock/p_stock2301_BSE"

//IncomeSheet 利润表 "http://webapi.cninfo.com.cn/api/stock/p_stock2301_BSE"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			rdate	报告期	string	报告期 可为空，为空取所有报告期
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type IncomeSheet struct {
	SECNAME     string `excel:"证券简称"` //证券简称
	SECCODE     string `excel:"证券代码"` //证券代码
	ORGNAME     string `excel:"机构名称"` //机构名称
	DECLAREDATE string `excel:"公告日期"` //公告日期
	ENDDATE     string `excel:"截止日期"` //截止日期

	F001D string  `excel:"报告年度"`                    //报告年度
	F002V string  `excel:"合并类型编码"`                  //合并类型编码
	F003V string  `excel:"合并类型"`                    //合并类型
	F004V string  `excel:"报表来源编码"`                  //报表来源编码
	F005V string  `excel:"报表来源"`                    //报表来源
	F035N float64 `excel:"一、营业总收入"`                 //一、营业总收入
	F006N float64 `excel:"其中：营业收入"`                 //其中：营业收入
	F033N float64 `excel:"利息收入"`                    //利息收入
	F034N float64 `excel:"已赚保费"`                    //已赚保费
	F042N float64 `excel:"手续费及佣金收入"`                //手续费及佣金收入
	F036N float64 `excel:"二、营业总成本"`                 //二、营业总成本
	F007N float64 `excel:"其中：营业成本"`                 //其中：营业成本
	F043N float64 `excel:"利息支出"`                    //利息支出
	F044N float64 `excel:"手续费及佣金支出"`                //手续费及佣金支出
	F045N float64 `excel:"退保金"`                     //退保金
	F046N float64 `excel:"赔付支出净额"`                  //赔付支出净额
	F047N float64 `excel:"提取保险合同准备金净额"`             //提取保险合同准备金净额
	F048N float64 `excel:"保单红利支出"`                  //保单红利支出
	F049N float64 `excel:"分保费用"`                    //分保费用
	F008N float64 `excel:"营业税金及附加"`                 //营业税金及附加
	F009N float64 `excel:"销售费用"`                    //销售费用
	F010N float64 `excel:"管理费用"`                    //管理费用
	F011N float64 `excel:"勘探费用"`                    //勘探费用
	F012N float64 `excel:"财务费用"`                    //财务费用
	F056N float64 `excel:"研发费用"`                    //研发费用
	F013N float64 `excel:"资产减值损失"`                  //资产减值损失
	F014N float64 `excel:"加：公允价值变动净收益"`             //加：公允价值变动净收益
	F015N float64 `excel:"投资收益"`                    //投资收益
	F016N float64 `excel:"其中：对联营企业和合营企业的投资收益"`      //其中：对联营企业和合营企业的投资收益
	F037N float64 `excel:"汇兑收益"`                    //汇兑收益
	F051N float64 `excel:"其它收入"`                    //其它收入
	F057N float64 `excel:"信用减值损失"`                  //信用减值损失
	F058N float64 `excel:"净敞口套期收益"`                 //净敞口套期收益
	F059N float64 `excel:"资产处置收益"`                  //资产处置收益
	F017N float64 `excel:"影响营业利润的其他科目"`             //影响营业利润的其他科目
	F018N float64 `excel:"三、营业利润"`                  //三、营业利润
	F019N float64 `excel:"加：补贴收入"`                  //加：补贴收入
	F020N float64 `excel:"营业外收入"`                   //营业外收入
	F050N float64 `excel:"其中：非流动资产处置利得"`            //其中：非流动资产处置利得
	F021N float64 `excel:"减：营业外支出"`                 //减：营业外支出
	F022N float64 `excel:"其中：非流动资产处置损失"`            //其中：非流动资产处置损失
	F023N float64 `excel:"加：影响利润总额的其他科目"`           //加：影响利润总额的其他科目
	F024N float64 `excel:"四、利润总额"`                  //四、利润总额
	F025N float64 `excel:"减：所得税"`                   //减：所得税
	F026N float64 `excel:"加：影响净利润的其他科目"`            //加：影响净利润的其他科目
	F027N float64 `excel:"五、净利润"`                   //五、净利润
	F060N float64 `excel:"持续经营净利润"`                 //持续经营净利润
	F061N float64 `excel:"终止经营净利润"`                 //终止经营净利润
	F028N float64 `excel:"归属于母公司所有者的净利润"`           //归属于母公司所有者的净利润
	F029N float64 `excel:"少数股东损益"`                  //少数股东损益
	F031N float64 `excel:"(一）基本每股收益"`               //(一）基本每股收益
	F032N float64 `excel:"（二）稀释每股收益"`               //（二）稀释每股收益
	F038N float64 `excel:"七、其他综合收益"`                //七、其他综合收益
	F039N float64 `excel:"八、综合收益总额"`                //八、综合收益总额
	F040N float64 `excel:"其中：归属于母公司"`               //其中：归属于母公司
	F041N float64 `excel:"其中：归属于少数股东"`              //其中：归属于少数股东
	MEMO  string  `excel:"备注"`                      //备注
	F062N float64 `excel:"其中：利息费用"`                 //其中：利息费用
	F063N float64 `excel:"其中：利息收入"`                 //其中：利息收入
	F064N float64 `excel:"信用减值损失（2019格式）"`          //信用减值损失（2019格式）
	F065N float64 `excel:"资产减值损失（2019格式）"`          //资产减值损失（2019格式）
	F066N float64 `excel:"其中：归属母公司所有者的其他综合收益的税后净额"` //其中：归属母公司所有者的其他综合收益的税后净额
	F067N float64 `excel:"其中：归属于少数股东的其他综合收益的税后净额"`  //其中：归属于少数股东的其他综合收益的税后净额
}
