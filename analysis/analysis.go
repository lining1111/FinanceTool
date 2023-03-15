package analysis

import (
	"fmt"
	"math"
	"sort"
	"time"
)

//GetAims 根据目标获得一些小目标
func GetAims(attr string) (aim []string, err error) {

	return
}

//GetRatios 将每一个小目标分解成若干实际的财务指标,需要输入特点的评价系统名称
func GetRatios(aim string, eval string) (ratio []float64, err error) {

	return
}

//GetScodes 根据各个指标的要求，根据比较分析法，选出当期符合要求的股票代码集合
func GetScodes(ratio float64) (scode []string, err error) {

	return
}

//DataUnit 单个指标的历史点记录
type DataUnit struct {
	Val  float64
	Date string //2023-03-31
}

var loc, _ = time.LoadLocation("Asia/Shanghai")

//Before 从时间上，判断d是否在d1的前面
func (d DataUnit) Before(d1 DataUnit) (bool, error) {
	time0, err := time.ParseInLocation("2006-01-02", d.Date, loc)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	time1, err := time.ParseInLocation("2006-01-02", d1.Date, loc)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	isBefore := time0.Before(time1)

	return isBefore, nil
}

//GetSingleRatioHistory 根据股票和单个指标获得历史数据集合
func GetSingleRatioHistory(scode string, ratio string) (ratioHistory []DataUnit, err error) {

	return
}

//RatioHistory 单个指标的历史记录
type RatioHistory struct {
	Ratio       string
	DataHistory []DataUnit
}

//GetRatiosHistory 根据单个股票，获取指标集合的历史数据
func GetRatiosHistory(scode string, ratios []string) (ratiosHistory []RatioHistory, err error) {

	return
}

//History 单只股票的历史指标集合
type History struct {
	scode         string
	RatiosHistory []RatioHistory
}

//根据各个指标的历史数据，通过统计方法，作出分析
type DH []DataUnit

func (dh DH) Len() int {
	return len(dh)
}

func (dh DH) Less(i, j int) bool {
	ret, err := dh[i].Before(dh[j])
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ret
}

func (dh DH) Swap(i, j int) {
	dh[i], dh[j] = dh[j], dh[i]
}

//Sort 按时间先后做排序
func (dh *DH) Sort() {
	sort.Sort(*dh)
}

func DHSortTest() {
	dh := DH{
		{1, "2023-12-31"},
		{2, "2023-09-30"},
		{3, "2023-06-30"},
		{4, "2023-03-31"},
		{5, "2022-03-31"},
		{6, "2022-06-30"},
		{7, "2022-09-30"},
		{8, "2022-12-31"},
	}
	dh.Sort()
	fmt.Println(dh)
}

//Avg 平均值
func (dh DH) Avg() float64 {
	sum := 0.0
	for _, v := range dh {
		sum += v.Val
	}
	return sum / float64(len(dh))
}

//AvgCha 与平均值的算术差值
func (dh DH) AvgCha() DH {
	avg := dh.Avg()
	ret := DH{}
	for _, v := range dh {
		item := DataUnit{}
		item.Date = v.Date
		item.Val = v.Val - avg
		ret = append(ret, item)
	}

	return ret
}

//AverageDeviation 平均差值：总体所有单位与其算术平均数的离差绝对值的算术平均数
func (dh DH) AverageDeviation() float64 {
	avgCha := dh.AvgCha()
	return avgCha.Avg()
}

//如果期望就是平均数，则均方误差就是方差

//Variance 方差：各组数据与它们的平均数的差的平方，和后的平均。衡量这组数据的波动大小
func (dh DH) Variance() float64 {
	avgCha := dh.AvgCha()
	sum := 0.0
	for _, v := range avgCha {
		sum += math.Pow(math.Abs(v.Val), 2)
	}
	return sum / float64(len(avgCha))
}

//StdDeviation 标准差，也叫均方差，即方差的开方。表示分散程度
func (dh DH) StdDeviation() float64 {
	return math.Sqrt(dh.Variance())
}

//Statistics 计算一组数据的统计指标
type Statistics struct {
	Avg    float64 //平均值
	AvgCha DH      //与平均值的差数组

	AverageDeviation float64 //平均差值
	Variance         float64 //方差
	StdDeviation     float64 //标准差
}

func (r *Statistics) Cal(src DH) {
	//计算平均值
	r.Avg = src.Avg()
	//与平均值的差数组
	avgCha := DH{}
	for _, v := range src {
		item := DataUnit{}
		item.Date = v.Date
		item.Val = v.Val - r.Avg
		avgCha = append(avgCha, item)
	}
	r.AvgCha = avgCha
	//平均差值
	r.AverageDeviation = r.AvgCha.Avg()
	//方差
	sum := 0.0
	for _, v := range r.AvgCha {
		sum += math.Pow(math.Abs(v.Val), 2)
	}
	r.Variance = sum / float64(len(avgCha))
	//标准差
	r.StdDeviation = math.Sqrt(r.Variance)
}

func ResultTest() {
	src := DH{
		{1, "2023-12-31"},
		{2, "2023-09-30"},
		{3, "2023-06-30"},
		{4, "2023-03-31"},
		{5, "2022-03-31"},
		{6, "2022-06-30"},
		{7, "2022-09-30"},
		{8, "2022-12-31"},
	}
	src.Sort()
	result := Statistics{}
	result.Cal(src)
	fmt.Println(result.AvgCha)

}

//Result 分析一组数据的结果
//长期投资的话，应该选择期望值对收益有帮助的且稳定性强的。
//短期投资应该在确认有上升空间后，选择稳定性差的，但稳定性差代表它上升和下降均有很大可能出现，所以不建议短期投资
type Result struct {
	Src        DH         //原始数据
	Statistics Statistics //统计中间结果
	//指标
	Hope float64 //期望值 这里取平均数 Avg

	Stability0 float64 //稳定性1 这里取平均差值 AverageDeviation
	Stability1 float64 //稳定性2 这里取标准差 StdDeviation

}
