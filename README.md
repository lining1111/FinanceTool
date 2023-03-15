#利用go程序来记录一些金融小常识
    纲领思想：从国家规定，结合各大财务评估系统的财务报表内容及指标出发，
        财务报表的个数是一定的:3张.
        财务报表个表的元素是一定的:此处有国家规定和指导文件，硬性要求.
        财务报表的指标内容是一定的:这里讲的指标内容为通过财务报表各表元素经过基本算术运算(+-*/)得出的初级财务指标。
                            具体内容选取应该结合国家规定和流行财务分析系统来制定。
        财务涉及金钱，所以个人的认知和判断，应该首先基于已经存在且不随环境变化的数值，
            1、这里通过干净的财务报表数据(即三表元素)，通过算术得到的初级指标，是可信的，100%
            2、然后根据国家规定及流行财务分析系统，进行二次加工，可信度降低。原因为：市场变化因素很多，往往成熟的分析师，
                都是事后诸葛亮行为，但历史因素具有周期性，所以从分析系统和分析师经验角度来剖析新的或者未知的财务情况，具有很高的参考性
            3、具体行业内的财务分析，应当结合行业特性，从供需、政策、产品工艺、从业人员水平等方面进行综合分析，筛出干扰信息。
            4、分析原则，应该保持平常心，对具体公司的分析结果应该是在本行业内的正态分布曲线内，过分好和过分差，说明风险过高，都应该引起重要的警惕。
##目的
    利用go的快捷性，以下列出：
    1 网络接口(获取数据，web接口通过网页展示数据)
    2 数据库接口(将网络数据存入数据库，方便进行历史数据的分析)
    3 文件接口(将网络数据或者本地数据库数据存到文件，供一般用户获取)
    4 便捷的序列化和反序列化操作(外部接口尽量采用json这种易用易读的方式来进行扩展)
##文件夹分布
    基本每个文件夹都会分为文本部分和程序部分：
    文本部分负责将书本知识转化到本地易读易懂话语，人类阅读
    程序部分负责将数据和课本数据进行展示和对比等操作，完成信息总结类作用，机器阅读，其实也是通过机器完成数据分析，再供人类使用
    -COM 通信及文件、数据库操作接口
        --cninfo 巨潮资讯数据访问接口
    按 债券、股票、基金 分类处理
    -Bond 债券
        --BaseInfo 基本信息
        --Market 市场信息

    -Stock 股票
        --HSB 沪深北
            ---BalanceSheet 资产负债表
            ---IncomeSheet 利润表
            ---CashFlowSheet 现金流量表
            ---FinanicalRatios 财务指标
            ---AuditOpinion 审计意见
            ---IMR 行业指标排名
            ---Market 股票相关信息
        --XSB 新三板
            ---BalanceSheet 资产负债表
            ---IncomeSheet 利润表
            ---CashFlowSheet 现金流量表
            ---FinanicalRatios 财务指标
            ---AuditOpinion 审计意见
            ---Market 股票相关信息
        --HK 港股
            ---BalanceSheet 资产负债表
            ---BalanceSheetBank 资产负债表银行
            ---IncomeSheet 利润表
            ---IncomeSheetBank 利润表银行
            ---CashFlowSheet 现金流量表
            ---FinanicalRatios 财务指标
            ---Market 股票相关信息
        --BSE 北交所
            ---BalanceSheet 资产负债表
            ---IncomeSheet 利润表
            ---CashFlowSheet 现金流量表
            ---FinanicalRatios 财务指标
            ---AuditOpinion 审计意见
            ---IMR 行业指标排名
            ---Market 股票相关信息
    -Fund 基金
        --BalanceSheet 基金资产负债表2009版
        --Benchmark 基金经营业绩及收益分配表2009版
        --FinancialRatios 基金主要财务指标2009版

