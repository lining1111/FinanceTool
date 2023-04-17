package analysis

import (
	"errors"
	"fmt"
	"math"
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
	max := da[0]
	index := 0
	for k, v := range da {
		if v > max {
			max = v
			index = k
		}
	}
	fmt.Printf("在第%d找到最大值%f", index, max)
	return max
}

//Min 最小值
func (da DA) Min() float64 {
	min := da[0]
	index := 0
	for k, v := range da {
		if v < min {
			min = v
			index = k
		}
	}
	fmt.Printf("在第%d找到最小值%f", index, min)
	return min
}

//Avg 平均值
func (da DA) Avg() float64 {
	sum := 0.0
	for _, v := range da {
		sum += v
	}
	return sum / float64(len(da))
}

//AvgChaArray 与平均值的算术差值数组
func (da DA) AvgChaArray() DA {
	avg := da.Avg()
	ret := DA{}
	for _, v := range da {
		item := v - avg
		ret = append(ret, item)
	}
	return ret
}

//AverageDeviation 平均差值：总体所有单位与其算术平均数的离差绝对值的算术平均数
func (da DA) AverageDeviation() float64 {
	aca := da.AvgChaArray()
	return aca.Avg()
}

//如果期望就是平均数，则均方误差就是方差

//Variance 方差：各组数据与它们的平均数的差的平方，和后的平均。衡量这组数据的波动大小
func (da DA) Variance() float64 {
	aca := da.AvgChaArray()
	sum := 0.0
	for _, v := range aca {
		sum += math.Pow(math.Abs(v), 2)
	}
	return sum / float64(len(aca))
}

//StdDeviation 标准差，也叫均方差，即方差的开方。表示分散程度
func (da DA) StdDeviation() float64 {
	return math.Sqrt(da.Variance())
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
	avg_da1 := da1.Avg()
	avg_da2 := da2.Avg()
	sum := 0.0
	for i := 0; i < len(da1); i++ {
		sum += ((da1[i] - avg_da1) * (da2[i] - avg_da2))
	}
	return sum / float64(len(da1)), nil

}

//Corr 相关系数  如果corr为1,则完全正线性相关;如果corr为-1,则完全负线性相关;corr为0，则无有线性相关系数;
//corr大于-1小于0，则负相关，系数的程序反映了相关性的强弱;
//corr大于0小于1,则正相关，系数的程序反映了相关性的强弱;
func Corr(da1 DA, da2 DA) (float64, error) {
	if len(da1) != len(da2) {
		return 0, errors.New("da1 len should equal da2")
	}
	cov, _ := Cov(da1, da2)
	return cov / (da1.StdDeviation() * da2.StdDeviation()), nil
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

//NetPresentValue NPV净现值
func NetPresentValue(rate float64, values []float64) float64 {
	npv := 0.0
	nper := len(values)
	for i := 1; i <= nper; i++ {
		npv += values[i-1] / math.Pow(1+rate, float64(i))
	}
	return npv
}

//Kvalue 资本化成本 c 原始成本 s 残值 r利率 m维修成本 n 使用年限
func Kvalue(c float64, s float64, r float64, m float64, n uint) float64 {
	npv_fix := (c - s) / (math.Pow(1+r, float64(n)) - 1) //无限维护成本的现值,复利下贴现计算，再算合
	npv_recover := m / r                                 //无限替换的现值，等比数列
	return c + npv_fix + npv_recover
}
