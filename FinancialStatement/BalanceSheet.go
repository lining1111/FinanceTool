package FinancialStatement

//BalanceSheet 资产负债表 内有恒等式(Assets==LiabilityAndShareHoldersEquity)即资产合计一定等于负债及股东权益合计
type BalanceSheet struct {
	A        Assets
	LiAndSHE LiabilityAndShareHoldersEquity
}

//Assets 资产
type Assets struct {
	CA    CurrentAssets
	LTA   LongTermAssets
	Total float64 //合计
}

//CurrentAssets 流动资产
type CurrentAssets struct {
	Cash               float64 //货币资金
	FVTOCI             float64 //Fair Value through Other Comprehensive Income以公允价值计量且起变动计入其他综合收益
	AccountsReceivable float64 //应收账款
	PrepaidExpense     float64 //预付账款
	Inventory          float64 //存货
	NACHSADGCHS        float64 //Non-current assets classified as held for sale and assets in disposal groups classified as held for sale被划分为持有待售的非流动资产及被划分为持有待售的处置组中的资产
	AFS                float64 //Available-for-sale financial assets可供出售金融资产
	HTMSecurities      float64 //held-to-maturity securities持有至到期投资
	DeferredExpense    float64 //待摊费用
	Others             float64 //其他
	Total              float64 //流动资产合计
}

//LongTermAssets 长期资产
type LongTermAssets struct {
	LongTermInvestment             float64 //长期投资
	REHI                           float64 //Real Estate Held for Investment投资性房地产
	FixedAssets                    float64 //固定资产
	BiologicalAsset                float64 //生物资产
	IntangibleAssetsAndOtherAssets float64 //无形及其他资产
	DeferredTaxAsset               float64 //递延所得税资产
	Total                          float64 //合计
}

//LiabilityAndShareHoldersEquity 负债及股东权益
type LiabilityAndShareHoldersEquity struct {
	CLi   CurrentLiability
	LTLi  LongTermLiability
	SHE   ShareholdersEquity
	Total float64 //合计
}

//CurrentLiability 流动负债
type CurrentLiability struct {
	ShortTermLoans      float64 //短期借款
	FVTPL               float64 //Fair Value through Profit and Loss以公允价值计量且其变动计入当期损益的金融负债
	AccountsPayable     float64 //应付账款
	UnearnedRevenue     float64 //预收款项
	CompensationPayable float64 //应付职工薪酬
	TaxPayable          float64 //应交税费
	CALHFS              float64 //Classified as liabilities held for sale被划分为持有待售的处置组中的负债
	Others              float64 //其他
	Total               float64 //流动负债合计
}

//LongTermLiability 长期负债
type LongTermLiability struct {
	LongTermLoans                float64 //长期借款
	BondsPayable                 float64 //应付债券
	LongTermPayable              float64 //长期应付款
	EstimatedLiabilities         float64 //预计负债
	DeferredIncomeTaxLiabilities float64 //递延所得税负债
	Others                       float64 //其他
	Total                        float64 //合计
}

//ShareholdersEquity 股东权益
type ShareholdersEquity struct {
	PaidInCapital        float64 //股本
	CapitalReserves      float64 //资本公积
	SurplusReserves      float64 //盈余公积
	UndistributedProfits float64 //未分配利润
	Total                float64 //合计
}
