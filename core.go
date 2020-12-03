package qywxapi

import (
	"github.com/thelark/request"
	"fmt"
	"path"
	"runtime"
	"strings"
)

const (
	// 通用域名 - 使用该域名将访问官方指定就近的接入点
	qywxAPIURL = "qyapi.weixin.qq.com"
)

var wxRequest = request.New(qywxAPIURL)

type response interface {
	Error() error
}

type ErrorReturn struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (err *ErrorReturn) Error() error {
	if err != nil {
		if err.ErrCode != 0 {
			return fmt.Errorf(err.ErrMsg)
		}
	}
	return nil
}
func checkError(rsp response) error {
	return rsp.Error()
}

type api interface {
	set(k, v string)
}

type option func(api)

// WithCorpID 企业ID参数
func WithCorpID(corpID string) option {
	return func(self api) {
		self.set("CorpID", corpID)
	}
}

// WithSecret 企业密钥
func WithSecret(secret string) option {
	return func(self api) {
		self.set("Secret", secret)
	}
}

// WithAccessToken AccessToken
func WithAccessToken(accessToken string) option {
	return func(self api) {
		self.set("AccessToken", accessToken)
	}
}

// 根据文件名称获取请求路由
func getBasePath() string {
	_, file, _, ok := runtime.Caller(1)
	if ok {
		path := strings.TrimSuffix(path.Base(file), path.Ext(file))
		path = strings.ReplaceAll(path, ".", "/")
		return path
	}
	return ""
}
