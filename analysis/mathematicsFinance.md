##数理金融
    此文件对应同名的go文件
### 1 数学简介
    SumAP 等差数列的和 first 开始值 c个数 d 差值
    SumGP 等比数列的和 first 开始值 c个数 r 比率
    Factorial 阶乘 n
    Amn 排列
    Cmn 组合

    Max 最大值
    Min 最小值
    Avg 平均值
    Variance 方差
    StdDeviation 标准差
    Cov 协方差
    Corr 相关系数  如果corr为1,则完全正线性相关;如果corr为-1,则完全负线性相关;corr为0，则无有线性相关系数;
    corr大于-1小于0，则负相关，系数的程序反映了相关性的强弱;
    corr大于0小于1,则正相关，系数的程序反映了相关性的强弱;
### 2 货币时间价值
    pv cv r n 是货币时间的基本元素，求一个元素的值，必须知道除它外的其余元素的值 
    i = fv- cv
    PVInSingleRate 单利模式下现值计算 i 利息总和 r 利率 n 到期期限
    FVInSingleRate 单利模式下将来值计算 pv 现值 r 利率 n 到期期限
    ERInSingleRate 单利模式下实际利率计算 or 名义利率
    ORInSingleRate 单利模式下名义利率计算 er 实际利率
    
    PVInMulRate 复利模式下现值计算 fv 将来值 r 利率 n 到期期限
    FVInMulRate 复利模式下将来值计算 pv 现值 r 利率 n 到期期限
    DFInMulRate 复利模式下贴现因子计算 r 利率 n 到期期限
    CPInMulRate 复合期间，即由将来值fv、现值pv、利率r 计算需要多久才能使现值变将来值
    ERInMulRate 复利模式下，由名义利率r 变换期m计算有效利率r 当m无穷大时，er=e^r-1
### 3 债务和租赁

### 4 资本预算和折旧
    NetPresentValue NPV净现值
    Kvalue 资本化成本
    具体的应用
    
    