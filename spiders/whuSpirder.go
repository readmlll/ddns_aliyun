package spiders

import "deep.thinks/common"

type WhuSpider struct {
	BASE_URL string
}

var log = common.LogUtils{}.NewLogUtils("", "", "", "INFO")

func (that WhuSpider) NewWhuSpider() WhuSpider {
	log.Info("测试日志打印情况")
	return WhuSpider{
		BASE_URL: "",
	}
}
