package alidayu

var (
	// AccessKeyId，请参阅<a href="https://ak-console.aliyun.com/">阿里云Access Key管理</a>
	AccessKeyId string
	// AccessKeySecret，请参阅<a href="https://ak-console.aliyun.com/">阿里云Access Key管理</a>
	AccessKeySecret string
)

const (
	host = "dysmsapi.aliyuncs.com"
)

// Response 请求返回结果
type Response struct {
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
 * sendSms 发送短信
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
func sendSms(signName, templateCode, phoneNumbers string, templateParam map[string]interface{}, outId, smsUpExtendCode string) {

}

/**
 * queryDetails 短信发送记录查询
 *
 * @param string phoneNumbers 必填, 短信接收号码 (e.g. 12345678901)
 * @param string sendDate 必填，短信发送日期，格式Ymd，支持近30天记录查询 (e.g. 20170710)
 * @param int32 pageSize 必填，分页大小
 * @param int32 currentPage 必填，当前页码
 * @param string bizId 选填，短信发送流水号 (e.g. abc123)
 * @return
 */
func queryDetails(phoneNumbers, sendDate string, pageSize, currentPage int32, bizId string) {

}
