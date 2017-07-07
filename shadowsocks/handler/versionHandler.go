package handler

import (
	"net/http"

	"github.com/andy-zhangtao/Sandstorm"
)

// GetVersion 获取当前版本信息
func GetVersion(w http.ResponseWriter, r *http.Request) {

	dv := "Dev Version: " + "r53+5M 947d446"
	rv := "  Release Version: 0.2"
	Sandstorm.HTTPSuccess(w, dv+rv)
}
