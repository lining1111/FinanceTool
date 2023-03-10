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