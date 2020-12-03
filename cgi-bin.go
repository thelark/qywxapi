package qywxapi

import (
	"github.com/thelark/request"
	"fmt"
	"reflect"
)

type cgiBin struct {
	CorpID      string
	Secret      string
	AccessToken string
}

func (t *cgiBin) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------
func (t *cgiBin) Department(opts ...option) *cgiBinDepartment {
	self := &cgiBinDepartment{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}

// 方法 --------------------------------------------------------------------

type cgiBinToken struct {
	*ErrorReturn
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetToken 获取access_token是调用企业微信API接口的第一步，相当于创建了一个登录凭证，其它的业务API接口，都需要依赖于access_token来鉴权调用者身份。因此开发者，在使用业务接口前，要明确access_token的颁发来源，使用正确的access_token。
// GET https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=ID&corpsecret=SECRET
func (t *cgiBin) GetToken() (*cgiBinToken, error) {
	rsp := new(cgiBinToken)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/gettoken", getBasePath()),
		request.WithParam("corpid", t.CorpID),
		request.WithParam("corpsecret", t.Secret),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	t.set("AccessToken", rsp.AccessToken)
	return rsp, nil
}
