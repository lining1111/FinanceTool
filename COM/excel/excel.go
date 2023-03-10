package excel

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
)

/**
excel操作类
*/

//SetExcelTitle 根据结构体tag设置excel指定工作表的各列标题
func SetExcelTitle(path string, sheet string, v interface{}) error {

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
	return f.SetSheetRow(sheet, "A1", &titles)
}
