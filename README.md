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
##20230227
    想想还是挺兴奋的，经济基础决定上层建筑，金融是一个复杂的学科，涉及数学、博弈、政策和无处不在的谎言。
    用数据分析的方式，分析金融数据背后的故事，应该很有意思。初步想法，先从会计学的三个报表开始，完成数据存储以及web上的展示(包括表格和柱状图)
    这尼玛，涉及前端就不会了，尬！！！