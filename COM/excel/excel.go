package excel

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
)

/**
excel操作类
*/

//SetExcelTitle 根据结构体tag设置excel指定工作表的各列标题
func SetExcelTitle(path string, sheet string, v interface{}) error {
	const start = "A1"
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return err
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
		tag := reflect.TypeOf(v).Field(i).Tag.Get("excel")
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
		return nil, errors.New("titiles empty")
	}

	return titles, nil
}

//SetExcelData 根据结构体tag设置excel指定工作表的各列标题
func SetExcelData(path string, sheet string, v interface{}) error {
	kind1 := reflect.TypeOf(v).Elem().Kind()
	len1 := reflect.ValueOf(v).Elem().Len()
	if kind1 != reflect.Slice || len1 <= 0 {
		return errors.New("info not a slice interface or len 0")
	}

	//fmt.Println(reflect.TypeOf(v).Elem())
	numTag := reflect.TypeOf(v).Elem().Elem().NumField()
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
	defer func() {
		if err := f.Save(); err != nil {
			fmt.Println(err)
		}
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	row := 2

	for m := 0; m < len1; m++ {
		//elemt:=reflect.TypeOf(v).Elem()
		//fmt.Println(elemt)
		//elemv:=reflect.ValueOf(v).Elem()
		//fmt.Println(elemv)
		kind2 := reflect.ValueOf(v).Elem().Index(m).Kind()
		//fmt.Println(kind2)
		if kind2 == reflect.Struct {
			curStructt := reflect.TypeOf(v).Elem().Elem()
			curStructv := reflect.ValueOf(v).Elem().Index(m)
			//获取v的tag
			for i := 0; i < curStructt.NumField(); i++ {
				//获取tag值
				tag := curStructt.Field(i).Tag.Get("excel")
				col := 0 //横坐标
				for i2, title := range titles {
					if title == tag {
						col = i2
						break
					}
				}

				//获取值
				name := curStructt.Field(i).Name    //得到结构体内变量名称
				val := curStructv.FieldByName(name) //通过名称得到数值
				fmt.Println(name, val)
				pos, _ := excelize.CoordinatesToCellName(col+1, row+m)
				f.SetCellValue(sheet, pos, val)
			}
		}
	}
	return nil
}

//GetExcelData 获取指定表格值
func GetExcelData(path string, sheet string, v interface{}) error {
	resultv := reflect.ValueOf(v).Elem()
	fmt.Println(resultv)
	resultt := reflect.TypeOf(v).Elem()
	kind1 := resultt.Kind()
	cap1 := resultv.Cap()
	if kind1 != reflect.Slice || cap1 <= 0 {
		return errors.New("info not a slice interface or cap 0")
	}

	//fmt.Println(reflect.TypeOf(v).Elem())
	numTag := reflect.TypeOf(v).Elem().Elem().NumField()
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
	defer func() {
		if err := f.Save(); err != nil {
			fmt.Println(err)
		}
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	elemt := reflect.TypeOf(v).Elem().Elem()
	rows, _ := f.GetRows(sheet)
	var b reflect.Value
	for i, row := range rows {
		if i > 0 {
			item := reflect.New(elemt)
			fmt.Println(item)
			for i1, row1 := range row {
				fmt.Println(row1)
				tag := titles[i1]
				//fmt.Println(elemt)
				for i2 := 0; i2 < elemt.NumField(); i2++ {
					if elemt.Field(i2).Tag.Get("excel") == tag {
						//招待tag相同的，进行相应的赋值
						name := elemt.Field(i2).Name
						fmt.Println(name)
						valType := elemt.Field(i2).Type.String()
						fmt.Println(valType)
						switch valType {
						case "string":
							//val:=row1
							item.Elem().Field(i2).SetString(row1)
							//fmt.Println(item)
						case "float64", "float32":
							val, _ := strconv.ParseFloat(row1, 10)
							item.Elem().Field(i2).SetFloat(val)
						case "uint64":
							val, _ := strconv.ParseUint(row1, 10, 64)
							item.Elem().Field(i2).SetUint(val)
						case "uint32":
							val, _ := strconv.ParseUint(row1, 10, 32)
							item.Elem().Field(i2).SetUint(val)
						case "int64":
							val, _ := strconv.ParseInt(row1, 10, 64)
							item.Elem().Field(i2).SetInt(val)
						case "int32":
							val, _ := strconv.ParseInt(row1, 10, 32)
							item.Elem().Field(i2).SetInt(val)

						}
						break
					}
				}
			}
			fmt.Println(item)
			//将得到的单元追加到结果
			b = reflect.ValueOf(v).Elem()
			b = reflect.Append(b, item.Elem())
			fmt.Println(b)
		}
	}

	//将临时数组传递出去
	if reflect.ValueOf(v).Elem().CanSet() {
		reflect.ValueOf(v).Elem().Set(b)
	}

	return nil
}
