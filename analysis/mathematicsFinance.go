package analysis

import (
	"errors"
	"fmt"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
	"math"
)

const (
	MaxIterations = 30
	Precision     = 1e-6
)

/**
本章主要是复写机械工业出版社的《数理金融》的公式
*/

//chapter 1 数学简介

//SumAP 等差数列的和 first 开始值 c个数 d 差值
func SumAP(first float64, c int, d float64) float64 {
	end := first + float64(c-1)*d
	return (float64(c) / 2.0) * (first + end)
}

func TestSumAP() {
	fmt.Println(SumAP(0.1, 5, 0.3))
}

//SumGP 等比数列的和 first 开始值 c个数 r 比率
func SumGP(first float64, c int, r float64) float64 {
	return first * ((1 - math.Pow(r, float64(c))) / (1 - r))
}

func TestSumGP() {
	fmt.Println(SumGP(0.1, 5, 3))
}

//Factorial 阶乘 n
func Factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

//Amn 排列
func Amn(m uint, n uint) (uint, error) {
	if m < n {
		return 0, errors.New("m should big than n")
	}
	return Factorial(m) / Factorial(m-n), nil
}

//Cmn 组合
func Cmn(m uint, n uint) (uint, error) {
	if m < n {
		return 0, errors.New("m should big than n")
	}
	return Factorial(m) / (Factorial(m-n) * Factorial(n)), nil
}

func TestAC() {
	amn, err := Amn(6, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(amn)
	}
	cmn, err := Cmn(6, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cmn)
	}
}

//期望 方差 标准差 协相关 相关系数

//DA 一维数组
type DA []float64

//Max 最大值
func (da DA) Max() float64 {
	return floats.Max(da)
}

//Min 最小值
func (da DA) Min() float64 {
	return floats.Min(da)
}

//Avg 平均值
func (da DA) Avg() float64 {
	return stat.Mean(da, nil)
}

//如果期望就是平均数，则均方误差就是方差

//Variance 方差：各组数据与它们的平均数的差的平方，和后的平均。衡量这组数据的波动大小
func (da DA) Variance() float64 {
	return stat.Variance(da, nil)
}

//StdDeviation 标准差，也叫均方差，即方差的开方。表示分散程度
func (da DA) StdDeviation() float64 {
	return stat.StdDev(da, nil)
}

//IsNormalDistribution 是否为正态分布
func (da DA) IsNormalDistribution() bool {
	//首先计算均值和标准差
	avg := da.Avg()
	stdD := da.StdDeviation()
	sum := 0.0
	for _, v := range da {
		sum += v
	}

	//看三次 均值+-标准差的覆盖率
	region := make([]float64, 2, 2) //区间
	pr := make([]float64, 3, 3)     //概率
	for i := 1; i < 4; i++ {
		region[0] = avg - float64(i)*stdD
		region[1] = avg + float64(i)*stdD
		total := 0.0
		for _, v := range da {
			if (v > region[0]) && (v <= region[1]) {
				total += v
			}
		}
		pr[i-1] = total / sum

		fmt.Println("区间 ", region, "概率 ", pr[i-1])
	}
	//函数曲线下68.268949%的面积在平均数左右的一个标准差范围内,
	//95.449974%的面积在平均数左右两个标准差的范围内,
	//99.730020%的面积在平均数左右三个标准差的范围内。
	if pr[0] >= 0.682 {
		if pr[1] >= 0.954 {
			if pr[2] >= 0.997 {
				return true
			}
		}
	}

	return false
}

//Cov 协方差
func Cov(da1 DA, da2 DA) (float64, error) {
	if len(da1) != len(da2) {
		return 0, errors.New("da1 len should equal da2")
	}
	cov := stat.Covariance(da1, da2, nil)

	return cov, nil

}

func TestCov() {
	da1 := []float64{1, 2, 3, 4, 5}
	da2 := []float64{0.2, 0.3, 0.4, 0.5, 0.6}
	cov, err := Cov(da1, da2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cov)
	}
}

//Corr 相关系数  如果corr为1,则完全正线性相关;如果corr为-1,则完全负线性相关;corr为0，则无有线性相关系数;
//corr大于-1小于0，则负相关，系数的程序反映了相关性的强弱;
//corr大于0小于1,则正相关，系数的程序反映了相关性的强弱;
func Corr(da1 DA, da2 DA) (float64, error) {
	if len(da1) != len(da2) {
		return 0, errors.New("da1 len should equal da2")
	}

	return stat.Correlation(da1, da2, nil), nil
}

func TestCorr() {
	da1 := []float64{1, 2, 3, 4, 5}
	da2 := []float64{0.2, 0.3, 0.4, 0.5, 0.6}
	corr, err := Corr(da1, da2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(corr)
	}
}

//chapter 2 货币的时间价值
//单利模式下，现值和将来值

//PVInSingleRate 单利模式下现值计算 i 利息总和 r 利率 n 到期期限
func PVInSingleRate(i float64, r float64, n uint) float64 {
	return i / (float64(n) * r)
}

//FVInSingleRate 单利模式下将来值计算 pv 现值 r 利率 n 到期期限
func FVInSingleRate(pv float64, r float64, n uint) float64 {
	return pv * (1 + r*float64(n))
}

//ARInSingleRate 单利模式下实际利率计算 or 名义利率
func ARInSingleRate(or float64) float64 {
	return or * (73 / 72)
}

//ORInSingleRate 单利模式下名义利率计算 ar 实际利率
func ORInSingleRate(ar float64) float64 {
	return ar * (72 / 73)
}