##企业会计准则具体细节
    企业会计准则
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/index.htm
    企业会计准则--基本准则
    http://www.gov.cn/govtest/content_270047.htm
    企业会计准则第1号——存货 
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46216.htm
    企业会计准则第2号——长期股权投资
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46215.htm
    企业会计准则第3号——投资性房地产
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46214.htm
    企业会计准则第4号——固定资产
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46213.htm
    企业会计准则第5号——生物资产
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46212.htm
    企业会计准则第6号——无形资产
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46242.htm
    企业会计准则第7号——非货币性资产交换（财会[2019]8号）
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/201909/t20190911_3384679.htm
    企业会计准则第8号——资产减值
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46240.htm
    企业会计准则第9号--职工薪酬
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46239.htm
    企业会计准则第10号--企业年金基金
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46238.htm
    企业会计准则第11号--股份支付
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46237.htm
    企业会计准则第12号——债务重组
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/201910/t20191028_3410789.htm
    企业会计准则第13号--或有事项
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46235.htm
    企业会计准则第14号——收入（财会[2017]22号）
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/201709/t20170907_2694006.htm
    企业会计准则第16号——政府补助（财会〔2017〕15号）
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46232.htm
    企业会计准则第17号——借款费用
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46231.htm
    企业会计准则第18号——所得税
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46230.htm
    企业会计准则第19号--外币折算
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46229.htm
    企业会计准则第20号——企业合并
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46228.htm
    企业会计准则第21号——租赁（财会〔2018〕35号）
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/201910/t20191028_3411190.htm
    企业会计准则第22号——金融工具确认和计量（财会[2017]7号）
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/201709/t20170908_2694655.htm
    企业会计准则第23号——金融资产转移（财会[2017]8号）
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/201709/t20170908_2694626.htm
    企业会计准则第24号——套期会计（财会[2017]9号）
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/201709/t20170908_2694624.htm
    企业会计准则第25号--原保险合同
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46223.htm
    企业会计准则第26号--再保险合同
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46222.htm
    企业会计准则第27号——石油天然气开采
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46221.htm
    企业会计准则第28号——会计政策、会计估计变更和差错更正
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46220.htm
    企业会计准则第29号——资产负债表日后事项
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46219.htm
    企业会计准则第30号——财务报表列报
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46218.htm
    企业会计准则第31号——现金流量表
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46250.htm
    企业会计准则第32号——中期财务报告
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46249.htm
    企业会计准则第33号——合并财务报表
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46248.htm
    企业会计准则第34号——每股收益
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46247.htm
    企业会计准则第35号——分部报告
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46246.htm
    企业会计准则第36号——关联方披露
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46245.htm
    企业会计准则第37号——金融工具列报（财会[2017]14号）
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/201709/t20170907_2694118.htm
    企业会计准则第37号--金融工具列报（财会[2014]23号）
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46244.htm
    企业会计准则第38号——首次执行企业会计准则
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/200806/t20080618_46243.htm
    企业会计准则第39号--公允价值计量
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/201512/t20151208_1602631.htm
    企业会计准则第40号——合营安排
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/201512/t20151208_1602633.htm
    企业会计准则第41号——在其他主体中权益的披露
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/201512/t20151208_1602637.htm
    企业会计准则第42号——持有待售的非流动资产、处置组和终止经营（财会〔2017〕13号）
    http://kjs.mof.gov.cn/zt/kjzzss/kuaijizhunzeshishi/201709/t20170907_2694145.htm
##财务报告数据获取方式
    各上市公司官网，此方法较为分散，不建议使用
    上海证券交所或者深圳证券交易所的官网 
    1.上海证券交所--披露 http://www.sse.com.cn/disclosure/listedinfo/regular/
    2.深圳证券交易所--信息披露--上市公司信息--定期报告  http://www.szse.cn/disclosure/listed/fixed/index.html
    3.中国证券网--信息披露平台 https://xinpi.cnstock.com/
        沪市 https://xinpi.cnstock.com/Search.aspx?Style=012001
        深市 https://xinpi.cnstock.com/Search.aspx?Style=012002
        中国证券网的web api感觉不太好总结，所以就先舍弃
    本所内的各个公司的财务报表格式是一致的
    他们的报告内容都是pdf格式的，在信息拆解的时候，需要进行pdf文件操作
    4.中国证监会指定披露平台 巨潮资讯 http://webapi.cninfo.com.cn/#/apiDoc
    
    ------本项目采取的方式是巨潮资讯上公开api接口，获取信息。
    数据中心 http://webapi.cninfo.com.cn/#/dataBrowse 选财务指标
    下面有细分项 报告期---财务指标---单季度---TTM
    再细分 报告期---个股报告期资产负债表---个股报告期利润表---个股报告期现金表
          财务指标---个股报告期指标表---财务指标行业排名
            单季度---个股单季财务利润表---个股单季现金流量表---个股单季财务指标

    FSMeasures结构体内容就是根据 财务指标---个股报告期指标表内容来做的
    http://webapi.cninfo.com.cn/api/stock/p_stock2303?
