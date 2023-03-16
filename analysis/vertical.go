package analysis

import (
	"fmt"
	"math"
	"sort"
	"time"
)

//-----------------------纵向分析--------------------//

//VData 单个指标的历史点记录
type VData struct {
	Val  float64
	Date string //2023-03-31
}

var loc, _ = time.LoadLocation("Asia/Shanghai")

//Before 从时间上，判断d是否在d1的前面
func (d VData) Before(d1 VData) (bool, error) {
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
func GetSingleRatioHistory(scode string, ratio string) (ratioHistory []VData, err error) {

	return
}

//RatioV 单个指标的历史记录
type RatioV struct {
	Ratio string
	VDA   []VData
}

//GetRatiosHistory 根据单个股票，获取指标集合的历史数据
func GetRatiosHistory(scode string, ratios []string) (ratiosHistory []RatioV, err error) {

	return
}

//VRecord 单只股票的历史指标集合
type VRecord struct {
	Scode string   //股票代码
	RVA   []RatioV //所有指标的历史合集
}

//根据各个指标的历史数据，通过统计方法，作出分析

type VDA []VData

func (vda VDA) Len() int {
	return len(vda)
}

func (vda VDA) Less(i, j int) bool {
	ret, err := vda[i].Before(vda[j])
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ret
}

func (vda VDA) Swap(i, j int) {
	vda[i], vda[j] = vda[j], vda[i]
}

//Sort 按时间先后做排序
func (vda *VDA) Sort() {
	sort.Sort(*vda)
}

func DHSortTest() {
	dh := VDA{
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

//VStatistics 计算一组数据的统计指标
type VStatistics struct {
	label []string //日期的数组，临时区用来做后续的统计计算
	da    DA       //剔除日期的数据数组，临时区用来做后续的统计计算
	Max   float64  //最大值
	Min   float64  //最小值
	Avg   float64  //平均值

	Chada DA //差数组，临时区用来做后续的统计计算

	AverageDeviation float64 //平均差值
	Variance         float64 //方差
	StdDeviation     float64 //标准差
}

func (vs *VStatistics) Cal(src VDA) {
	//将数值和label分离开处理
	for _, v := range src {
		vs.label = append(vs.label, v.Date)
		vs.da = append(vs.da, v.Val)
	}

	//计算最大值、最小值、平均值
	vs.Max = vs.da.Max()
	vs.Min = vs.da.Min()
	vs.Avg = vs.da.Avg()
	//与平均值的差数组
	for _, v := range vs.da {
		item := v - vs.Avg
		vs.Chada = append(vs.Chada, item)
	}

	//平均差值
	vs.AverageDeviation = vs.Chada.Avg()
	//方差
	sum := 0.0
	for _, v := range vs.Chada {
		sum += math.Pow(math.Abs(v), 2)
	}
	vs.Variance = sum / float64(len(vs.Chada))
	//标准差
	vs.StdDeviation = math.Sqrt(vs.Variance)
}

func ResultTest() {
	src := VDA{
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
	result := VStatistics{}
	result.Cal(src)
	fmt.Println(result.StdDeviation)

}

//VResult 分析一组数据的结果
//长期投资的话，应该选择期望值对收益有帮助的且稳定性强的。
//短期投资应该在确认有上升空间后，选择稳定性差的，但稳定性差代表它上升和下降均有很大可能出现，所以不建议短期投资
type VResult struct {
	Src        VDA         //原始数据
	Statistics VStatistics //统计中间结果
	//指标
	Hope float64 //期望值 这里取平均数 Avg

	Stability0 float64 //稳定性1 这里取平均差值 AverageDeviation
	Stability1 float64 //稳定性2 这里取标准差 StdDeviation

}
