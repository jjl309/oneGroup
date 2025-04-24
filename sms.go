package oneGroup

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func CreateClient() (_result *dysmsapi20170525.Client, _err error) {
	//// 工程代码建议使用更安全的无AK方式，凭据配置方式请参见：https://help.aliyun.com/document_detail/378661.html。
	//credential, _err := credential.NewCredential(nil)
	//if _err != nil {
	//	return _result, _err
	//}

	config := &openapi.Config{
		AccessKeyId:     tea.String(NaCos.Aliyun.AccessKeyID),
		AccessKeySecret: tea.String(NaCos.Aliyun.AccessKeySecret),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func Sms(mobile, code string) (resp *dysmsapi20170525.SendSmsResponse, _err error) {
	client, _err := CreateClient()
	if _err != nil {
		return resp, _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("阿里云短信测试"),
		TemplateCode:  tea.String("SMS_154950909"),
		PhoneNumbers:  tea.String(mobile),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	runtime := &util.RuntimeOptions{}

	resp, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
	if _err != nil {
		return resp, _err
	}

	console.Log(util.ToJSONString(resp))

	return resp, _err

}
