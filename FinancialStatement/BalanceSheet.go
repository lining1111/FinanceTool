package FinancialStatement

//BalanceSheet 资产负债表 内有恒等式(Assets==LiabilityAndShareHoldersEquity)即资产合计一定等于负债及股东权益合计
type BalanceSheet struct {
	Assets                         Assets
	LiabilityAndShareHoldersEquity LiabilityAndShareHoldersEquity
}

//Assets 资产
type Assets struct {
	CurrentAssets  CurrentAssets
	LongTermAssets LongTermAssets
	Total          float64 //合计
}

//CurrentAssets 流动资产
type CurrentAssets struct {
	Cash               float64 //货币资金
	AccountsReceivable float64 //应收账款
	Prepayment         float64 //预付账款
	Inventory          float64 //存货
	DeferredExpense    float64 //待摊费用
	Others             float64 //其他
	Total              float64 //流动资产合计
}

//LongTermAssets 长期资产
type LongTermAssets struct {
	LongTermInvestment             float64 //长期投资
	FixedAssets                    float64 //固定资产
	IntangibleAssetsAndOtherAssets float64 //无形及其他资产
	Total                          float64 //合计
}

//LiabilityAndShareHoldersEquity 负债及股东权益
type LiabilityAndShareHoldersEquity struct {
	CurrentLiability   CurrentLiability
	LongTermLiability  LongTermLiability
	ShareholdersEquity ShareholdersEquity
	Total              float64 //合计
}

//CurrentLiability 流动负债
type CurrentLiability struct {
	ShortTermLoans  float64 //短期借款
	AccountsPayable float64 //应付账款
	Others          float64 //其他
	Total           float64 //流动负债合计
}

//LongTermLiability 长期负债
type LongTermLiability struct {
	LongTermLoans float64 //长期借款
	BondsPayable  float64 //应付债券
	Others        float64 //其他
	Total         float64 //合计
}

//ShareholdersEquity 股东权益
type ShareholdersEquity struct {
	PaidInCapital        float64 //股本
	CapitalReserves      float64 //资本公积
	SurplusReserves      float64 //盈余公积
	UndistributedProfits float64 //未分配利润
	Total                float64 //合计
}
