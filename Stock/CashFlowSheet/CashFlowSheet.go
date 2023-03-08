package CashFlowSheet

import (
	"FinanceTool/COM/cninfo"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/goinggo/mapstructure"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

/**
现金流量表，数据类计算
*/

const urlApi = "http://webapi.cninfo.com.cn/api/stock/p_stock2302"

//CashFlowSheet 现金流量表 "http://webapi.cninfo.com.cn/api/stock/p_stock2302"
type CashFlowSheet struct {
	SECNAME     string //证券简称
	SECCODE     string //证券代码
	ORGNAME     string //机构名称
	DECLAREDATE string //公告日期
	ENDDATE     string //截止日期
	F001D       string //报告年度
	F002V       string //合并类型编码
	F003V       string //合并类型
	F004V       string //报表来源编码
	F005V       string //报表来源

	F006N float64 //销售商品、提供劳务收到的现金
	F072N float64 //客户存款和同业存放款项净增加额
	F073N float64 //向中央银行借款净增加额
	F074N float64 //向其他金融机构拆入资金净增加额
	F077N float64 //收到原保险合同保费取得的现金
	F078N float64 //收到再保险业务现金净额
	F079N float64 //保户储金及投资款净增加额
	F080N float64 //处置以公允价值计量且其变动计入当期损益的金融资产净增加额
	F081N float64 //收取利息、手续费及佣金的现金
	F082N float64 //拆入资金净增加额
	F083N float64 //回购业务资金净增加额
	F007N float64 //收到的税费返还
	F008N float64 //收到其他与经营活动有关的现金
	F009N float64 //经营活动现金流入小计

	F010N float64 //购买商品、接受劳务支付的现金
	F084N float64 //客户贷款及垫款净增加额
	F085N float64 //存放中央银行和同业款项净增加额
	F086N float64 //支付原保险合同赔付款项的现金
	F087N float64 //支付利息、手续费及佣金的现金
	F088N float64 //支付保单红利的现金
	F011N float64 //支付给职工以及为职工支付的现金
	F012N float64 //支付的各项税费
	F013N float64 //支付其他与经营活动有关的现金
	F014N float64 //经营活动现金流出小计

	F015N float64 //经营活动产生的现金流量净额

	F016N float64 //收回投资收到的现金
	F017N float64 //取得投资收益收到的现金
	F018N float64 //处置固定资产、无形资产和其他长期资产收回的现金净额
	F019N float64 //处置子公司及其他营业单位收到的现金净额
	F020N float64 //收到其他与投资活动有关的现金
	F021N float64 //投资活动现金流入小计

	F022N float64 //购建固定资产、无形资产和其他长期资产支付的现金
	F023N float64 //投资支付的现金
	F075N float64 //质押贷款净增加额
	F024N float64 //取得子公司及其他营业单位支付的现金净额
	F025N float64 //支付其他与投资活动有关的现金
	F026N float64 //投资活动现金流出小计

	F027N float64 //投资活动产生的现金流量净额

	F028N float64 //吸收投资收到的现金
	F089N float64 //其中：子公司吸收少数股东投资收到的现金
	F029N float64 //取得借款收到的现金
	F076N float64 //发行债券收到的现金
	F030N float64 //收到其他与筹资活动有关的现金
	F031N float64 //筹资活动现金流入小计

	F032N float64 //偿还债务支付的现金
	F033N float64 //分配股利、利润或偿付利息支付的现金
	F090N float64 //其中：子公司支付给少数股东的股利、利润
	F034N float64 //支付其他与筹资活动有关的现金
	F035N float64 //筹资活动现金流出小计

	F036N float64 //筹资活动产生的现金流量净额

	F037N float64 //四、汇率变动对现金的影响
	F038N float64 //四(2)、其他原因对现金的影响

	F039N float64 //五、现金及现金等价物净增加额
	F040N float64 //期初现金及现金等价物余额
	F041N float64 //期末现金及现金等价物余额
	F044N float64 //净利润

	F045N float64 //加：资产减值准备
	F046N float64 //固定资产折旧、油气资产折耗、生产性生物资产折旧
	F091N float64 //投资性房地产的折旧及摊销
	F047N float64 //无形资产摊销
	F048N float64 //长期待摊费用摊销
	F049N float64 //处置固定资产、无形资产和其他长期资产的损失
	F050N float64 //固定资产报废损失
	F051N float64 //公允价值变动损失
	F052N float64 //财务费用
	F053N float64 //投资损失
	F054N float64 //递延所得税资产减少
	F055N float64 //递延所得税负债增加
	F056N float64 //存货的减少
	F057N float64 //经营性应收项目的减少
	F058N float64 //经营性应付项目的增加

	F059N float64 //其他

	F060N float64 //经营活动产生的现金流量净额2
	F062N float64 //债务转为资本
	F063N float64 //一年内到期的可转换公司债券
	F064N float64 //融资租入固定资产
	F066N float64 //现金的期末余额
	F067N float64 //减：现金的期初余额
	F068N float64 //加：现金等价物的期末余额
	F069N float64 //减：现金等价物的期初余额
	F070N float64 //加：其他原因对现金的影响2
	F071N float64 //现金及现金等价物净增加额2
	F096N float64 //信用减值损失
}

//GetFromCNINFByScode 从巨潮资讯平台以股票代码和日期为输入获取
func (cs *CashFlowSheet) GetFromCNINFByScode(scode string, year string, q string) (bool, error) {
	url := urlApi
	params := map[string]string{
		"scode": scode,
		"rdate": cninfo.Getrdate(year, q),
		"type":  "071001",
	}
	resp, err := cninfo.Post(url, nil, params, nil)
	defer resp.Body.Close()
	if err != nil {
		return false, err
	} else {
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			return false, err1
		} else {
			//打印结果
			//fmt.Println(string(body))
			var result map[string]interface{}
			err2 := json.Unmarshal(body, &result)
			if err2 != nil {
				return false, err2
			} else {
				if result["resultcode"].(float64) != http.StatusOK {
					glog.Error("result:%s", string(body))
					return false, errors.New("http req err:" + string(body))
				}
				//打印下数组数量
				fmt.Println("回复结果数为 ", result["total"])
				if result["total"].(float64) == 0 {
					return false, errors.New("http result total 0")
				}
				//反序列化
				err3 := mapstructure.Decode(result["records"].([]interface{})[0], cs)
				if err3 != nil {
					return false, err2
				}
			}
		}
	}
	return true, nil
}

func GetFromCNINFByScode_test() {
	cs := new(CashFlowSheet)
	_, err := cs.GetFromCNINFByScode("000001", "2021", cninfo.Q1)
	if err != nil {
		fmt.Println(err)
	}
}