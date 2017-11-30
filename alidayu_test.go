package alidayu

import "testing"

const (
	accessKeyId     = ""
	accessKeySecret = ""
)

func init() {
	InitAPI(accessKeyId, accessKeySecret)
}

// TestSendSms 测试发送短信
func TestSendSms(t *testing.T) {

	res, err := SendSms("测试", "test", "18628338385", "2", nil, "")
	t.Errorf("返回参数%v, E=%v", res, err)
}

// TestQuerySendDetails 测试短信查询
func TestQuerySendDetails(t *testing.T) {

}
