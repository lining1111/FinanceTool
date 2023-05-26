package main

import (
	"FinanceTool/analysis"
	"flag"
	"github.com/golang/glog"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 10000, "本地的服务端端口")
	flag.Parse()
	defer glog.Flush()
	glog.Info("欢迎来到金融与go的世界")
	glog.Info(TopCreed)
	glog.Info(Counsel)
	analysis.Testgonum()
	//HSB.BSGetFromCNINFByScode_test()
	//HSB.ISGetFromCNINFByScode_test()
	//HSB.CSGetFromCNINFByScode_test()
	//HSB.IMRGetFromCNINFByScode_test()
	//HSB.AOGetFromCNINFByScode_test()
	//HSB.FRGetFromCNINFByScode_test()
	//cninfo.PublicCodeData_test()
	//Stock.Test1()
	analysis.ResultTest()
}

//func main() {
//	var path string
//	flag.StringVar(&path, "path", "", "要替换的文件")
//	flag.Parse()
//	f, err := os.OpenFile(path, os.O_RDWR, 0666)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer f.Close()
//
//	reader := bufio.NewReader(f)
//	var results []string
//	for {
//		line, _, err := reader.ReadLine()
//		if err == io.EOF {
//			break
//		}
//		if strings.Contains(string(line), "//") &&
//			(strings.Contains(string(line), "float") || strings.Contains(string(line), "uint") || strings.Contains(string(line), "string") || strings.Contains(string(line), "int")) {
//			//需要增加tag的行
//			contents := strings.Split(string(line), " ")
//			for _, v := range contents {
//				if strings.HasPrefix(v, "//") {
//					zhushi := strings.TrimLeft(v, "//")
//					zhushi1:=strings.Split(zhushi," ")
//					str := fmt.Sprintln("`excel:\"", zhushi1[0], "\"`", v)
//					results = append(results,str)
//				}else {
//					results = append(results,v)
//				}
//			}
//		}else{
//			results = append(results,string(line))
//		}
//	}
//	fmt.Println(results)
//}
