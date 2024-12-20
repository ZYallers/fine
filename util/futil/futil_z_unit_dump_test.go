package futil_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/os/ftime"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/util/futil"
)

func Test_Dump(t *testing.T) {
	type CommonReq struct {
		AppId      int64  `json:"appId" v:"required" in:"path" des:"应用Id" sum:"应用Id Summary"`
		ResourceId string `json:"resourceId" in:"query" des:"资源Id" sum:"资源Id Summary"`
	}
	type SetSpecInfo struct {
		StorageType string   `v:"required|in:CLOUD_PREMIUM,CLOUD_SSD,CLOUD_HSSD" des:"StorageType"`
		Shards      int32    `des:"shards 分片数" sum:"Shards Summary"`
		Params      []string `des:"默认参数(json 串-ClickHouseParams)" sum:"Params Summary"`
	}
	type CreateResourceReq struct {
		CommonReq
		f.Meta    `path:"/CreateResourceReq" method:"POST" tags:"default" sum:"CreateResourceReq sum"`
		Name      string
		CreatedAt *ftime.Time
		SetMap    map[string]*SetSpecInfo
		SetSlice  []SetSpecInfo
		Handler   http.HandlerFunc
		internal  string
	}
	req := &CreateResourceReq{
		CommonReq: CommonReq{
			AppId:      12345678,
			ResourceId: "tdchqy-xxx",
		},
		Name:      "john",
		CreatedAt: ftime.Now(),
		SetMap: map[string]*SetSpecInfo{
			"test1": {
				StorageType: "ssd",
				Shards:      2,
				Params:      []string{"a", "b", "c"},
			},
			"test2": {
				StorageType: "ass",
				Shards:      10,
				Params:      []string{},
			},
		},
		SetSlice: []SetSpecInfo{
			{
				StorageType: "ass",
				Shards:      10,
				Params:      []string{"h"},
			},
		},
	}
	ftest.C(t, func(t *ftest.T) {
		futil.Dump(map[int]int{100: 100})
		futil.Dump(req)
		futil.Dump(true, false)
		futil.Dump(make(chan int))
		futil.Dump(func() {})
		futil.Dump(nil)
	})
}

func Test_DumpTo(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		buffer := bytes.NewBuffer(nil)
		m := f.Map{"k1": f.Map{"k2": "v2"}}
		futil.DumpTo(buffer, m, futil.DumpOption{})
		t.Assert(buffer.String(), `{
    "k1": {
        "k2": "v2",
    },
}`)
	})
}

func Test_DumpWithType(t *testing.T) {
	type CommonReq struct {
		AppId      int64  `json:"appId" v:"required" in:"path" des:"应用Id" sum:"应用Id Summary"`
		ResourceId string `json:"resourceId" in:"query" des:"资源Id" sum:"资源Id Summary"`
	}
	type SetSpecInfo struct {
		StorageType string   `v:"required|in:CLOUD_PREMIUM,CLOUD_SSD,CLOUD_HSSD" des:"StorageType"`
		Shards      int32    `des:"shards 分片数" sum:"Shards Summary"`
		Params      []string `des:"默认参数(json 串-ClickHouseParams)" sum:"Params Summary"`
	}
	type CreateResourceReq struct {
		CommonReq
		f.Meta    `path:"/CreateResourceReq" method:"POST" tags:"default" sum:"CreateResourceReq sum"`
		Name      string
		CreatedAt *ftime.Time
		SetMap    map[string]*SetSpecInfo `v:"required" des:"配置Map"`
		SetSlice  []SetSpecInfo           `v:"required" des:"配置Slice"`
		Handler   http.HandlerFunc
		internal  string
	}
	req := &CreateResourceReq{
		CommonReq: CommonReq{
			AppId:      12345678,
			ResourceId: "tdchqy-xxx",
		},
		Name:      "john",
		CreatedAt: ftime.Now(),
		SetMap: map[string]*SetSpecInfo{
			"test1": {
				StorageType: "ssd",
				Shards:      2,
				Params:      []string{"a", "b", "c"},
			},
			"test2": {
				StorageType: "ass",
				Shards:      10,
				Params:      []string{},
			},
		},
		SetSlice: []SetSpecInfo{
			{
				StorageType: "ass",
				Shards:      10,
				Params:      []string{"h"},
			},
		},
	}
	ftest.C(t, func(t *ftest.T) {
		futil.DumpWithType(map[int]int{100: 100})
		futil.DumpWithType(req)
		futil.DumpWithType([][]byte{[]byte("hello")})
	})
}

func Test_Dump_Slashes(t *testing.T) {
	type Req struct {
		Content string
	}
	req := &Req{
		Content: `{"name":"john", "age":18}`,
	}
	ftest.C(t, func(t *ftest.T) {
		futil.Dump(req)
		futil.Dump(req.Content)
		futil.DumpWithType(req)
		futil.DumpWithType(req.Content)
	})
}

func Test_DumpJson(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var jsonContent = `{"a":1,"b":2}`
		futil.DumpJson(jsonContent)
	})
}
