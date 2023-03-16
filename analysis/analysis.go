package analysis

//分析的原理，就是通过统计学原理，对数据进行质量分析

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
