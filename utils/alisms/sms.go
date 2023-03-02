package alisms

import (
	"fmt"
	"gin-icqqg/config"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

//阿里云短信发送
var aly = config.AppConfig.AlySms

//SendSmS 阿里云短信发送
//param phone,params string
//return bool
func SendSmS(phone, params string) bool {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", aly.AliYunSmsAk, aly.AliYunSmsAs)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone    //接收短信的手机号码
	request.SignName = aly.SingName //短信签名名称
	request.TemplateCode = aly.Code //短信模板ID
	request.TemplateParam = params  //短信参数 为json字符串
	response, err := client.SendSms(request)
	if err != nil {
		config.ErrorLog(fmt.Sprintf("短信发送失败，错误详情是：%v", err.Error()))
		return false
	}
	if response.Code == "OK" && response.Message == "OK" {
		return true
	} else {
		config.ErrorLog(fmt.Sprintf("短信发送失败，错误详情是：%v", response.Message))
		return false
	}

}
