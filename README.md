#利用go程序来记录一些金融小常识
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
    FinancialStatement  财务报表三张表
    FinancialStatement/FSConcreteDetails 财务报表三张表内根据各行各业具体分类的细节
##20230227
    想想还是挺兴奋的，经济基础决定上层建筑，金融是一个复杂的学科，涉及数学、博弈、政策和无处不在的谎言。
    用数据分析的方式，分析金融数据背后的故事，应该很有意思。初步想法，先从会计学的三个报表开始，完成数据存储以及web上的展示(包括表格和柱状图)
    这尼玛，涉及前端就不会了，尬！！！

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
    

