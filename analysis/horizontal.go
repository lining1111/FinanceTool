package analysis

import (
	"math"
	"sort"
)

//------------------------横向分析---------------------------//

//HData 某个指标某行业的一个时间点的所有股票值
type HData struct {
	Scode string  //股票代码
	Val   float64 //数值
}

func (h HData) Less(h1 HData) bool {
	cha := h.Val - h1.Val
	if cha < 0 {
		return true
	} else {
		return false
	}
}

type HDA []HData

func (hda HDA) Len() int {
	return len(hda)
}

func (hda HDA) Less(i, j int) bool {
	return hda[i].Less(hda[j])
}

func (hda HDA) Swap(i, j int) {
	hda[i], hda[j] = hda[j], hda[i]
}

func (hda *HDA) Sort() {
	sort.Sort(*hda)
}

//对于按行业区分、某个时间点、某个指标的数据分析，属于横向分析，应该从正态分布方向出发

type HStatistics struct {
	label []string //股票代码的数组，临时区用来做后续的统计计算
	da    DA       //剔除股票的数据数组，临时区用来做后续的统计计算
	Max   float64  //最大值
	Min   float64  //最小值
	Avg   float64  //平均值

	Chada DA //差数组，临时区用来做后续的统计计算
	//纵向-基于行业区分，看行业稳定性
	AverageDeviation float64 //平均差值
	Variance         float64 //方差
	StdDeviation     float64 //标准差

	//横向-基于指标值从小到大的股票数量和的数组
	StepLabel      []float64 //数值数组
	StepCountArray []float64 //股票数量累计数组

}

//Cal step:衡量该指标时的步进值 ，比如xx率就是0.01 而金额就是1
func (hs *HStatistics) Cal(src HDA, step float64) {
	//将数值和label分离开处理
	for _, v := range src {
		hs.label = append(hs.label, v.Scode)
		hs.da = append(hs.da, v.Val)
	}

	//计算最大值、最小值、平均值
	hs.Max = hs.da.Max()
	hs.Min = hs.da.Min()
	hs.Avg = hs.da.Avg()
	//与平均值的差数组
	for _, v := range hs.da {
		item := v - hs.Avg
		hs.Chada = append(hs.Chada, item)
	}

	//纵向
	//平均差值
	hs.AverageDeviation = hs.Chada.Avg()
	//方差
	sum := 0.0
	for _, v := range hs.Chada {
		sum += math.Pow(math.Abs(v), 2)
	}
	hs.Variance = sum / float64(len(hs.Chada))
	//标准差
	hs.StdDeviation = math.Sqrt(hs.Variance)
	//横向
	//得到累计值数组
	length := uint64((hs.Max-hs.Min)/step + 1)
	hs.StepLabel = make([]float64, length, length)
	hs.StepCountArray = make([]float64, length, length)
	for _, v := range hs.da {
		index := 0
		for k1, v1 := range hs.StepLabel {

			cha := v - v1
			cha1 := v - (v1 + step)
			if cha == 0 {
				index = k1
				break
			} else if cha1 == 0 {
				index = k1 + 1
				break
			} else if (cha > 0) && (cha1 < 0) {
				index = k1
				break
			}
		}
		hs.StepCountArray[index]++
	}

}

//RatioH 单个指标在某个时间点，某行业的记录
type RatioH struct {
	Ratio string //指标名称
	Eval  string //行业名称
	HDA   HDA    //当前时间所有该行业的
}

//RHA 单个指标在某个时间点，所有行业的记录
type RHA []RatioH
