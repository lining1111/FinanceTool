package cninfo

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/golang/glog"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var usr = "wOtj05dLpabViMzo1tIN5VJ2X4krNDfm"
var passwd = "iJSoU1ak7wa3SsW2L2CvIJ8DWQowwhiS"

type TokenInfo struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token,omitempty"`
	Expires_in    int    `json:"expires_in"`
}

type LocalTokenInfo struct {
	tokenInfo TokenInfo
	GetTime   time.Time
	mtx       sync.RWMutex
}

var localToken LocalTokenInfo

var usToken = false

func init() {
	if usToken {

		//判断token是否为空，空的话获取一次
		localToken.mtx.Lock()
		if len(localToken.tokenInfo.Access_token) == 0 {
			token, err := GetToken()
			if err != nil {
				log.Println("本地token为空，获取token 错误%v", err)
			} else {
				localToken.tokenInfo = token
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
					glog.Info("token 过期，重新获取token")
					token, err := GetToken()
					if err != nil {
						glog.Info("获取token 错误%v", err)
					} else {
						localToken.tokenInfo = token
						localToken.GetTime = time.Now()
						glog.Info("获取token:%v,time:%v",
							localToken.tokenInfo, localToken.GetTime.String())
					}
				}
				localToken.mtx.Unlock()
			}
		}()
		wg.Wait()
	}
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

var Mcode = "" //这个是后台验证的，是时间戳的base64

func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	//new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		glog.Error(err)
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
	glog.Info("Go ", http.MethodPost, " URL:", req.URL.String())
	return client.Do(req)
}

func Post(url string, body interface{}, params map[string]string, headers map[string]string) (*http.Response, error) {
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
	client := &http.Client{}
	glog.Info("Go ", http.MethodPost, " URL:", req.URL.String())
	return client.Do(req)
}

func PostWithoutToken(url string, body interface{}, params map[string]string, headers map[string]string) (*http.Response, error) {
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
	//req.Header.Set("Referer","strict-origin-when-cross-origin")
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
	//glog.Info("Go %s URL : %s \n", http.MethodPost, req.URL.String())
	return client.Do(req)
}

func GetToken() (TokenInfo, error) {
	var tokenInfo TokenInfo
	params := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     usr,
		"client_secret": passwd,
	}
	resp, err := PostWithoutToken("http://webapi.cninfo.com.cn/api-cloud-platform/oauth2/token", nil, params, nil)
	if err != nil {
		return tokenInfo, err
	} else {
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			return tokenInfo, err1
		}

		err2 := json.Unmarshal(body, &tokenInfo)
		if err2 != nil {
			return tokenInfo, err2
		} else {
			return tokenInfo, nil
		}

	}
}
