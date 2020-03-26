package tools

import (
	"strconv"
	"time"
)

const (
	TenantScope   = "tenant"   //租户后台
	BackendScope  = "backend"  //应用后台
	FrontendScope = "frontend" //应用前台
	WebApp        = "h5"       //h5端
	MpApp         = "mp"       //小程序
	AppNative     = "app"      //原生app
)

func UUID18(prefix string) string {
	return prefix + strconv.Itoa(int(time.Now().UnixNano()))[:18-len(prefix)]
}
