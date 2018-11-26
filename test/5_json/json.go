package main

import (
	"encoding/json"
	"fmt"
	"github.com/luoyelaiwanfan/log/test/5_json/st"
)

////TreeInfo from go-carbon project
//type GlobResponse struct {
//	Name    string       `json:"name,omitempty"`
//	Matches []*GlobMatch `json:"matches,omitempty"`
//}
//
//type GlobMatch struct {
//	Path   string `json:"path,omitempty"`
//	IsLeaf bool   `json:"isLeaf,omitempty"`
//}
//
//type TreeInfo struct {
//	Metrics []*GlobResponse `json:"metrics"`
//}



//TreeInfo from go-carbon project
type GlobResponse struct {
	Name    string       `json:"name,omitempty"`
	Matches []*GlobMatch `json:"matches,omitempty"`
}



// Fake single interval set for graphite
type IntervalSet struct {
	Start int32 `json:"start,omitempty"`
	End   int32 `json:"end,omitempty"`
}
type GlobMatch struct {
	Path   string `json:"path,omitempty"`
	IsLeaf bool   `json:"is_leaf,omitempty"`
	//Intervals []IntervalSet `json:"intervals,omitempty"`
}

type TreeInfo struct {
	Metrics []*GlobResponse `json:"metrics"`
}
func main() {

	//
	//	byt := []byte{123,34,109,101,116,114,105,99,115,34,58,91,123,34,110,97,109,101,34,58,34,42,34,44,34,109,97,116,99,104,101,115,34,58,91,123,34,112,97,116,104,34,58,34,97,97,97,34,125,44,123,34,112,97,116,104,34,58,34,97,98,99,34,125,44,123,34,112,97,116,104,34,58,34,99,97,114,98,111,110,34,125,44,123,34,112,97,116,104,34,58,34,109,121,116,101,115,116,34,125,44,123,34,112,97,116,104,34,58,34,115,116,97,116,115,34,125,93,125,93,125}
	//
	//fmt.Println("bytStr", string(byt))
	byt := []byte(`{"metrics":[{"name":"*","matches":[{"path":"aaa"},{"path":"abc"},{"path":"carbon"},{"path":"mytest"},{"path":"stats"}]}]}`)
	resp := &st.MultiGlobResponse{}
	resp.Metrics = make([]st.GlobResponse, 0)

	//err := resp.Unmarshal(byt)
	err := json.Unmarshal(byt, &resp)

	fmt.Printf("resp:%+v", resp)
	fmt.Println("err:", err)

	//byt1 := []byte{91,123,34,105,115,95,108,101,97,102,34,58,32,102,97,108,115,101,44,32,34,112,97,116,104,34,58,32,34,115,116,97,116,115,46,107,118,34,125,44,32,123,34,105,115,95,108,101,97,102,34,58,32,102,97,108,115,101,44,32,34,112,97,116,104,34,58,32,34,115,116,97,116,115,46,110,111,99,34,125,44,32,123,34,105,115,95,108,101,97,102,34,58,32,102,97,108,115,101,44,32,34,112,97,116,104,34,58,32,34,115,116,97,116,115,46,111,112,101,110,97,112,105,34,125,44,32,123,34,105,115,95,108,101,97,102,34,58,32,102,97,108,115,101,44,32,34,112,97,116,104,34,58,32,34,115,116,97,116,115,46,111,112,101,110,97,112,105,95,112,114,111,102,105,108,101,34,125,44,32,123,34,105,115,95,108,101,97,102,34,58,32,102,97,108,115,101,44,32,34,112,97,116,104,34,58,32,34,115,116,97,116,115,46,116,105,109,101,114,115,34,125,93}

	byt3 := []byte(`[{"is_leaf": false, "path": "stats.kv"}, {"is_leaf": false, "path": "stats.noc"}, {"is_leaf": false, "path": "stats.openapi"}, {"is_leaf": false, "path": "stats.openapi_profile"}, {"is_leaf": false, "path": "stats.timers"}]`)

	//resp1 := &st.MultiGlobResponse{}

	resp3 := make([]GlobMatch,0)


	//resp1.Metrics = make([]st.GlobResponse, 0)
	err = json.Unmarshal(byt3, &resp3)
	fmt.Println("err:", err)
	fmt.Printf("resp3:%+v\n", resp3)


	//byt2 := []byte(`{"name":"mytest.*"}`)
	//resp2 := &st.MultiGlobResponse{}
	//resp2.Metrics = make([]st.GlobResponse, 0)
	//json.Unmarshal(byt2, &resp2)
	//fmt.Printf("resp2:%+v", resp2)

	resp4 := make([]GlobMatch,0)
	byt4 := []byte(`[{"is_leaf": true, "path": "stats.kv.openapi.access.usergraph.aliyun_tc.total._.bytes", "intervals": [{"start": 1507605944.086437, "end": 1539141685.328313}]}, {"is_leaf": true, "path": "stats.kv.openapi.access.usergraph.aliyun_tc.total._.status", "intervals": [{"start": 1507605944.06086, "end": 1539141861.5280707}]}]`)
	err = json.Unmarshal(byt4, &resp4)
	fmt.Println("err:", err)
	fmt.Printf("resp4:%+v\n", resp4)


	resp5:=GlobResponse{}
	//resp5 := make([]GlobMatch,0)
	byt5 := []byte(`{"name":"stats_byhost.----mcached.*","matches":[{"path":"stats_byhost.----mcached.7481"}]}`)
	err = json.Unmarshal(byt5, &resp5)
	fmt.Println("err:", err)
	fmt.Printf("resp5:%+v\n", resp5)
	fmt.Printf("resp5:%+v\n", resp5.Matches[0])


}
