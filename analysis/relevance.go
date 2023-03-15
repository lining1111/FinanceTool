package analysis

import "fmt"

//Element 指标与元素的关联关系
type Element struct {
	Name    string  //会计科目名称
	Account float64 //金额
}

type Relevance struct {
	Action  string    //行为描述
	Account float64   //金额
	Dr      []Element //借方
	Cr      []Element //贷方
}

//Check 查看这次行为后的正确性，即借贷双方的和相等
func (r Relevance) Check() bool {
	sumDr := 0.0
	sumCr := 0.0
	for _, v := range r.Dr {
		sumDr += v.Account
	}
	for _, v := range r.Cr {
		sumCr += v.Account
	}
	if sumDr != sumCr {
		fmt.Println("借方和 ", sumDr, "贷方合 ", sumCr)
		return false
	}

	return true
}

//下面的金融行为模板，是根据会计学ISBN：978-7-300-27234-4内容编写的。
//在实际中，这块的内容是固定不变的，金额只存在公司的会计系统内，是不会通过财务报告来向外透露的，
//但是可以反向推理从异常指标或者异常的报表元素来找其相关的元素变化情况，起到预警。
//=====对比能够发现问题的原因，实际的金融活动都有对应的元素变动+会计凭证。这些也是稽查的东西。企业造假的过程就是在组合这些金融行为后，躲过稽查，完成不可告人的目的