//PVInMulRate 复利模式下现值计算 fv 将来值 r 利率 n 到期期限
func PVInMulRate(fv float64, r float64, n uint) float64 {
	return fv * (math.Pow(1+r, float64(-n)))
}

//FVInMulRate 复利模式下将来值计算 pv 现值 r 利率 n 到期期限
func FVInMulRate(pv float64, r float64, n uint) float64 {
	return pv * (math.Pow(1+r, float64(n)))
}

//DFInMulRate 复利模式下贴现因子计算 r 利率 n 到期期限
func DFInMulRate(r float64, n uint) float64 {
	return 1 / (math.Pow(1+r, float64(n)))
}

//CPInMulRate 复合期间，即由将来值fv、现值pv、利率r 计算需要多久才能使现值变将来值
func CPInMulRate(fv float64, pv float64, r float64) float64 {
	return math.Log(fv/pv) / math.Log(1+r)
}

//ERInMulRate 复利模式下，由名义利率r 变换期m计算有效利率r
func ERInMulRate(r float64, m uint) float64 {
	return math.Pow(1+(r/float64(m)), float64(m)) - 1
}

//chapter 3 债务和租赁

//chapter 4 资本预算和折旧

//NetPresentValue NPV净现值 rate 利率 values 历史回报 i0初始投资
func NetPresentValue(rate float64, values []float64, i0 float64) float64 {
	npv := -i0
	nper := len(values)
	for i := 1; i <= nper; i++ {
		npv += values[i-1] / math.Pow(1+rate, float64(i))
	}
	return npv
}

//IIR 内部回报率 values 历史回报 i0 初始投资
func IIR(values []float64, i0 float64) (float64, error) {
	//以穷举法计算
	var r float64
	for i := 1; i < 1/Precision; i++ {
		r = float64(i) * Precision
		npv := NetPresentValue(r, values, i0)
		if npv < 1 && npv > 0 {
			break
		}
	}
	return r, nil
}

func TestIIR() {
	v := []float64{3600, 4200, 5500, 6300, 7500}
	i0 := 15000.0
	r, _ := IIR(v, i0)
	fmt.Println(r)
}

//Kvalue 资本化成本 c 原始成本 s 残值 r利率 m维修成本 n 使用年限
func Kvalue(c float64, s float64, r float64, m float64, n uint) float64 {
	npv_fix := (c - s) / (math.Pow(1+r, float64(n)) - 1) //无限维护成本的现值,复利下贴现计算，再算合
	npv_recover := m / r                                 //无限替换的现值，等比数列
	return c + npv_fix + npv_recover
}

func TestKvalue() {
	c := 65000.0
	s := 5000.0
	n := uint(12)
	r := 0.08
	m := 3000.0
	Kb := Kvalue(c, s, r, m, n)
	fmt.Println(Kb)
	//已知y 球x，则是求一元一次
	var x float64 //x消耗值
	a11 := 1.0
	a12 := 1 / (math.Pow(1+r, float64(n)) - 1)
	a := a11 + a12
	//  65000+x +(65000+x-5000)/a12+m/r
	//
	x = ((Kb / 20000 * 30000) - m/r - 65000 - 60000*a12) / a
	fmt.Println(x)
	//用矩阵的方式求解
	matrixRow := mat.Dense{}
	matrixCol := mat.NewDense(1, 1, []float64{a})
	d := mat.Dense{}
	d.Inverse(matrixCol)
	matrixResult := mat.NewDense(1, 1, []float64{Kb/20000*30000 - 65000 - 60000*a12 - m/r})
	matrixRow.Mul(matrixResult, &d)
	fmt.Println(mat.Formatted(&matrixRow, mat.Prefix(" "), mat.Squeeze()))
}

//chapter 5 盈亏平衡点和杠杆作用

//BEQ 盈亏平衡产量 fc固定成本 p单位价格 v单位成本 tp 目标利润
func BEQ(fc float64, p float64, v float64, tp float64) float64 {
	return (fc + tp) / (p - v)
}

//BER 盈亏平衡收益 fc固定成本 p单位价格 v单位成本 tp 目标利润
func BER(fc float64, p float64, v float64, tp float64) float64 {
	return (fc + tp) / (1 - v/p)
}

//CM 边际贡献 p单位价格 v单位成本
func CM(p float64, v float64) float64 {
	return p - v
}

//CBEQ 现金盈亏平衡产量 fc固定成本 nc任意非现金费用(可以是折旧费等) p单位价格 v单位成本 tp 目标利润
func CBEQ(fc float64, nc float64, p float64, v float64, tp float64) float64 {
	return (fc + tp - nc) / (p - v)
}

//BET 盈亏平衡时间 fc固定成本 p单位价格 v单位成本 q生产率(生产量除以时间) tl时滞(产品从销售出去拿到收益的时间)
func BET(fc float64, p float64, v float64, q float64, tl float64) float64 {
	return (fc + p*q*tl) / (q * (p - v))
}

//DOL 运营杠杆程度DOL=Q(p-v)/(Q(p-v)-FC)    固定成本FC 产出Q 单位成本v 单位价格p
func DOL(Q uint, p float64, v float64, FC float64) float64 {
	return float64(Q) * (p - v) / (float64(Q)*(p-v) - FC)
}

//DCL DCL=Q(p-v)/(Q(p-v)-FC-I-(Dps/(1-T))) 产量Q 单位价格p 单位成本v 固定成本FC 利息I 优先股分红Dps 税率T
func DCL(Q uint, p float64, v float64, FC float64, I float64, Dps float64, T float64) float64 {
	return float64(Q) * (p - v) / (float64(Q)*(p-v) - FC - I - (Dps / (1 - T)))
}

// chapter 6 投资
