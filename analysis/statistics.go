package analysis

import (
	"fmt"
	"math"
)

//针对数组的统计计算相关

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
