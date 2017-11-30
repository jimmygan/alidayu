package alidayu

import (
	"net/url"
	"math/rand"
	"time"
	"encoding/json"
	"strings"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/base64"
	"fmt"
	"net/http"
	"net"
	"errors"
)

var (
	// AccessKeyId，请参阅<a href="https://ak-console.aliyun.com/">阿里云Access Key管理</a>
	AccessKeyId string
	// AccessKeySecret，请参阅<a href="https://ak-console.aliyun.com/">阿里云Access Key管理</a>
	AccessKeySecret string
	// 签名算法
	SignatureMethod = "HMAC-SHA1"
	// 返回数据类型
	Format = "JSON"
	// 签名版本
	SignatureVersion = "1.0"
	// API支持的RegionID，如短信API的值为：cn-hangzhou
	RegionId = "cn-hangzhou"
)

const (
	host = "dysmsapi.aliyuncs.com"
)

// Result 请求返回结果
type Result struct {
	RequestId         string           `json:"RequestId"`
	Code              string           `json:"Code"`
	Message           string           `json:"Message"`
	BizId             string           `json:"BizId"`
	TotalCount        int32            `json:"TotalCount"`
	TotalPage         int32            `json:"TotalPage"`
	SmsSendDetailDTOs []*SmsSendDetail `json:"smsSendDetailDTOs"`
}

// SmsSendDetail 短信详情
type SmsSendDetail struct {
	PhoneNum     string `json:"phoneNum"`
	SendStatus   int    `json:"sendStatus"`
	ErrCode      string `json:"errCode"`
	TemplateCode string `json:"templateCode"`
	Content      string `json:"content"`
	SendDate     string `json:"sendDate"`
	ReceiveDate  string `json:"receiveDate"`
	OutId        string `json:"outId"`
}

// InitAPI 初始化api
func InitAPI(accessKeyId, accessKeySecret string) {
	AccessKeyId = accessKeyId
	AccessKeySecret = accessKeySecret
}

/**
 * SendSms 发送短信
 *
 * @param string signName
 * 		必填, 短信签名，应严格"签名名称"填写，参考：<a href="https://dysms.console.aliyun.com/dysms.htm#/sign">短信签名页</a>
 * @param string templateCode
 * 		必填, 短信模板Code，应严格按"模板CODE"填写, 参考：<a href="https://dysms.console.aliyun.com/dysms.htm#/template">短信模板页</a>
 * (e.g. SMS_0001)
 * @param string phoneNumbers 必填, 短信接收号码 (e.g. 12345678901,13456789011)
 * @param map|nil templateParam
 *    选填, 假如模板中存在变量需要替换则为必填项 (e.g. map("code"=>"12345", "product"=>"阿里通信"))
 * @param string|nil outId [optional] 选填, 发送短信流水号 (e.g. 1234)
 * @param string|nil smsUpExtendCode [optional] 选填，上行短信扩展码（扩展码字段控制在7位或以下，无特殊需求用户请忽略此字段）
 * @return
 */
func SendSms(signName, templateCode, phoneNumbers, outId string, templateParam map[string]interface{}, smsUpExtendCode string) (*Result, error) {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	params := url.Values{}
	params.Add("Timestamp", timestamp)
	params.Add("Action", "SendSms")
	params.Add("SignName", signName)
	params.Add("TemplateCode", templateCode)
	params.Add("PhoneNumbers", phoneNumbers)
	params.Add("Version", "2017-05-25")

	if outId != "" {
		params.Add("OutId", outId)
	}
	if templateParam != nil {
		body, err := json.Marshal(templateParam)
		if err != nil {
			return nil, err
		}
		params.Add("TemplateParam", string(body))
	}

	if smsUpExtendCode != "" {
		params.Add("SmsUpExtendCode", smsUpExtendCode)
	}

	resp, err := Request(params)
	fmt.Println(resp)
	return resp, err
}


 // QuerySendDetails 短信发送记录查询
 // @param string phoneNumbers 必填, 短信接收号码 (e.g. 12345678901)
 // @param string sendDate 必填，短信发送日期，格式Ymd，支持近30天记录查询 (e.g. 20170710)
 // @param int32 pageSize 必填，分页大小
 // @param int32 currentPage 必填，当前页码
 // @param string bizId 选填，短信发送流水号 (e.g. abc123)
 // @return
func QuerySendDetails(phoneNumbers, sendDate, bizId string, pageSize, currentPage int32) {

}

// Sign 签名
func Sign(method string,params url.Values) url.Values {
	params.Add("SignatureMethod", SignatureMethod)
	params.Add("AccessKeyId", AccessKeyId)
	params.Add("SignatureVersion", SignatureVersion)
	params.Add("Format", Format)
	params.Add("SignatureNonce", RandString(16))

	// hmac签名
	data := params.Encode()
	data = strings.ToUpper(method)+ "&/&" + data
	data = SpecialUrlEncode(data)
	// accessSecret：你的AccessKeyId对应的秘钥AccessSecret，特别说明：POP要求需要后面多加一个“&”字符，即accessSecret + “&”
	sign := HmacSha1(data, AccessKeySecret + "&")
	sign = SpecialUrlEncode(sign)

	params.Add("Signature", sign)
	return params
}

// Request 向网关发起请求
func Request(params url.Values) (*Result, error) {
	resp := new(Result)
	method := "GET"
	params = Sign(method, params)
	urlstring := "http://" + host + "/?" + params.Encode()
	req := http.NewRequest(method, urlstring, nil)

	client := http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				deadline := time.Now().Add(25 * time.Second)
				c, err := net.DialTimeout(network, addr, time.Second * 20)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	// 请求成功
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("http请求失败")
	}

	return resp, nil
}

// HmacSha1签名算法
func HmacSha1(data string, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// SpecialUrlEncode 替换特殊字符串
func SpecialUrlEncode(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}

// RandString 生成随机字符串
func RandString(length int) string {
	key := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = key[b%byte(len(key))]
	}
	return string(bytes)
}
