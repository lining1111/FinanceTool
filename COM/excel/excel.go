package excel

import (
	"FinanceTool/Stock/HSB"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
)

/**
excel操作类
*/
const tagExcel = "excel"

//SetExcelTitle 根据结构体tag设置excel指定工作表的各列标题
func SetExcelTitle(path string, sheet string, v interface{}) error {
	const start = "A1"
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if _, ok := f.Sheet.Load(sheet); !ok {
		fmt.Println("sheet not exist create", sheet)
		_, err := f.NewSheet(sheet)
		if err != nil {
			fmt.Println(err)
		}
	}

	defer func() {
		if err := f.Save(); err != nil {
			fmt.Println(err)
		}
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("type ", reflect.TypeOf(v), " val ", reflect.ValueOf(v))
	kind := reflect.ValueOf(v).Kind()
	if kind != reflect.Struct {
		fmt.Println("expect struct")
		return errors.New("expect struct")
	}
	num := reflect.ValueOf(v).NumField()
	fmt.Println("结构体有字段", num)
	//遍历字段
	titles := make([]string, 0, num)
	for i := 0; i < num; i++ {
		tag := reflect.TypeOf(v).Field(i).Tag.Get(tagExcel)
		if len(tag) > 0 {
			titles = append(titles, tag)
		}
	}
	return f.SetSheetRow(sheet, start, &titles)
}

func GetExcelTtile(path string, sheet string) ([]string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if _, ok := f.Sheet.Load(sheet); !ok {
		fmt.Println("sheet not exist create", sheet)
		_, err := f.NewSheet(sheet)
		if err != nil {
			fmt.Println(err)
		}
	}
	defer func() {
		if err := f.Save(); err != nil {
			fmt.Println(err)
		}
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	//遍历字段
	titles := make([]string, 0, 1000)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}
	for i, row := range rows {
		if i == 0 {
			for _, colCell := range row {
				titles = append(titles, colCell)
			}

			break
		}
	}
	if len(titles) == 0 {
		return nil, errors.New("titles empty")
	}

	return titles, nil
}

//SetExcelData 根据结构体tag设置excel指定工作表的各列标题
func SetExcelData(path string, sheet string, v interface{}) error {
	vt := reflect.TypeOf(v).Elem()
	vv := reflect.ValueOf(v).Elem()
	if vt.Kind() != reflect.Slice || vt.Elem().Kind() != reflect.Struct || vv.Len() <= 0 {
		return errors.New("info not a slice interface or len 0")
	}

	//fmt.Println(reflect.TypeOf(v).Elem())
	numTag := vt.Elem().NumField()
	//获取标题的数组
	titles, err := GetExcelTtile(path, sheet)
	if err != nil {
		return err
	}
	if numTag != len(titles) {
		return errors.New("titles num not equal")
	}

	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if _, ok := f.Sheet.Load(sheet); !ok {
		fmt.Println("sheet not exist create", sheet)
		_, err := f.NewSheet(sheet)
		if err != nil {
			fmt.Println(err)
		}
	}
	defer func() {
		if err := f.Save(); err != nil {
			fmt.Println(err)
		}
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	row := 2

	for i := 0; i < vv.Len(); i++ {
		curvv := vv.Index(i)
		//获取v的tag
		for j := 0; j < numTag; j++ {
			//获取tag值
			tag := vt.Elem().Field(j).Tag.Get(tagExcel)
			col := 0 //横坐标
			for m, title := range titles {
				if title == tag {
					col = m
					break
				}
			}

			//获取值
			name := vt.Elem().Field(j).Name //得到结构体内变量名称
			val := curvv.FieldByName(name)  //通过名称得到数值
			fmt.Println(i, name, val)
			pos, _ := excelize.CoordinatesToCellName(col+1, row+i)
			f.SetCellValue(sheet, pos, val)
		}
	}
	return nil
}

//GetExcelData 获取指定表格值
func GetExcelData(path string, sheet string, v interface{}) error {
	vt := reflect.TypeOf(v).Elem()
	vv := reflect.ValueOf(v).Elem()
	if vt.Kind() != reflect.Slice || vt.Elem().Kind() != reflect.Struct {
		return errors.New("info not a slice interface")
	}
	vtt := vt.Elem()
	numTag := vtt.NumField()
	//获取标题的数组
	titles, err := GetExcelTtile(path, sheet)
	if err != nil {
		return err
	}
	if numTag != len(titles) {
		return errors.New("titles num not equal")
	}

	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if _, ok := f.Sheet.Load(sheet); !ok {
		fmt.Println("sheet not exist create", sheet)
		_, err := f.NewSheet(sheet)
		if err != nil {
			fmt.Println(err)
		}
	}
	defer func() {
		if err := f.Save(); err != nil {
			fmt.Println(err)
		}
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, _ := f.GetRows(sheet)
	vv.SetCap(len(rows))

	b := vv
	for i, row := range rows {
		if i > 0 {
			item := reflect.New(vtt)
			for j, row1 := range row {
				tag := titles[j]
				for m := 0; m < vtt.NumField(); m++ {
					if vtt.Field(m).Tag.Get(tagExcel) == tag {
						//招待tag相同的，进行相应的赋值
						name := vtt.Field(m).Name
						valType := vtt.Field(m).Type.Kind()
						//fmt.Println(m, name, valType)
						valp := item.Elem().FieldByName(name)
						switch valType {
						case reflect.String:
							//val:=row1
							valp.SetString(row1)
							//fmt.Println(item)
						case reflect.Float64, reflect.Float32:
							val, _ := strconv.ParseFloat(row1, 10)
							valp.SetFloat(val)
						case reflect.Uint64:
							val, _ := strconv.ParseUint(row1, 10, 64)
							valp.SetUint(val)
						case reflect.Uint32:
							val, _ := strconv.ParseUint(row1, 10, 32)
							valp.SetUint(val)
						case reflect.Int64:
							val, _ := strconv.ParseInt(row1, 10, 64)
							valp.SetInt(val)
						case reflect.Int32:
							val, _ := strconv.ParseInt(row1, 10, 32)
							valp.SetInt(val)
						default:
							fmt.Println("unknown type ", valType)
						}
						break
					}
				}
			}
			//fmt.Println(item)
			//将得到的单元追加到结果
			b = reflect.Append(b, item.Elem())
		}
	}
	//fmt.Println(b)

	//将临时数组传递出去
	if vv.CanSet() {
		vv.Set(b)
	} else {
		return errors.New("v interface can not set")
	}

	return nil
}

func ExcelTest() {
	err := SetExcelTitle("./excel/format/HSB.xlsx", "资产负债表", HSB.BalanceSheet{})
	if err != nil {
		fmt.Println(err)
	}
	titles, err := GetExcelTtile("./excel/format/HSB.xlsx", "资产负债表")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(titles)
	}

	bsu := HSB.BSU{
		{
			SECNAME: "nihao",
			SECCODE: "000001",
			F041N:   1000},
	}
	SetExcelData("./excel/format/HSB.xlsx", "资产负债表", &bsu)
	bsu1 := make([]HSB.BalanceSheet, 0, 100)
	GetExcelData("./excel/format/HSB.xlsx", "资产负债表", &bsu1)
	fmt.Println(bsu1)
}
