package client

/**
 * 签名方法
这里介绍API请求中签名 ( signature ) 的生成方法。签名需要你先在控制台创建 Access Key ，获得 accesss_key_id 和 secret_access_key，这里我们假设

access_key_id = 'QYACCESSKEYIDEXAMPLE'
secret_access_key = 'SECRETACCESSKEY'
例如我们的请求参数如下:

{
  "count":1,
  "vxnets.1":"vxnet-0",
  "zone":"pek1",
  "instance_type":"small_b",
  "signature_version":1,
  "signature_method":"HmacSHA256",
  "instance_name":"demo",
  "image_id":"centos64x86a",
  "login_mode":"passwd",
  "login_passwd":"QingCloud20130712",
  "version":1,
  "access_key_id":"QYACCESSKEYIDEXAMPLE",
  "action":"RunInstances",
  "time_stamp":"2013-08-27T14:30:10Z"
}
注解 你可以使用上述的 AccessKey 和 Request 调试你的代码， 当得到跟后面一致的签名结果后(即表示你的代码是正确的)， 可再换为你自己的 AccessKey 和其他 API 请求。
签名步骤

1. 按参数名进行升序排列

排序后的参数为:

{
  "access_key_id":"QYACCESSKEYIDEXAMPLE",
  "action":"RunInstances",
  "count":1,
  "image_id":"centos64x86a",
  "instance_name":"demo",
  "instance_type":"small_b",
  "login_mode":"passwd",
  "login_passwd":"QingCloud20130712",
  "signature_method":"HmacSHA256",
  "signature_version":1,
  "time_stamp":"2013-08-27T14:30:10Z",
  "version":1,
  "vxnets.1":"vxnet-0",
  "zone":"pek1"
}
2. 对参数名称和参数值进行URL编码

编码后的请求串为:

{
  "access_key_id":"QYACCESSKEYIDEXAMPLE",
  "action":"RunInstances",
  "count":1,
  "image_id":"centos64x86a",
  "instance_name":"demo",
  "instance_type":"small_b",
  "login_mode":"passwd",
  "login_passwd":"QingCloud20130712",
  "signature_method":"HmacSHA256",
  "signature_version":1,
  "time_stamp":"2013-08-27T14%3A30%3A10Z",
  "version":1,
  "vxnets.1":"vxnet-0",
  "zone":"pek1"
}
3. 构造URL请求

参数名和参数值之间用 “=” 号连接，参数和参数之间用 “＆” 号连接，构造后的URL请求为

access_key_id=QYACCESSKEYIDEXAMPLE&action=RunInstances&count=1&image_id=centos64x86a&instance_name=demo&instance_type=small_b&login_mode=passwd&login_passwd=QingCloud20130712&signature_method=HmacSHA256&signature_version=1&time_stamp=2013-08-27T14%3A30%3A10Z&version=1&vxnets.1=vxnet-0&zone=pek1
4. 构造被签名串

被签名串的构造规则为: 被签名串 = HTTP请求方式 + "\n" + uri + "\n" + url 请求串

注解 注意不要将 “\” 转义
假设 HTTP 请求方法为 GET, 请求的uri路径为 "/iaas/", 则被签名串为

GET\n/iaas/\naccess_key_id=QYACCESSKEYIDEXAMPLE&action=RunInstances&count=1&image_id=centos64x86a&instance_name=demo&instance_type=small_b&login_mode=passwd&login_passwd=QingCloud20130712&signature_method=HmacSHA256&signature_version=1&time_stamp=2013-08-27T14%3A30%3A10Z&version=1&vxnets.1=vxnet-0&zone=pek1
5. 计算签名

计算被签名串的签名 signature。

将API密钥的私钥 ( secret_access_key ) 作为key，生成被签名串的 HMAC-SHA256 或者 HMAC-SHA1 签名，更多信息可参见 RFC2104
将签名进行 Base64 编码
将 Base64 编码后的结果进行 URL 编码 （注意，当 Base64 编码后存在空格时，不要对空格进行 URL 编码，而要直接将空格转为 “+”）
将签名用于参数”signature”的值。
以 Python 代码为例 (其他语言类似，需要使用 sha256 + base64 编码，最后需要再进行 URL 编码，URL 编码时需要将原有的空格 ” ” 转为 “+”)

import base64
import hmac
import urllib
from hashlib import sha256

h = hmac.new(secret_access_key, digestmod=sha256)
h.update("string_to_sign") # 前面生成的被签名串
sign = base64.b64encode(h.digest()).strip()
signature = urllib.quote_plus(sign)
6. 添加签名

将签名参数附在原有请求串的最后面。 最终的HTTP请求串为(为了查看方便，我们人为地将参数之间用回车分隔开)

access_key_id=QYACCESSKEYIDEXAMPLE
&action=RunInstances
&count=1
&image_id=centos64x86a
&instance_name=demo
&instance_type=small_b
&login_mode=passwd
&login_passwd=QingCloud20130712
&signature_method=HmacSHA256
&signature_version=1
&time_stamp=2013-08-27T14%3A30%3A10Z
&version=1
&vxnets.1=vxnet-0
&zone=pek1
&signature=32bseYy39DOlatuewpeuW5vpmW51sD1A%2FJdGynqSpP8%3D
完整的请求URL为(为了查看方便，我们人为地将参数之间用回车分隔开)

https://api.qingcloud.com/iaas/?access_key_id=QYACCESSKEYIDEXAMPLE
&action=RunInstances
&count=1
&image_id=centos64x86a
&instance_name=demo
&instance_type=small_b
&login_mode=passwd
&login_passwd=QingCloud20130712
&signature_method=HmacSHA256
&signature_version=1
&time_stamp=2013-08-27T14%3A30%3A10Z
&version=1
&vxnets.1=vxnet-0
&zone=pek1
&signature=32bseYy39DOlatuewpeuW5vpmW51sD1A%2FJdGynqSpP8%3D
*/
import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type Param struct {
	Name  string
	Value interface{}
}

type Params []*Param

func (ps Params) Len() int {
	return len(ps)
}

func (ps Params) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }

type ParamsSortByName struct{ Params }

func (s ParamsSortByName) Less(i, j int) bool { return s.Params[i].Name < s.Params[j].Name }

func (c Client) GenSignature(params map[string]string) string {
	return ""
}

func sortParamsByKey(ps Params) Params {

	sort.Sort(ParamsSortByName{ps})
	return ps

}

func urlEscapeParams(ps Params) Params {
	for i, _ := range ps {
		if str, ok := ps[i].Value.(string); ok {
			ps[i].Value = url.QueryEscape(str)
		}
	}
	return ps
}

func generateUrlByParams(ps Params) string {
	var urls []string
	for _, v := range ps {
		urls = append(urls, fmt.Sprintf("%v=%v", v.Name, v.Value))
	}
	return strings.Join(urls, "&")
}
func genSignatureUrl(httpMethod string, uri string, url string) string {
	return fmt.Sprintf("%v\n%v\n%v", httpMethod, uri, url)
}

func genSignature(signUrl, secret string) string {

	key := []byte(secret)
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(signUrl))

	sEnc := b64.StdEncoding.EncodeToString(mac.Sum(nil))

	strings.Replace(sEnc, " ", "+", -1)
	fin := url.QueryEscape(sEnc)
	return fin
}
