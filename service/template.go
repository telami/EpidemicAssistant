package service

import (
	"EpidemicAssistant/model"
	"fmt"
	"strings"
)

func RiskGradeInfoTemplate(result model.RiskGrade) string {

	if !model.IsSuccessful(result) {
		return ""
	}

	data := result.Data

	title := `# 全国中高风险地区一览


<font color="blue">截止至%s，全国疫情</font>：


<font color="red">%d</font>个高风险等级地区；<font color="warning">%d</font>个中风险等级地区


<font color="red">高风险等级地区</font>

`

	topInfo := fmt.Sprintf(title, data.End_update_time, data.Hcount, data.Mcount)

	var highPart strings.Builder
	for _, high := range data.Highlist {
		areaName := "#### " + high.Area_name
		for _, community := range high.Communitys {
			areaName = areaName + "\n" + community + "\n"
		}
		highPart.WriteString(areaName + "\n\n")
	}

	var middlePart strings.Builder
	for _, high := range data.Middlelist {
		areaName := "#### " + high.Area_name
		for _, community := range high.Communitys {
			areaName = areaName + "\n" + community + "\n"
		}
		middlePart.WriteString(areaName + "\n\n")
	}

	topInfo = topInfo + highPart.String() + "<font color=\"warning\">中风险等级地区</font>\n" + middlePart.String()

	return topInfo
}
