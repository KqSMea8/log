package main

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

func main() {
	//addr := "http://127.0.0.1:8080/render/?target=stats_byhost.openapi_profile.activity_rpc-yf-docker.DB.m4562i_eos_grid_sina_com_cn_4562.total_count&target=stats_byhost.openapi_profile.activity_rpc-yf-docker.DB.s3601o_apollo_grid_sina_com_cn_3601.slow_count&format=pickle&local=1&noCache=1&from=1515656744&until=1515657344&request_id=711032"
	addr := "/render/?target=stats_byhost.openapi_profile.activity_rpc-yf-docker.DB.m4562i_eos_grid_sina_com_cn_4562.total_count&target=stats_byhost.openapi_profile.activity_rpc-yf-docker.DB.s3601o_apollo_grid_sina_com_cn_3601.slow_count&format=pickle&local=1&noCache=1&from=1515656744&until=1515657344&request_id=711032"

	//addr := "http://10.13.80.70/render/?width=586&height=308&_salt=1537327227.35&from=-24minutes&target=stats_byhost.openapi_profile.activity_rpc-yf-docker.DB.m4562i_eos_grid_sina_com_cn_4562.total_count&target=stats_byhost.openapi_profile.activity_rpc-yf-docker.HTTP.all_http___i_api_card_weibo_com_page_finance_dynamicinfo_json.total_count&target=stats_byhost.openapi_profile.activity_rpc-yf-docker.DB.s3601o_apollo_grid_sina_com_cn_3601.total_count&target=stats_byhost.openapi_profile.activity_rpc-yf-docker.DB.s3601o_apollo_grid_sina_com_cn_3601.slow_count&target=stats_byhost.openapi_profile.activity_rpc-yf-docker.DB.m4562i_eos_grid_sina_com_cn_4562.interval5"
	u, _ := url.Parse(addr)
	//type URL struct {
	//	Scheme     string
	//	Opaque     string    // encoded opaque data
	//	User       *Userinfo // username and password information
	//	Host       string    // host or host:port
	//	Path       string    // path (relative paths may omit leading slash)
	//	RawPath    string    // encoded path hint (see EscapedPath method)
	//	ForceQuery bool      // append a query ('?') even if RawQuery is empty
	//	RawQuery   string    // encoded query values, without '?'
	//	Fragment   string    // fragment for references, without '#'
	//}
	fmt.Println("u:", u.Scheme)
	fmt.Println("Opaque",u.Opaque)
	fmt.Println("User",u.User)
	fmt.Println("host",u.Host)
	fmt.Println("path",u.Path)
	fmt.Println("RawPath",u.RawPath)
	fmt.Println("ForceQuery",u.ForceQuery)
	fmt.Println("RawQuery",u.RawQuery)
	fmt.Println("Fragment",u.Fragment)

	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("m:", m)

	str1 := "stats_byhost.dpool_sso.security_weibo_com_yf2"

	str2 := "stats_byhost.dpool_sso.security_weibo_com_yf"

	num := strings.Index(str1, str2)

	fmt.Println("num:", num)

	fmt.Println(time.Now().Unix())



	uu := url.Values{}
	uu.Set("m", "avg:720s-avg:noc-merge.stats_byhost.openapi_profile.activity_rpc-yf-docker.DB.__.total_count{api=m4562i_eos_grid_sina_com_cn_4563}")
	uu.Set("start", "1515656880")
	uu.Set("end", "1515657600")
	uu.Set("request_id", "unknow")
	fmt.Println("uu:", uu.Encode())

}
