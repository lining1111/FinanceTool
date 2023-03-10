package cninfo

import "fmt"

func init() {
	if usToken {
		go tokenFresh()
	}
	//获取公共编码数据
	fmt.Println("获取公共编码数据")
	MPublicCodeL0.Get()             //获取总类及其子类的树形结构
	PubilicCodeL0.GetType()         //获取总类名称及编码
	PublicCode071.GetSubtype("071") //获取071子类
	PublicCode033.GetSubtype("033") //获取033子类
	P0005GetAll()                   //获取股票类型编码
}

var (
	Q1 = "-03-31"
	Q2 = "-06-30"
	Q3 = "-09-30"
	Q4 = "-12-31"
)

func Getrdate(year string, q string) string {
	return year + q
}

const APIPublicCodeData = "http://webapi.cninfo.com.cn/api/public/p_public0006"

//PublicCodeData 公共编码数据 "http://webapi.cninfo.com.cn/api/public/p_public0006"
//subtype	总类编码	string	否	总类编码，可以为空
//format	结果集格式	string	否	设置结果返回的格式，可选的有xml、json、csv、dbf
//@column	结果列选择	string	否	选择结果集中所需要的字段，多列用逗号分隔，如@column=a,b
//@limit	结果条数限制	int	否	设置结果返回的条数
//@orderby	结果集排序	string	否	设置结果集的格式，如 @orderby=id:desc @orderby=id:asc
type PublicCodeData struct {
	OBJECTID string //记录ID	BIGINT
	SORTCODE string //类目编码	VARCHAR
	SORTNAME string //类目名称	VARCHAR

	F001V string //总类编码	VARCHAR
	F002V string //总类名称	VARCHAR
	F004D string //终止日期	DATE
	F005V string //类目名称（英文）	VARCHAR
}

func getPublicCode(subtype string, result *[]PublicCodeData) error {
	params := map[string]string{
		"subtype": subtype,
	}
	return GetInfoByScodeDate(APIPublicCodeData, params, result)

}

type PC map[string]string

func (pc *PC) GetType() error {
	result := make([]PublicCodeData, 1, 10000)
	err := getPublicCode("", &result)
	if err != nil {
		return err
	} else {
		pc1 := make(PC)
		for _, v := range result {
			_, ok := pc1[v.F002V]
			if !ok {
				pc1[v.F002V] = v.F001V
			}
		}
		*pc = pc1

		return nil
	}
}

func (pc *PC) GetSubtype(subtype string) error {
	result := make([]PublicCodeData, 1, 10000)
	err := getPublicCode(subtype, &result)
	if err != nil {
		return err
	} else {
		pc1 := make(PC)
		for _, v := range result {
			pc1[v.SORTNAME] = v.SORTCODE
		}
		*pc = pc1

		return nil
	}
}

//总类编码 名称 编码
var PubilicCodeL0 PC

//key类目名称SORTNAME val 类目编码SORTCODE
var PublicCode071 PC //合并本期 合并上期 母公司本期 母公司上期
var PublicCode033 PC //定期报告 审计报告

func PublicCodeData_test() {
	result := make([]PublicCodeData, 1, 1000)
	params := map[string]string{
		"subtype": "071",
	}
	err := GetInfoByScodeDate(APIPublicCodeData, params, &result)
	fmt.Println(err)
}

type PublicCodeL0 struct {
	L1 []PublicCodeL1 //总类
}
type PublicCodeL1 struct {
	Code  string         //总类编码
	Name  string         //总类名称
	EDate string         //中止上市时间
	EName string         //中文名称
	L2    []PublicCodeL2 //子类
}
type PublicCodeL2 struct {
	Code string //子类编码
	Name string //子类名称
}

var MPublicCodeL0 PublicCodeL0

func (pcl0 *PublicCodeL0) Get() error {
	result := make([]PublicCodeData, 0, 10000)
	err := getPublicCode("", &result)
	if err != nil {
		return err
	} else {
		curNameL1 := ""
		for _, v := range result {
			if curNameL1 != v.F002V {
				//新的总类
				curNameL1 = v.F002V
				l1 := PublicCodeL1{}
				l1.Code = v.F001V
				l1.Name = v.F002V
				l2 := PublicCodeL2{}
				l2.Code = v.SORTCODE
				l2.Name = v.SORTNAME
				l1.L2 = append(l1.L2, l2)
				pcl0.L1 = append(pcl0.L1, l1)
			} else {
				//旧的总类
				for k1, v1 := range pcl0.L1 {
					if v1.Name == v.F002V {
						//找到旧类的数组索引
						l2 := PublicCodeL2{}
						l2.Code = v.SORTCODE
						l2.Name = v.SORTNAME
						pcl0.L1[k1].L2 = append(pcl0.L1[k1].L2, l2)
					}
				}
			}

		}
		return nil
	}
}

//Public0005 证券类别编码数据 http://webapi.cninfo.com.cn/api/public/p_public0005
type Public0005 struct {
	PARENTCODE string //父类编码	VARCHAR
	SORTCODE   string //类目编码	VARCHAR
	SORTNAME   string //类目名称	VARCHAR
	F002V      string //类目名称（英文）	VARCHAR
	F001D      string //终止日期	DATE
}

var P0005 []Public0005

func P0005GetAll() error {
	url := "http://webapi.cninfo.com.cn/api/public/p_public0005"
	p5 := make([]Public0005, 1, 10000)
	params := map[string]string{}
	err := GetInfoByScodeDate(url, params, &p5)
	if err != nil {
		return err
	} else {
		for _, v := range p5 {
			P0005 = append(P0005, v)
		}
		return nil
	}
}

//获取A股的所有股票代码信息
