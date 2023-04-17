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
    假如我们有一台打印机，原始成本是65000美元，12年后估计残值为5000美元，机器生产率是每年20000本书，维护成本是3000美元，
    公司的工程师决定安装一个额外的配件，能够把机器的生产效率提高到每年30000本书，而不会影响到维护或使用寿命，如果投资率是8%
    那么公司能够接受多少钱达到提高生产率的目的
    这里前提是 花钱后，至少平均资本产出是相等的
    即 Kb/Pb=Ka/Pa 公式1
    技术改进前的Kb计算
    Kb = c+（c-s）/(（1+r）^n-1)+m/r = 65000+ 39521+ 37500=142021
    技术改进后的ka计算
    Ka=c+x+（c+x-s）/(（1+r）^n-1)+m/r=(227271+3x)/1.5
    通过公式1的到
    x=30758