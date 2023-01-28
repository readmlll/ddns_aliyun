package main

import (
	"github.com/alibabacloud-go/tea/tea"
	"strconv"
	"time"
)
import "deep.thinks/common"
import "deep.thinks/api"

func main() {
	jsonText := common.FileUtils{}.ReadTextFileAll("./dns.json")
	jsonText = common.StrUtils{}.Trim(common.StrUtils{}.Replace(jsonText, "\n", "", -1))
	jsonDict := common.EdUtils{}.Str2jsonDict(jsonText)
	accessKeyId := tea.ToString(jsonDict["accessKeyId"])
	accessKeySecret := tea.ToString(jsonDict["accessKeySecret"])
	endPoint := tea.ToString(jsonDict["endPoint"])
	domainPreKey := tea.ToString(jsonDict["domainPreKey"])
	domain := tea.ToString(jsonDict["domain"])
	ttlSecond, _ := strconv.ParseInt(tea.ToString(jsonDict["ttlSecond"]), 10, 64)
	whileSecond, _ := strconv.ParseInt(tea.ToString(jsonDict["whileSecond"]), 10, 64)
	getCurIpCmd := tea.ToString(jsonDict["getCurIpCmd"])
	for true {
		api.AliDnsApi{}.ChangeAliDnsAInfoWarp(accessKeyId, accessKeySecret, endPoint, domainPreKey, domain, ttlSecond, getCurIpCmd)
		time.Sleep(time.Duration(whileSecond) * time.Second)
	}

}
