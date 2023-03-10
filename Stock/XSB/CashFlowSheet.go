package XSB

/**
现金流量表，数据类计算
*/

const APICS = "http://webapi.cninfo.com.cn/api/neeq/p_neeq6020"

//CashFlowSheet 现金流量表 "http://webapi.cninfo.com.cn/api/neeq/p_neeq6020"
//params:	scode	股票代码	string	输入不超过50只股票代码，用逗号分隔；输入多个代码时，不允许报告期为空.
//			rdate	报告期	string	报告期 可为空，为空取所有报告期
//			type	合并类型	string	通过p_public0006可获取，对应的总类编码为‘071’
type CashFlowSheet struct {
	ORGNAME     string  `excel:"机构名称"`                         //机构名称	varchar
	SECCODE     string  `excel:"证券代码"`                         //证券代码	varchar
	SECNAME     string  `excel:"证券简称"`                         //证券简称	varchar
	DECLAREDATE string  `excel:"公告日期"`                         //公告日期	date
	STARTDATE   string  `excel:"开始日期"`                         //开始日期	date
	ENDDATE     string  `excel:"截止日期"`                         //截止日期	date
	F001D       string  `excel:"报告年度"`                         //报告年度	date
	F002V       string  `excel:"合并类型编码"`                       //合并类型编码	varchar		071001合并本期071002合并上期071003母公司本期071004母公司上期
	F003V       string  `excel:"合并类型"`                         //合并类型	varchar		详见合并类型编码备注
	F004V       string  `excel:"报表来源编码"`                       //报表来源编码	varchar		033003定期报告033006转让说明书033013其他
	F005V       string  `excel:"报表来源"`                         //报表来源	varchar		详见报表来源编码备注
	F006N       float64 `excel:"销售商品、提供劳务收到的现金"`               //销售商品、提供劳务收到的现金	decimal		单位：元
	F072N       float64 `excel:"客户存款和同业存放款项净增加额"`              //客户存款和同业存放款项净增加额	decimal		单位：元
	F073N       float64 `excel:"向中央银行借款净增加额"`                  //向中央银行借款净增加额	decimal		单位：元
	F074N       float64 `excel:"向其他金融机构拆入资金净增加额"`              //向其他金融机构拆入资金净增加额	decimal		单位：元
	F077N       float64 `excel:"收到原保险合同保费取得的现金"`               //收到原保险合同保费取得的现金	decimal		单位：元
	F078N       float64 `excel:"收到再保险业务现金净额"`                  //收到再保险业务现金净额	decimal		单位：元
	F079N       float64 `excel:"保户储金及投资款净增加额"`                 //保户储金及投资款净增加额	decimal		单位：元
	F080N       float64 `excel:"处置以公允价值计量且其变动计入当期损益的金融资产净增加额"` //处置以公允价值计量且其变动计入当期损益的金融资产净增加额	decimal		单位：元
	F081N       float64 `excel:"收取利息、手续费及佣金的现金"`               //收取利息、手续费及佣金的现金	decimal		单位：元
	F082N       float64 `excel:"拆入资金净增加额"`                     //拆入资金净增加额	decimal		单位：元
	F083N       float64 `excel:"回购业务资金净增加额"`                   //回购业务资金净增加额	decimal		单位：元
	F007N       float64 `excel:"收到的税费返还"`                      //收到的税费返还	decimal		单位：元
	F008N       float64 `excel:"收到其他与经营活动有关的现金"`               //收到其他与经营活动有关的现金	decimal		单位：元
	F009N       float64 `excel:"经营活动现金流入小计"`                   //经营活动现金流入小计	decimal		单位：元
	F010N       float64 `excel:"购买商品、接受劳务支付的现金"`               //购买商品、接受劳务支付的现金	decimal		单位：元
	F084N       float64 `excel:"客户贷款及垫款净增加额"`                  //客户贷款及垫款净增加额	decimal		单位：元
	F085N       float64 `excel:"存放中央银行和同业款项净增加额"`              //存放中央银行和同业款项净增加额	decimal		单位：元
	F086N       float64 `excel:"支付原保险合同赔付款项的现金"`               //支付原保险合同赔付款项的现金	decimal		单位：元
	F087N       float64 `excel:"支付利息、手续费及佣金的现金"`               //支付利息、手续费及佣金的现金	decimal		单位：元
	F088N       float64 `excel:"支付保单红利的现金"`                    //支付保单红利的现金	decimal		单位：元
	F011N       float64 `excel:"支付给职工以及为职工支付的现金"`              //支付给职工以及为职工支付的现金	decimal		单位：元
	F012N       float64 `excel:"支付的各项税费"`                      //支付的各项税费	decimal		单位：元
	F013N       float64 `excel:"支付其他与经营活动有关的现金"`               //支付其他与经营活动有关的现金	decimal		单位：元
	F014N       float64 `excel:"经营活动现金流出小计"`                   //经营活动现金流出小计	decimal		单位：元
	F015N       float64 `excel:"经营活动产生的现金流量净额"`                //经营活动产生的现金流量净额	decimal		单位：元
	F016N       float64 `excel:"收回投资收到的现金"`                    //收回投资收到的现金	decimal		单位：元
	F017N       float64 `excel:"取得投资收益收到的现金"`                  //取得投资收益收到的现金	decimal		单位：元
	F018N       float64 `excel:"处置固定资产、无形资产和其他长期资产收回的现金净额"`    //处置固定资产、无形资产和其他长期资产收回的现金净额	decimal		单位：元
	F019N       float64 `excel:"处置子公司及其他营业单位收到的现金净额"`          //处置子公司及其他营业单位收到的现金净额	decimal		单位：元
	F020N       float64 `excel:"收到其他与投资活动有关的现金"`               //收到其他与投资活动有关的现金	decimal		单位：元
	F021N       float64 `excel:"投资活动现金流入小计"`                   //投资活动现金流入小计	decimal		单位：元
	F022N       float64 `excel:"购建固定资产、无形资产和其他长期资产支付的现金"`      //购建固定资产、无形资产和其他长期资产支付的现金	decimal		单位：元
	F023N       float64 `excel:"投资支付的现金"`                      //投资支付的现金	decimal		单位：元
	F075N       float64 `excel:"质押贷款净增加额"`                     //质押贷款净增加额	decimal		单位：元
	F024N       float64 `excel:"取得子公司及其他营业单位支付的现金净额"`          //取得子公司及其他营业单位支付的现金净额	decimal		单位：元
	F025N       float64 `excel:"支付其他与投资活动有关的现金"`               //支付其他与投资活动有关的现金	decimal		单位：元
	F026N       float64 `excel:"投资活动现金流出小计"`                   //投资活动现金流出小计	decimal		单位：元
	F027N       float64 `excel:"投资活动产生的现金流量净额"`                //投资活动产生的现金流量净额	decimal		单位：元
	F028N       float64 `excel:"吸收投资收到的现金"`                    //吸收投资收到的现金	decimal		单位：元
	F089N       float64 `excel:"其中：子公司吸收少数股东投资收到的现金"`          //其中：子公司吸收少数股东投资收到的现金	decimal		单位：元
	F029N       float64 `excel:"取得借款收到的现金"`                    //取得借款收到的现金	decimal		单位：元
	F076N       float64 `excel:"发行债券收到的现金"`                    //发行债券收到的现金	decimal		单位：元
	F030N       float64 `excel:"收到其他与筹资活动有关的现金"`               //收到其他与筹资活动有关的现金	decimal		单位：元
	F031N       float64 `excel:"筹资活动现金流入小计"`                   //筹资活动现金流入小计	decimal		单位：元
	F032N       float64 `excel:"偿还债务支付的现金"`                    //偿还债务支付的现金	decimal		单位：元
	F033N       float64 `excel:"分配股利、利润或偿付利息支付的现金"`            //分配股利、利润或偿付利息支付的现金	decimal		单位：元
	F090N       float64 `excel:"其中：子公司支付给少数股东的股利、利润"`          //其中：子公司支付给少数股东的股利、利润	decimal		单位：元
	F034N       float64 `excel:"支付其他与筹资活动有关的现金"`               //支付其他与筹资活动有关的现金	decimal		单位：元
	F035N       float64 `excel:"筹资活动现金流出小计"`                   //筹资活动现金流出小计	decimal		单位：元
	F036N       float64 `excel:"筹资活动产生的现金流量净额"`                //筹资活动产生的现金流量净额	decimal		单位：元
	F037N       float64 `excel:"四、汇率变动对现金的影响"`                 //四、汇率变动对现金的影响	decimal		单位：元
	F038N       float64 `excel:"四(2)、其他原因对现金的影响"`              //四(2)、其他原因对现金的影响	decimal		单位：元
	F039N       float64 `excel:"五、现金及现金等价物净增加额"`               //五、现金及现金等价物净增加额	decimal		单位：元
	F040N       float64 `excel:"期初现金及现金等价物余额"`                 //期初现金及现金等价物余额	decimal		单位：元
	F041N       float64 `excel:"期末现金及现金等价物余额"`                 //期末现金及现金等价物余额	decimal		单位：元
	F042N       float64 `excel:"附注"`                           //附注：	decimal		单位：元
	F043N       float64 `excel:"1、将净利润调节为经营活动现金流量"`            //1、将净利润调节为经营活动现金流量：	decimal		单位：元
	F044N       float64 `excel:"净利润"`                          //净利润	decimal		单位：元
	F045N       float64 `excel:"加：资产减值准备"`                     //加：资产减值准备	decimal		单位：元
	F046N       float64 `excel:"固定资产折旧、油气资产折耗、生产性生物资产折旧"`      //固定资产折旧、油气资产折耗、生产性生物资产折旧	decimal		单位：元
	F091N       float64 `excel:"投资性房地产的折旧及摊销"`                 //投资性房地产的折旧及摊销	decimal		单位：元
	F047N       float64 `excel:"无形资产摊销"`                       //无形资产摊销	decimal		单位：元
	F048N       float64 `excel:"长期待摊费用摊销"`                     //长期待摊费用摊销	decimal		单位：元
	F049N       float64 `excel:"处置固定资产、无形资产和其他长期资产的损失"`        //处置固定资产、无形资产和其他长期资产的损失	decimal		单位：元
	F050N       float64 `excel:"固定资产报废损失"`                     //固定资产报废损失	decimal		单位：元
	F051N       float64 `excel:"公允价值变动损失"`                     //公允价值变动损失	decimal		单位：元
	F052N       float64 `excel:"财务费用"`                         //财务费用	decimal		单位：元
	F053N       float64 `excel:"投资损失"`                         //投资损失	decimal		单位：元
	F054N       float64 `excel:"递延所得税资产减少"`                    //递延所得税资产减少	decimal		单位：元
	F055N       float64 `excel:"递延所得税负债增加"`                    //递延所得税负债增加	decimal		单位：元
	F056N       float64 `excel:"存货的减少"`                        //存货的减少	decimal		单位：元
	F057N       float64 `excel:"经营性应收项目的减少"`                   //经营性应收项目的减少	decimal		单位：元
	F058N       float64 `excel:"经营性应付项目的增加"`                   //经营性应付项目的增加	decimal		单位：元
	F059N       float64 `excel:"其他"`                           //其他	decimal		单位：元
	F060N       float64 `excel:"经营活动产生的现金流量净额（附注）"`            //经营活动产生的现金流量净额（附注）	decimal		单位：元
	F061N       float64 `excel:"2、不涉及现金收支的重大投资和筹资活动"`          //2、不涉及现金收支的重大投资和筹资活动：	decimal		单位：元
	F062N       float64 `excel:"债务转为资本"`                       //债务转为资本	decimal		单位：元
	F063N       float64 `excel:"一年内到期的可转换公司债券"`                //一年内到期的可转换公司债券	decimal		单位：元
	F064N       float64 `excel:"融资租入固定资产"`                     //融资租入固定资产	decimal		单位：元
	F065N       float64 `excel:"3、现金及现金等价物净变动情况"`              //3、现金及现金等价物净变动情况：	decimal		单位：元
	F066N       float64 `excel:"现金的期末余额"`                      //现金的期末余额	decimal		单位：元
	F067N       float64 `excel:"减：现金的期初余额"`                    //减：现金的期初余额	decimal		单位：元
	F068N       float64 `excel:"加：现金等价物的期末余额"`                 //加：现金等价物的期末余额	decimal		单位：元
	F069N       float64 `excel:"减：现金等价物的期初余额"`                 //减：现金等价物的期初余额	decimal		单位：元
	F070N       float64 `excel:"加：其他原因对现金的影响2"`                //加：其他原因对现金的影响2	decimal		单位：元
	F071N       float64 `excel:"现金及现金等价物净增加额2"`                //现金及现金等价物净增加额2	decimal		单位：元
	MEMO        string  `excel:"备注"`                           //备注	varchar
	F092N       float64 `excel:"为交易目的而持有金融资产净增加额"`             //为交易目的而持有金融资产净增加额	decimal(18,2)	元	20210826新增
	F093N       float64 `excel:"拆出资金净增加额"`                     //拆出资金净增加额	decimal(18,2)	元	20210826新增
	F094N       float64 `excel:"信用减值损失"`                       //信用减值损失	decimal(18,2)	元	20210826新增
	F095N       float64 `excel:"少数股东损益"`                       //少数股东损益	decimal(18,2)	元	20210826新增
	F096N       float64 `excel:"待摊费用的减少"`                      //待摊费用的减少	decimal(18,2)	元	20210826新增
	F097N       float64 `excel:"预提费用的增加"`                      //预提费用的增加	decimal(18,2)	元	20210826新增
	F098N       float64 `excel:"递延收益摊销"`                       //递延收益摊销	decimal(18,2)	元	20210826新增
	F099N       float64 `excel:"预计负债的增加"`                      //预计负债的增加	decimal(18,2)	元	20210826新增
}
