package cninfo

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"sync"
	"time"
)

type LocalTokenInfo struct {
	tokenInfo TokenInfo
	GetTime   time.Time
	mtx       sync.RWMutex
}

var localToken LocalTokenInfo

const usToken = false

func tokenFresh() {
	//判断token是否为空，空的话获取一次
	localToken.mtx.Lock()
	if len(localToken.tokenInfo.Access_token) == 0 {
		err := localToken.tokenInfo.Get()
		if err != nil {
			log.Println("本地token为空，获取token 错误:", err)
		} else {
			localToken.GetTime = time.Now()
			log.Println("本地token为空，获取token:", localToken.tokenInfo,
				",time:", localToken.GetTime.String())
		}
	}
	localToken.mtx.Unlock()
	//启动一个线程，如果本地token信息过期的话，刷新token
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		wg.Done()
		for true {
			time.Sleep(time.Duration(60) * time.Second)
			localToken.mtx.Lock()
			//判断是否过期
			du := time.Now().Unix() - localToken.GetTime.Unix()
			if du >= int64(localToken.tokenInfo.Expires_in) {
				log.Println("token 过期，重新获取token")
				err := localToken.tokenInfo.Get()
				if err != nil {
					log.Printf("获取token 错误%v", err)
				} else {
					localToken.GetTime = time.Now()
					log.Printf("获取token:%v,time:%v",
						localToken.tokenInfo, localToken.GetTime.String())
				}
			}
			localToken.mtx.Unlock()
		}
	}()
	wg.Wait()
}

const periodControl = 50

func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	//通过访问频率
	time.Sleep(time.Duration(periodControl) * time.Millisecond)
	//new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		glog.V(1).Info(err)
		return nil, errors.New("new request is fail ")
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	seconds := time.Now().Unix()

	req.AddCookie(&http.Cookie{
		Name:  "Hm_lvt_489bd07e99fbfc5f12cbb4145adb0a9b",
		Value: strconv.Itoa(int(seconds)),
	})

	req.AddCookie(&http.Cookie{
		Name:  "Hm_lpvt_489bd07e99fbfc5f12cbb4145adb0a9b",
		Value: strconv.Itoa(int(seconds + 60)),
	})
	req.Header.Add("Host", "webapi.cninfo.com.cn")
	req.Header.Add("Origin", "http://webapi.cninfo.com.cn")
	req.Header.Add("Referer", "http://webapi.cninfo.com.cn/")
	timestampBase64 := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(seconds))))

	req.Header.Add("mcode", timestampBase64)

	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	if usToken {

		localToken.mtx.RLock()
		token := localToken.tokenInfo.Access_token
		localToken.mtx.RUnlock()
		if token != "" {
			q.Add("access_token", token)
		}
	}
	req.URL.RawQuery = q.Encode()

	//http client
	client := &http.Client{}
	glog.V(1).Info("Go ", http.MethodPost, " URL:", req.URL.String())
	return client.Do(req)
}

func Post(url string, body interface{}, params map[string]string, headers map[string]string) (*http.Response, error) {
	//通过访问频率
	time.Sleep(time.Duration(periodControl) * time.Millisecond)
	//add post body
	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			glog.V(1).Info(err)
			return nil, errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		glog.Error(err)
		return nil, errors.New("new request is fail: %v \n")
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
	}
	if usToken {

		localToken.mtx.RLock()
		token := localToken.tokenInfo.Access_token
		localToken.mtx.RUnlock()
		if token != "" {
			q.Add("access_token", token)
		}
	}
	req.URL.RawQuery = q.Encode()

	seconds := time.Now().Unix()

	req.AddCookie(&http.Cookie{
		Name:  "Hm_lvt_489bd07e99fbfc5f12cbb4145adb0a9b",
		Value: strconv.Itoa(int(seconds)),
	})

	req.AddCookie(&http.Cookie{
		Name:  "Hm_lpvt_489bd07e99fbfc5f12cbb4145adb0a9b",
		Value: strconv.Itoa(int(seconds + 60)),
	})
	req.Header.Add("Host", "webapi.cninfo.com.cn")
	req.Header.Add("Origin", "http://webapi.cninfo.com.cn")
	req.Header.Add("Referer", "http://webapi.cninfo.com.cn/")
	timestampBase64 := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(seconds))))

	req.Header.Add("mcode", timestampBase64)

	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	//http client
	client := &http.Client{
		Timeout: time.Duration(5) * time.Second,
	}
	glog.V(1).Info("Go ", http.MethodPost, " URL:", req.URL.String())
	return client.Do(req)
}

func PostWithoutToken(url string, body interface{}, params map[string]string, headers map[string]string) (*http.Response, error) {
	//通过访问频率
	time.Sleep(time.Duration(periodControl) * time.Millisecond)
	//add post body
	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			glog.Error(err)
			return nil, errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		glog.Error(err)
		return nil, errors.New("new request is fail: %v \n")
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
	}
	req.URL.RawQuery = q.Encode()
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	return client.Do(req)
}

var usr = "wOtj05dLpabViMzo1tIN5VJ2X4krNDfm"
var passwd = "iJSoU1ak7wa3SsW2L2CvIJ8DWQowwhiS"

const APIGetToken = "http://webapi.cninfo.com.cn/api-cloud-platform/oauth2/token"

type TokenInfo struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token,omitempty"`
	Expires_in    int    `json:"expires_in"`
}

func (t *TokenInfo) Get() error {
	params := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     usr,
		"client_secret": passwd,
	}
	resp, err := PostWithoutToken(APIGetToken, nil, params, nil)
	if err != nil {
		return err
	} else {
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			return err1
		}
		return json.Unmarshal(body, t)

	}
}

const MaxResultLimit = 20000

func GetInfoByScodeDate(url string, params map[string]string, info interface{}) error {
	type rspResult struct {
		Resultcode int         `json:"resultcode"`
		Resultmsg  string      `json:"resultmsg"`
		Count      int         `json:"count,omitempty"`
		Total      int         `json:"total,omitempty"`
		Records    interface{} `json:"records,omitempty"`
	}

	if len(url) == 0 {
		return errors.New("url empty")
	}
	//获取类型
	kind := reflect.TypeOf(info).Elem().Kind()
	if kind != reflect.Slice {
		return errors.New("info not a slice interface")
	}
	//获取数组的长度
	capacity := reflect.ValueOf(info).Elem().Cap()

	params["@limit"] = strconv.Itoa(capacity)
	resp, err := Post(url, nil, params, nil)
	if err != nil {
		return err
	} else {
		defer resp.Body.Close()
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			return err1
		} else {
			//打印结果
			//fmt.Println(string(body))
			result := &rspResult{}
			err2 := json.Unmarshal(body, result)
			if err2 != nil {
				return err2
			} else {
				if result.Resultcode != http.StatusOK {
					glog.V(1).Info("result:%s", string(body))
					return errors.New("http req err:" + string(body))
				}

				//打印下数组数量
				glog.V(1).Info(fmt.Sprintf("回复count:%d,total:%d ", result.Count, result.Total))
				if result.Count == 0 {
					return errors.New("http result count 0")
				}
				//反序列化
				kind := reflect.TypeOf(result.Records).Kind()
				if kind != reflect.Slice {
					return errors.New("result records not a slice")
				}
				b, err := json.Marshal(result.Records)
				if err != nil {
					return err
				} else {
					return json.Unmarshal(b, info)

				}
			}
		}
	}
}
