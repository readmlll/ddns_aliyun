package api

import (
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)
import "deep.thinks/common"

var log = common.LogUtils{}.NewLogUtils("", "", "", "INFO")

type AliDnsApi struct {
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func (that AliDnsApi) createClient(accessKeyId *string, accessKeySecret *string, endPoint *string) (_result *alidns20150109.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String(*endPoint)
	_result = &alidns20150109.Client{}
	_result, _err = alidns20150109.NewClient(config)
	return _result, _err
}

func (that AliDnsApi) getDnsARecordInfo(client *alidns20150109.Client, aPreKey string, domain string) (_result *alidns20150109.DescribeDomainRecordsResponse, _err error) {
	describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: tea.String(domain),
		PageNumber: tea.Int64(1),
		PageSize:   tea.Int64(1),
		RRKeyWord:  tea.String(aPreKey),
	}
	runtime := &util.RuntimeOptions{}
	_result = nil
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_result, _err = client.DescribeDomainRecordsWithOptions(describeDomainRecordsRequest, runtime)
		if _err != nil {
			return _err
		}
		return _err
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		var msg *string
		msg, _err = util.AssertAsString(error.Message)
		if _err != nil {

			return nil, _err
		}
		log.Error(*msg)
	}
	return _result, _err
}

func (that AliDnsApi) updateDnsARecordInfo(client *alidns20150109.Client, aPreKey string, recordId string, ip string, ttl int64) (_result *alidns20150109.UpdateDomainRecordResponse, _err error) {
	updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
		RR:       tea.String(aPreKey),
		Type:     tea.String("A"),
		Value:    tea.String(ip),
		RecordId: tea.String(recordId),
		TTL:      tea.Int64(ttl),
	}

	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_result, _err = client.UpdateDomainRecordWithOptions(updateDomainRecordRequest, runtime)
		if _err != nil {
			return _err
		}
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		var msg *string
		msg, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return nil, _err
		}
		log.Error(*msg)
	}
	return _result, _err
}

func (that AliDnsApi) ChangeAliDnsAInfoWarp(accessKeyId string, accessKeySecret string, endPoint string, domainPreKey string, domain string, ttlSecond int64, getCurIpCmd string) (_err error) {
	// 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	changeIp := common.SysUtils{}.CmdExec("sh", "-c", getCurIpCmd)
	changeIp = common.StrUtils{}.Trim(changeIp)

	client, _err := that.createClient(tea.String(accessKeyId), tea.String(accessKeySecret), tea.String(endPoint))
	if _err != nil {
		log.Error("初始化客户端失败")
		return _err
	}
	log.Info("初始化客户端成功")
	var recordInfo *alidns20150109.DescribeDomainRecordsResponse
	recordInfo, _err = that.getDnsARecordInfo(client, domainPreKey, domain)
	if _err != nil {
		log.Error("获取dns记录id失败")
		return _err
	}
	recordId := ""
	recordIp := ""
	var updateResult *alidns20150109.UpdateDomainRecordResponse
	recordId = *recordInfo.Body.DomainRecords.Record[0].RecordId
	recordIp = *recordInfo.Body.DomainRecords.Record[0].Value
	recordIp = common.StrUtils{}.Trim(recordIp)
	log.Info("获取dns记录id成功-当前信息-{" + recordId + "}-{" + recordIp + "}")
	if changeIp == "" {
		log.Error("更改的目标ip为空")
		return _err
	}
	if changeIp == recordIp {
		log.Info("更改的目标ip和记录ip一致无需更改")
		return _err
	}

	updateResult, _err = that.updateDnsARecordInfo(client, domainPreKey, recordId, changeIp, ttlSecond)
	if _err != nil {
		log.Error("更新ip失败")
		return _err
	}
	if updateResult == nil {
		log.Error("更新ip失败，结果为空")
		return _err
	}
	log.Info("更新ip成功，获取最新的信息...")

	recordInfo, _err = that.getDnsARecordInfo(client, domainPreKey, domain)
	if _err != nil {
		log.Error("获取最新dns记录id失败")
		return _err
	}
	recordId = *recordInfo.Body.DomainRecords.Record[0].RecordId
	recordIp = *recordInfo.Body.DomainRecords.Record[0].Value
	log.Info("获取最新dns记录id成功-当前信息-{" + recordId + "}-{" + recordIp + "}")
	return _err
}
