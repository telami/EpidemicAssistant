package main

import (
	"EpidemicAssistant/cron"
	"EpidemicAssistant/model"
	"EpidemicAssistant/service"
	"encoding/json"
)

const provinceTemplate = "确诊人数<font color=\"info\">%d</font>,疑似感染人数<font color=\"info\">%d</font>,治愈人数<font color=\"info\">%d</font>,死亡人数<font color=\"info\">%d</font>"

func main() {
	//result := service.OverAllInfo()
	//log.Info(result)

	//beijing := getProvinceInfo("北京市")
	//tianjin := getProvinceInfo("天津市")
	//
	//bj := fmt.Sprintf(provinceTemplate, beijing.Results[0].ConfirmedCount, beijing.Results[0].SuspectedCount, beijing.Results[0].CuredCount, beijing.Results[0].DeadCount)
	//tj := fmt.Sprintf(provinceTemplate, tianjin.Results[0].ConfirmedCount, tianjin.Results[0].SuspectedCount, tianjin.Results[0].CuredCount, tianjin.Results[0].DeadCount)
	//
	//msgTemplate := `**今日疫情信息同步:** \n
	//> 北京:
	//> ` + bj + `
	//> 天津:
	//> ` + tj + `
	//>[更多数据请查看](https://news.qq.com/zt2020/page/feiyan.htm) \n`
	//
	//log.Info(msgTemplate)
	//riskGradeInfo := getRiskGradeInfo()
	////拼接模板
	//template := service.RiskGradeInfoTemplate(riskGradeInfo)
	////发送到企业微信
	//service.SendInfo(template)

	cron.StartSchedule()
}

func getProvinceInfo(provinceName string) model.Response {
	provinceInfo := service.GetProvinceInfo(provinceName)
	var province model.Response
	_ = json.Unmarshal([]byte(provinceInfo), &province)
	return province
}

//获取高中风险地区信息
func getRiskGradeInfo() model.RiskGrade {
	riskGradeInfo := service.GetRiskGradeInfo()
	var riskGrade model.RiskGrade
	_ = json.Unmarshal([]byte(riskGradeInfo), &riskGrade)
	return riskGrade
}
