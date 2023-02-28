package BalanceSheet

import "errors"

//BalanceSheet 资产负债表 内有恒等式(Assets==LiabilityAndShareHoldersEquity)即资产合计一定等于负债及股东权益合计
type BalanceSheet struct {
	A   Assets
	LEQ LiabilityAndEquity
}

func (bs BalanceSheet) CheckIdentity() (bool, error) {
	if bs.A.Total > bs.LEQ.Total {
		return false, errors.New("资产多于负债")
	} else if bs.A.Total < bs.LEQ.Total {
		return false, errors.New("资产少于负债")
	} else {
		return true, nil
	}
}

//Assets 资产
type Assets struct {
	CA    CurrentAssets
	LTA   LongTermAssets
	Total float64 //合计
}

func (a *Assets) CalTotal() float64 {
	a.Total = a.CA.Total + a.LTA.Total
	return a.Total
}

//CurrentAssets 流动资产
type CurrentAssets struct {
	Cash               float64 //货币资金
	FVTOCI             float64 //Fair Value through Other Comprehensive Income以公允价值计量且起变动计入其他综合收益
	AccountsReceivable float64 //应收账款
	PrepaidExpense     float64 //预付账款
	Inventory          float64 //存货
	NACHSADGCHS        float64 //Non-current assets classified as held for sale and assets in disposal groups classified as held for sale被划分为持有待售的非流动资产及被划分为持有待售的处置组中的资产
	AFSFA              float64 //Available-for-sale financial assets可供出售金融资产
	HTMSecurities      float64 //held-to-maturity securities持有至到期投资
	DeferredExpense    float64 //待摊费用
	Others             float64 //其他
	Total              float64 //流动资产合计
}

func (ca *CurrentAssets) CalTotal() float64 {
	ca.Total = ca.Cash + ca.FVTOCI + ca.AccountsReceivable + ca.PrepaidExpense +
		ca.Inventory + ca.NACHSADGCHS + ca.AFSFA + ca.HTMSecurities + ca.DeferredExpense + ca.Others
	return ca.Total
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

func (lta *LongTermAssets) CalTotal() float64 {
	lta.Total = lta.LongTermInvestment + lta.REHI + lta.FixedAssets + lta.BiologicalAsset +
		lta.IntangibleAssetsAndOtherAssets + lta.DeferredTaxAsset
	return lta.Total
}

//LiabilityAndEquity 负债及股东权益
type LiabilityAndEquity struct {
	Li    Liability
	SHE   Equity
	Total float64 //合计
}

func (leq *LiabilityAndEquity) CalTotal() float64 {
	leq.Total = leq.Li.Total + leq.SHE.Total
	return leq.Total
}

//Liability 负债
type Liability struct {
	Cli   CurrentLiability
	Ltli  LongTermLiability
	Total float64
}

func (li Liability) CalTotal() float64 {
	li.Total = li.Cli.Total + li.Ltli.Total
	return li.Total
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

func (cli *CurrentLiability) CalTotal() float64 {
	cli.Total = cli.ShortTermLoans + cli.FVTPL + cli.AccountsPayable + cli.UnearnedRevenue +
		cli.CompensationPayable + cli.TaxPayable + cli.CALHFS + cli.Others
	return cli.Total
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

func (ltli *LongTermLiability) CalTotal() float64 {
	ltli.Total = ltli.LongTermLoans + ltli.BondsPayable + ltli.LongTermPayable + ltli.EstimatedLiabilities +
		ltli.DeferredIncomeTaxLiabilities + ltli.Others
	return ltli.Total
}

//Equity 股东权益
type Equity struct {
	PaidInCapital        float64 //股本
	CapitalReserves      float64 //资本公积
	SurplusReserves      float64 //盈余公积
	UndistributedProfits float64 //未分配利润
	Total                float64 //合计
}

func (eq *Equity) CalTotal() float64 {
	eq.Total = eq.PaidInCapital + eq.CapitalReserves + eq.SurplusReserves + eq.UndistributedProfits
	return eq.Total
}