##20230309
    俗话说：凡人畏果，圣人畏因。
    投资本质就是投资人与融资人的一场博弈。
    博弈的取胜就是基于尽可能全的正确信息完成正确的决策
    这里分为3步走：
        1、学会分解信息，从信息本身的产生出发，回溯原因，将信息分为真实信息和虚假信息两类，分别处理
        2、真实信息处理---从真实信息对过去行为产生的风险和收益作为判断，进而影响未来投资行为
        3、虚假信息处理---分析历史，看虚假信息与哪些异常历史数据有关，此时要更为谨慎的分析融资人的行为意识，及时止损。
##20230311
    记录心得，分为两个部分记录
    --Accounting 会计心得
    --Financial 金融心得
    代码的作用是从金融数据库获取信息，落表(excel) 分析或者 程序内分析体系 分析。
        代码分三部分 
            获取数据库
            落表
            分析体系
##20230314
    粗略学会会计学后的，分析方法总结：
        A1、会计学主要是讲报表内各个元素的填写方法，即 金融行为-->报表元素变化
            设：报表元素为YJitem,YDitem(YJitem为借方数组、YDitem为贷方数组);金融行为为Xaction(Xation为单元素)
            根据借贷记账法得出 ：YJitem=j(Xation)
                             YDitem=d(Xation)
                             Sum(YJitem)=Sum(YDitem)
            它们为一个方程组，且为一元一次的方程组，这个从书上的例子得到一个这样的事实
                正向---金融行为是一个应用题的题目，而这个题目一定会出现一个特定的一元一次方程组。
                        这个是考生视角，从有特定的财会软件得到可信度100%
                反向推理---这个是出题人视角，
                            1、从方程组得到题目，会得到有限集合的金融行为，因为金融行为是要有会计凭证供审计的，可信度100%
                            2、然后从公司管理层的意图分解成若干组(恶意、平意、善意),这里的善意个人应该很少或者没有(自私角度，个人只对个人的利益负责)
                            3、根据各个组合结合报表中的说明和实际的外部环境，分别赋予可信值。
                            4、这些可信值会再次进入完整的评价决策体系来完成输出。
        A2、各个阶段的报表内容，从审计意见可以依次得出计算正常及名面证据关联度的100%可信。
            设：报表的指标Ytarget(Ytarget为单一指标),报表的元素Xitem(Xitem为数组),
            根据会计准则得出：Ytarget=T(Xitem)
            它是一个多元一次方程组。
            这里可以全部取可信度为100%的指标，因为根据披露的审计意见和公司报表的目的(求得更高的投资),
            它们最终会趋于100%(这里指的是报表相关佐证的到指标的计算正确) 
        A3、若干个指标组成一个实际的要求(比如股东的基本收益保证为：净资产报酬率=总资产周转率*销售净利率*权益乘数)
            设：一个明确的要求Yget(单个要求)，报表指标Xtarget(Xtarget为数组)
            根据特定的评价体系得出Yget=U(Xtarget)
            它是一个多元不定次的方程，其中系数的确定、指标的选取就是体系的关键。
        A4、一个特定要求的投资目的，可以分解成若干实际的要求。
            设：一个投资目的Yattr(具体的要求，比如长期收益率)，若干要求Xget(Xget为数组)
            得到Yattr=U(Xget)
            它是一个多元一次方程，其中系数的确定、要求的选取也是关键。
        其实上面的A3、A4可以合并为一个就是Yattr=U(Xtarget)
    综上所述、投资初期和后续执行按步骤应该是
            投资初期：
                周期以季度为单位，以获取历史信息为主
                1、设定一个明确的目标，可控且可接受本金损失的风险。这里需要确定以下几个量：
                    1）投资类型(这里类型分为期限+所投领域，短期投资的评价体系和长期的评价不同，
                                领域取决于知识储备对报表后的实际金融行为认知程度，即不要试图赚取认知范围外的钱)
                    2）投资的本金(以固定收入-固定支出-抵御风险的备用金)
                    3）投资回报率(这里的回报率要细分为时间+利息，因为利率是带时间属性的乘法因子。确定的时间会带来机会成本)
                2、根据1的描述、经过A4分解成若干小目的。一维数组。(目标)
                3、根据2的结果，经过A3将每个小目标根据评价体系分解成若干。二维数组(目标、指标)
                4、根据3的结果，通过比较分析法，选出近期(比如一年内)财务报表的表现符合要求的预投资对象。一维数组(股票代码)
                5、根据4的结果，通过因素分析法，对数组内的代码进行长历史记录(比如从企业上市开始时间，一直到现在)。得到三维数组(股票代码，指标、历史记录)
                6、根据5的结果，以股票代码为单元，将各个指标的历史曲线绘图。做单个指标分析。得到三维数组(股票代码，指标，指标的历史韧性)
                7、根据6的结果，以股票代码为单元，经过评价体系加权，作出单个股票的综合分析，得到二维数组(股票代码，分析结果)
                8、根据7的结果，通过分析结果对比，选出股票组合(这里一定要是有风险对冲效果的组合，即：一方跌必然或不一定立即引起另一方的涨)
                风险对冲的方式，有点类似通信中用双绞线和差分线来应对信号噪声一样。
            投资后续：
                周期以天为单位，长期投资战略的话，以周加权为单位，以获取反馈信息为主
                1、根据选定的股票代码，定时获取该股票的实时行情，通过它的行情来对比同行业的行情，作出决策的反馈信息。记录。一维数组(股票代码的反馈信息)
                2、根据后续的季度性报告，对比1的反馈信息记录，来质证报告内容。得到报告内容的可信度。一维数组(股票代码的指标可信度)
                3、根据2的指标可信度，按从低到高(个人对个人的利益负责),按照会计准则的到与其相关的若干元素组合。三维维数组(指标可信度、元素组合数组、元素组合)
                4、根据3的结果，以指标为单元，完成(元素组合)与 报表披露内容及1中反馈信息记录、还有大环境因子的对应关系。四维数组(指标可信度、元素组合数组、元素组合、[反馈信息+大环境因子说明])
                5、综合4的结果，按照指标可信度，从评价体系中计算当前股票的收益。二维数组(股票代码的现时收益预测)
                6、根据5的结果，结合投资周期的战略，作出调整或者不调整投资组合的决定。

        综合上面的步骤，输入量为[投资周期、本金、利息、投资对象类型[]]。输出量[股票代码(投资金额、投资周期、预期利息)]
        投资初期步骤
        1、根据定义的总要求分解成若干小目标
            func GetAims(attr string) (aim []string, err error)
        2、将每一个小目标分解成若干实际的财务指标,需要输入特点的评价系统名称
            func GetRatios(aim string,eval string)(ratio []float64,err error)
        3、根据各个指标的要求，根据比较分析法，选出当期符合要求的股票代码集合
            func GetScodes(ratio float64) (scode []string, err error)
        4、根据单个股票，获取指标集合的历史数据
                根据股票和单个指标获得历史数据集合 这里提出一个数据单元结构体，就是 历史指标值(数值和时间)
                func GetSingleRatioHistory(scode string,ratio string) (ratioHistory []DataUnit,err error)
                将单个指标的历史数据做成一个集合 这里提出一个数据单元集合结构体,就是 指标历史(指标名称+历史指标值数组)
            func GetRatiosHistory(scode string, ratios []string) (ratiosHistory []RatioHistory, err error)
        5、根据各个指标的历史数据，通过统计方法，作出分析
            