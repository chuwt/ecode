package ecode

import "net/http"

var (
	OK = NewCode(0, http.StatusOK, "")

	ServerErr = NewCode(500, http.StatusInternalServerError, "服务器错误") // 服务器错误
)
