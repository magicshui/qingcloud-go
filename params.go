package qingcloud

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

func sortParamsByKey(ps Params) Params {

	sort.Sort(ParamsSortByName{ps})
	return ps

}

func urlEscapeParams(ps Params) Params {
	for i, _ := range ps {
		if str, ok := ps[i].Value.(string); ok {
			ps[i].Value = strings.Replace(url.QueryEscape(str), "+", "%20", -1)
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
