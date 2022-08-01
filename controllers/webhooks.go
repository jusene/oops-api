package controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"gitlab.hho-inc.com/devops/oops-api/models"
	"gitlab.hho-inc.com/devops/oops-api/utils"
	"io/ioutil"
	"strings"
	"text/template"
	"time"
)

// @summary alertmanager 钉钉机器人webhook
// @description 发送dingding机器人
// @tags webhook
// @produce json
// @accept json
// @param token query string true "token"
// @param template query string true "模板"
// @success 200 {object} models.Fail
// @failure 404 {object} models.Success
// @router /webhook/dingding [post]
func WebHookSendDing(c *gin.Context) {
	// 获取alertmanager hook的信息
	v := viper.New()
	v.SetConfigType("json")
	v.ReadConfig(c.Request.Body)
	alerts := v.Get("alerts")

	// 获取钉钉token
	token, ok := c.GetQuery("token")
	if  !ok {
		CheckError(c, 404, "未获取token")
		return
	}

	// 获取consul上的告警模板
	tmp, ok := c.GetQuery("template")
	if !ok {
		CheckError(c, 404, "未获取template")
		return
	}

	// 将模板渲染成告警信息
	consul_client := utils.NewConsul("/devops/alarm/template")
	t, err := consul_client.GetKV(tmp)
	if err != nil {
		CheckError(c, 500, "模板不存在")
		return
	}

	for _, alert := range alerts.([]interface{}) {
		title := alert.(map[string]interface{})["labels"].(map[string]interface{})["alertname"]
		instance := alert.(map[string]interface{})["labels"].(map[string]interface{})["instance"]
		status := alert.(map[string]interface{})["status"]
		severity := alert.(map[string]interface{})["labels"].(map[string]interface{})["severity"]
		job := alert.(map[string]interface{})["labels"].(map[string]interface{})["job"]
		alarmAttr := &models.AlertManager{
			TITLE:    title.(string),
			INSTANCE: instance.(string),
			STATUS:   status.(string),
			SEVERITY: severity.(string),
			JOB:      job.(string),
			TIME:     time.Now().Format("2006-01-02 15:04:05"),
		}
		var buf bytes.Buffer
		alertTemp, _ := template.New("alert").Parse(string(t))
		alertTemp.Execute(&buf, alarmAttr)
		// 发送dingding告警
		ding_client := utils.NewSendDing(token, buf.Bytes())
		_, err := ding_client.Send()

		if err != nil {
			fmt.Println(instance, " send fail!")
			c.JSON(500, models.Fail{
				Code: 500,
				Msg:  err.Error(),
			})
			return
		}
		fmt.Println(instance, " send ok")
	}

	c.JSON(200, gin.H{
		"Msg": "send ok",
	})
}


func CheckError(c *gin.Context, code int, msg string) {
	c.JSON(code, models.Fail{
		Code: code,
		Msg: msg,
	})
}

// @summary alertmanager aliyun语音服务webhook接口
// @description alertmanager 打电话告警
// @tags webhook
// @produce json
// @accept json
// @success 200 {object} models.Fail
// @failure 404 {object} models.Success
// @router /webhook/callvoice [post]
func WebHookCallVoice(c *gin.Context)  {
	// 获取alertmanager hook的信息
	v := viper.New()
	v.SetConfigType("json")
	req, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Println(string(req))
	v.ReadConfig(bytes.NewReader(req))
	instance := strings.Split(v.GetString("alerts.0.labels.instance"), ":")[0]

	// 获取告警list
	consul_client := utils.NewConsul("/devops/alarm")
	alarmList, _ := consul_client.GetKV("alarm_list")
	v.ReadConfig(bytes.NewReader(alarmList))
	list := v.Get("node").([]interface{})

	// 将告警的ip与主机名进行映射
	consul_data, _, _:= consul_client.Client.Catalog().Service("node-exporter", "", &api.QueryOptions{})
	for _, s := range consul_data {
		if s.ServiceAddress == instance {
			instance = s.ServiceID
		}
	}

	// 开始电话告警
	voiceClient := utils.NewVoiceClient()

	for _, u := range list {
		fmt.Println("call:",u)
		resp := voiceClient.AppCall(instance, u.(string))
		v.ReadConfig(bytes.NewReader([]byte(resp)))
		c.JSON(200, v.AllSettings())
	}
}

// @summary loki aliyun语音服务webhook接口
// @description loki 打电话告警
// @tags webhook
// @produce json
// @accept json
// @success 200 {object} models.Fail
// @failure 404 {object} models.Success
// @router /webhook/loki/callvoice [post]
func WebHookLokiCallVoice(c *gin.Context)  {
	// 获取alertmanager hook的信息
	v := viper.New()
	v.SetConfigType("json")
	data, _ := ioutil.ReadAll(c.Request.Body)
	v.ReadConfig(bytes.NewReader(data))
	alerts := v.Get("alerts").([]interface{})
	receiver := v.GetString("receiver")
	fmt.Println(receiver)

	// 获取告警list
	consul_client := utils.NewConsul("/devops/alarm")
	alarmList, _ := consul_client.GetKV("alarm_list")
	v.ReadConfig(bytes.NewReader(alarmList))
	list := v.Get(receiver).([]interface{})

	// 开始电话告警
	voiceClient := utils.NewVoiceClient()

	for _, a := range alerts {
		if a.(map[string]interface{})["status"].(string) == "firing" {
			info := a.(map[string]interface{})["labels"].(map[string]interface{})["alertname"].(string)
			for _, u := range list {
				fmt.Println("call:",u)
				resp := voiceClient.AppCall(info, u.(string))
				v.ReadConfig(bytes.NewReader([]byte(resp)))
				c.JSON(200, v.AllSettings())
			}
		}
	}
}