package qywxapi

import (
	"github.com/thelark/request"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type cgiBinDepartment struct {
	AccessToken string
}

func (t *cgiBinDepartment) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

type cgiBinDepartmentList struct {
	*ErrorReturn
	Department []struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		NameEN   string `json:"name_en"`
		ParentID int64  `json:"parentid"`
		Order    int64  `json:"order"`
	} `json:"department"`
}

// List 获取部门列表
// GET https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=ACCESS_TOKEN&id=ID
func (t *cgiBinDepartment) List(id int64) (*cgiBinDepartmentList, error) {
	rsp := new(cgiBinDepartmentList)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/list", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithParam("id", strconv.FormatInt(id, 10)),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type cgiBinDepartmentCreate struct {
	*ErrorReturn
	ID int64 `json:"id"`
}

// Create 创建部门 注意，部门的最大层级为15层；部门总数不能超过3万个；每个部门下的节点不能超过3万个。建议保证创建的部门和对应部门成员是串行化处理。
// POST https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=ACCESS_TOKEN
func (t *cgiBinDepartment) Create(body struct {
	Name     string `json:"name"`
	NameEN   string `json:"name_en"`
	ParentID int64  `json:"parentid"`
	Order    int64  `json:"order"`
	ID       int64  `json:"id"`
}) (*cgiBinDepartmentCreate, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinDepartmentCreate)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/create", getBasePath()),
		request.WithBody(buf),
		request.WithParam("access_token", t.AccessToken),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type cgiBinDepartmentUpdate struct {
	*ErrorReturn
}

// Update 更新部门
// POST https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=ACCESS_TOKEN
func (t *cgiBinDepartment) Update(body struct {
	Name     string `json:"name,omitempty"`
	NameEN   string `json:"name_en,omitempty"`
	ParentID int64  `json:"parentid,omitempty"`
	Order    int64  `json:"order,omitempty"`
	ID       int64  `json:"id"`
}) (*cgiBinDepartmentCreate, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinDepartmentCreate)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/update", getBasePath()),
		request.WithBody(buf),
		request.WithParam("access_token", t.AccessToken),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type cgiBinDepartmentDelete struct {
	*ErrorReturn
}

// Delete 更新部门
// GET https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=ACCESS_TOKEN&id=ID
func (t *cgiBinDepartment) Delete(id int64) (*cgiBinDepartmentDelete, error) {
	rsp := new(cgiBinDepartmentDelete)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/delete", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithParam("id", strconv.FormatInt(id, 10)),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
