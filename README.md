# ddns_aliyun
实现基于阿里云dns api的ddns服务
我采用软路由拨号，光纤宽带上网。
ddns的思路就是检测本地的公网ip同步更改域名的a记录
我使用阿里云的dns服务。使用阿里云dns的api更改并获取最新的a记录信息，使用shell命令获取pppoe最新的公网ip 使得两者保持一致以实现ddns


目录结构
├── DDns
├── DDns.go
├── api
│   └── AliDnsApi.go
├── build_linux.sh
├── common
│   ├── DateUtils.go
│   ├── EdUtils.go
│   ├── FileUtils.go
│   ├── LogUtils.go
│   ├── RequestUtils.go
│   ├── RuntimeUtils.go
│   ├── StrUtils.go
│   └── SysUtils.go
├── dns.json
├── go.mod
├── go.sum


DDns.go 为入口文件
build_linux.sh 为编译linux版本二进制的脚本文件
dns.json 为所需参数

dns.json 需要的参数如下 带Desc的参数可以去除 是说明用的。dns.json需要和编译后的二进制在同一个目录（也可自行更改源码）
{
  "accessKeyId":"阿里云accessKeyId",
  "accessKeySecret": "阿里云accessKeySecret",
  "endPoint": "阿里云dnsapi的访问网关alidns.cn-shenzhen.aliyuncs.com",
  "domainPreKey": "你的域名前缀 如www.baidu.com 就填写www",
  "domain": "你的根域名 如www.baidu.com 就填写baidu.com",
  "ttlSecond": 600,
  "ttlSecondDesc": "ttlSecond 为设置dns的ttl值 阿里云免费的dns默认是10分钟也就是600秒",
  "getCurIpCmd": "ifconfig | grep pppoe -A 1  | grep inet | awk {'print $2'} | awk -F : {'print $2'}",
  "getCurIpCmdDesc": "获取当前公网ip的命令,我是通过pppoe拨号的光纤宽带，所以在软路由上查看网卡信息即可得到公网ip",
  "whileSecond": 180,
  "whileSecondDesc": "设置检测循环的时间，秒为单位，如180则表示 每180秒执行getCurIpCmd设置的命令检测阿里云dns的解析ip是否一致 不一致则重新设置最新的ip"
}