//TemplateFinancial 金融行为模板
var TemplateFinancial = []Relevance{
	//1 现金存入银行
	{Action: "把现金存入银行",
		Dr: []Element{
			{Name: "银行存款"},
		},
		Cr: []Element{
			{Name: "库存现金"},
		},
	},
	//2 签发支票，从银行存款账户取走
	{Action: "签发支票，从银行存款账户取走",
		Dr: []Element{
			{Name: "库存现金"},
		},
		Cr: []Element{
			{Name: "银行存款"},
		},
	},
	//3 出差借支
	{Action: "出差借支",
		Dr: []Element{
			{Name: "其他应收款"},
		},
		Cr: []Element{
			{Name: "库存现金"},
		},
	},
	//4 出差回来报销15000 交回现金6000
	{Action: "出差报销",
		Dr: []Element{
			{Name: "库存现金"},
			{Name: "管理费用"},
		},
		Cr: []Element{
			{Name: "其他应收款"},
		},
	},
	//5 盘点库存现金，短缺 9000
	{Action: "盘点库存现金，短缺",
		Dr: []Element{
			{Name: "待处理财产损益-待处理流动资产损益"},
		},
		Cr: []Element{
			{Name: "库存现金"},
		},
	},
	//5 追回员工借款 3000
	{Action: "追回员工借款",
		Dr: []Element{
			{Name: "其他应收款"},
			{Name: "管理费用"},
		},
		Cr: []Element{
			{Name: "待处理财产损益-待处理流动资产损益"},
		},
	},
	//6 现金溢余 2000
	{Action: "现金溢余",
		Dr: []Element{
			{Name: "库存现金"},
		},
		Cr: []Element{
			{Name: "待处理财产损益-待处理流动资产损益"},
		},
	},
	//7 计入营业外收入 2000
	{Action: "计入营业外收入",
		Dr: []Element{
			{Name: "待处理财产损益-待处理流动资产损益"},
		},
		Cr: []Element{
			{Name: "营业外收入"},
		},
	},
	//8 银行存入
	{Action: "银行存入",
		Dr: []Element{
			{Name: "其他货币资金"},
		},
		Cr: []Element{
			{Name: "银行存款"},
		},
	},
	//9 通过支票同于开支
	{Action: "通过支票同于开支",
		Dr: []Element{
			{Name: "在途物资、库存商品"},
		},
		Cr: []Element{
			{Name: "其他货币资金"},
		},
	},
	//10 用银行存款支付采购费用
	{Action: "用银行存款支付采购费用",
		Dr: []Element{
			{Name: "在途物资、库存商品"},
			{Name: "应交税费"},
		},
		Cr: []Element{
			{Name: "其他货币资金"},
		},
	},
	//11 用银行存款支付采购费用
	{Action: "用银行存款支付采购费用",
		Dr: []Element{
			{Name: "在途物资、库存商品"},
			{Name: "应交税费"},
		},
		Cr: []Element{
			{Name: "其他货币资金"},
		},
	},
	//12 用银行存款支付业务招待费
	{Action: "用银行存款支付业务招待费",
		Dr: []Element{
			{Name: "管理费用"},
		},
		Cr: []Element{
			{Name: "其他货币资金"},
		},
	},
	//13 收到银行存款的利息
	{Action: "收到银行存款的利息",
		Dr: []Element{
			{Name: "其他货币资金"},
		},
		Cr: []Element{
			{Name: "财务费用"},
		},
	},
	//14 向证券公司划入资金，准备投资
	{Action: "向证券公司划入资金，准备投资",
		Dr: []Element{
			{Name: "其他货币资金"},
		},
		Cr: []Element{
			{Name: "银行存款"},
		},
	},
	//15 用准备投资的钱，买了一支股票
	{Action: "用准备投资的钱，买了一支股票",
		Dr: []Element{
			{Name: "交易性金融资产"},
		},
		Cr: []Element{
			{Name: "其他货币资金"},
		},
	},
	//16 出售商品，缴税并代垫付运费 ====这里一个知识点，就是金融活动发生时，必有税费产生，即不同时期，主营业务收入和税费的比率应该相差不大
	{Action: "出售商品，缴税并代垫付运费",
		Dr: []Element{
			{Name: "应收账款"},
		},
		Cr: []Element{
			{Name: "主营业务收入"},
			{Name: "应交税费"},
			{Name: "库存现金"},
		},
	},
	//17 出售商品，缴税，收到不带息商业承兑汇票
	{Action: "出售商品，缴税，收到不带息商业承兑汇票",
		Dr: []Element{
			{Name: "应收票据"},
		},
		Cr: []Element{
			{Name: "主营业务收入"},
			{Name: "应交税费"},
		},
	},
	//18 购买商品，预付账款
	{Action: "购买商品，预付账款",
		Dr: []Element{
			{Name: "预付账款"},
		},
		Cr: []Element{
			{Name: "银行存款"},
		},
	},
	//19 采购
	{Action: "采购",
		Dr: []Element{
			{Name: "库存商品"},
			{Name: "应交税费"},
		},
		Cr: []Element{
			{Name: "银行存款"},
		},
	},
	//20 销售
	{Action: "销售",
		Dr: []Element{
			{Name: "银行存款"},
		},
		Cr: []Element{
			{Name: "主营业务收入"},
			{Name: "应交税费"},
		},
	},
	//21 结转销货成本
	{Action: "结转销货成本",
		Dr: []Element{
			{Name: "主营业务收入"},
		},
		Cr: []Element{
			{Name: "库存商品"},
		},
	},
	//21 缴纳税费
	{Action: "缴纳税费",
		Dr: []Element{
			{Name: "应交税费"},
		},
		Cr: []Element{
			{Name: "银行存款"},
		},
	},
	//22 采购原材料
	{Action: "采购原材料",
		Dr: []Element{
			{Name: "在途物资"},
			{Name: "应交税费"},
		},
		Cr: []Element{
			{Name: "银行存款"},
		},
	},
	//22 采购原材料
	{Action: "采购原材料",
		Dr: []Element{
			{Name: "在途物资"},
			{Name: "应交税费"},
		},
		Cr: []Element{
			{Name: "银行存款"},
		},
	},
	//23 原材料入库
	{Action: "原材料入库",
		Dr: []Element{
			{Name: "原材料"},
		},
		Cr: []Element{
			{Name: "在途物资"},
		},
	},
	//24 领取原材料进行生产
	{Action: "领取原材料进行生产",
		Dr: []Element{
			{Name: "生产成本"},
		},
		Cr: []Element{
			{Name: "原材料"},
		},
	},
	//25 计发工人工资
	{Action: "计发工人工资",
		Dr: []Element{
			{Name: "生产成本"},
		},
		Cr: []Element{
			{Name: "应付职工薪资"},
		},
	},
	//26 实发工人工资
	{Action: "实发工人工资",
		Dr: []Element{
			{Name: "应付职工薪资"},
		},
		Cr: []Element{
			{Name: "银行存款"},
		},
	},
	//27 计算生产中间费用，包括水电气费用和车间主任薪酬
	{Action: "计算生产中间费用，包括水电气费用和车间主任薪酬",
		Dr: []Element{
			{Name: "制造费用"},
		},
		Cr: []Element{
			{Name: "应付职工薪资"},
			{Name: "银行存款"},
			{Name: "累计折旧"},
		},
	},
	//28 计算最终的生产成本
	{Action: "计算最终的生产成本",
		Dr: []Element{
			{Name: "生产成本"},
		},
		Cr: []Element{
			{Name: "制造费用"},
		},
	},
	//29 生产的物品入库
	{Action: "生产的物品入库",
		Dr: []Element{
			{Name: "库存商品"},
		},
		Cr: []Element{
			{Name: "生产成本"},
		},
	},
	//30 出售库存商品
	{Action: "出售库存商品",
		Dr: []Element{
			{Name: "应收账款"},
			{Name: "主营业务成本"},
		},
		Cr: []Element{
			{Name: "主营业务收入"},
			{Name: "应交税费"},
		},
	},
	//31 取得销售收入
	{Action: "取得销售收入",
		Dr: []Element{
			{Name: "银行存款"},
		},
		Cr: []Element{
			{Name: "主营业务收入"},
			{Name: "应交税费"},
		},
	},
	//31 结转销售成本
	{Action: "结转销售成本",
		Dr: []Element{
			{Name: "主营业务成本"},
		},
		Cr: []Element{
			{Name: "库存商品"},
		},
	},
	//32 委托加工-交付材料和库存商品
	{Action: "委托加工-交付材料和库存商品",
		Dr: []Element{
			{Name: "委托加工物资"},
		},
		Cr: []Element{
			{Name: "原材料"},
			{Name: "库存商品"},
		},
	},
	//32 委托加工-支付加工费用
	{Action: "委托加工-支付加工费用",
		Dr: []Element{
			{Name: "委托加工物资"},
			{Name: "应交税费"},
		},
		Cr: []Element{
			{Name: "银行存款"},
		},
	},
	//33 委托加工-收回完成品
	{Action: "委托加工-收回完成品",
		Dr: []Element{
			{Name: "库存商品"},
		},
		Cr: []Element{
			{Name: "委托加工物资"},
		},
	},
	//34 销售货物
	{Action: "销售货物",
		Dr: []Element{
			{Name: "应收账款"},
		},
		Cr: []Element{
			{Name: "主营业务收入"},
			{Name: "应交税费"},
		},
	},
	//35 结转成本
	{Action: "结转成本",
		Dr: []Element{
			{Name: "主营业务成本"},
		},
		Cr: []Element{
			{Name: "库存商品"},
		},
	},
	//36 计提存货跌价
	{Action: "计提存货跌价",
		Dr: []Element{
			{Name: "资产减值损失"},
		},
		Cr: []Element{
			{Name: "存货跌价准备"},
		},
	},
	//37 存货盘盈
	{Action: "存货盘盈",
		Dr: []Element{
			{Name: "库存商品"},
		},
		Cr: []Element{
			{Name: "待处理财产损溢-待处理流动资产损溢"}, //最终会计入 管理费用
		},
	},
	//38 存货盘亏
	{Action: "存货盘盈",
		Dr: []Element{
			{Name: "待处理财产损溢-待处理流动资产损溢"}, //最终会计入 管理费用
		},
		Cr: []Element{
			{Name: "库存商品"},
		},
	},
	//38 购买债券
	{Action: "购买债券",
		Dr: []Element{
			{Name: "债权投资"},
			{Name: "应收利息"},
		},
		Cr: []Element{
			{Name: "银行存款"},
		},
	},
	//38 收到利息
	{Action: "收到利息",
		Dr: []Element{
			{Name: "银行存款"},
		},
		Cr: []Element{
			{Name: "应收利息"},
		},
	},
	//后续还有通过子公司和控股公司的收益
}
