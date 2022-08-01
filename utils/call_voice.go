package utils

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dyvmsapi"
	)

type VoiceClient struct {
	client *dyvmsapi.Client
}

func NewVoiceClient() *VoiceClient {
	ak := ""
	sk := ""
	client, _ := dyvmsapi.NewClientWithAccessKey("cn-hangzhou", ak, sk)
	return &VoiceClient{
		client: client,
	}
}

func (v *VoiceClient) AppCall(app, callNumber string) string {
	request := dyvmsapi.CreateSingleCallByTtsRequest()
	request.Scheme = "http"
	request.AcceptFormat = "json"

	request.CalledNumber = callNumber
	request.TtsCode = "TTS_242708453"
	request.CalledShowNumber = "057128014906"
	request.TtsParam = fmt.Sprintf("{\"app\":\"%s\"}", app)
	response := dyvmsapi.CreateAddRtcAccountResponse()
	err := v.client.DoAction(request,response)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
	return response.String()
}