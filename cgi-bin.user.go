package qywxapi

import (
	"github.com/thelark/request"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type cgiBinUser struct {
	AccessToken string
}

func (t *cgiBinUser) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

type cgiBinUserSimpleList struct {
	*ErrorReturn
	UserList []struct {
		UserID     string  `json:"userid"`      // 成员UserID。对应管理端的帐号
		Name       string  `json:"name"`        // 成员名称，此字段从2019年12月30日起，对新创建第三方应用不再返回真实name，使用userid代替name，2020年6月30日起，对所有历史第三方应用不再返回真实name，使用userid代替name，后续第三方仅通讯录应用可获取，第三方页面需要通过通讯录展示组件来展示名字
		Department []int64 `json:"department"`  // 成员所属部门列表。列表项为部门ID，32位整型
		OpenUserID string  `json:"open_userid"` // 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取
	} `json:"userlist"`
}

// SimpleList 获取部门成员
// GET https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=ACCESS_TOKEN&department_id=DEPARTMENT_ID&fetch_child=FETCH_CHILD
func (t *cgiBinUser) SimpleList(departmentID int64, fetchChild int) (*cgiBinUserSimpleList, error) {
	rsp := new(cgiBinUserSimpleList)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/simplelist", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithParam("department_id", strconv.FormatInt(departmentID, 10)),
		request.WithParam("fetch_child", strconv.Itoa(fetchChild)),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type cgiBinUserList struct {
	*ErrorReturn
	UserList []struct {
		UserID         string  `json:"userid"`            // 成员UserID。对应管理端的帐号
		Name           string  `json:"name"`              // 成员名称；第三方不可获取，调用时返回userid以代替name；对于非第三方创建的成员，第三方通讯录应用也不可获取；第三方页面需要通过通讯录展示组件来展示名字
		Department     []int64 `json:"department"`        // 成员所属部门id列表，仅返回该应用有查看权限的部门id
		Order          []int64 `json:"order"`             // 部门内的排序值，默认为0。数量必须和department一致，数值越大排序越前面。值范围是[0, 2^32)
		Position       string  `json:"position"`          // 职务信息；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
		Mobile         string  `json:"mobile"`            // 手机号码，第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
		Gender         string  `json:"gender"`            // 性别。0表示未定义，1表示男性，2表示女性
		Email          string  `json:"email"`             // 邮箱，第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
		IsLeaderInDept []int64 `json:"is_leader_in_dept"` // 表示在所在的部门内是否为上级。；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
		Avatar         string  `json:"avatar"`            // 头像url。 第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
		ThumbAvatar    string  `json:"thumb_avatar"`      // 头像缩略图url。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
		Telephone      string  `json:"telephone"`         // 座机。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
		Alias          string  `json:"alias"`             // 别名；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
		Status         int     `json:"status"`            // 激活状态: 1=已激活，2=已禁用，4=未激活，5=退出企业。 已激活代表已激活企业微信或已关注微工作台（原企业号）。未激活代表既未激活企业微信又未关注微工作台（原企业号）。
		Address        string  `json:"address"`           // 地址。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
		HideMobile     int     `json:"hide_mobile"`       //
		EnglishName    string  `json:"english_name"`      //
		OpenUserID     string  `json:"open_userid"`       // 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取
		MainDepartment int64   `json:"main_department"`   // 主部门
		Extattr        struct {
			Attrs []struct {
				Type int    `json:"type"`
				Name string `json:"name"`
				Text struct {
					Value string `json:"value"`
				} `json:"text"`
				Web struct {
					URL   string `json:"url"`
					Title string `json:"title"`
				} `json:"web"`
			} `json:"attrs"`
		} `json:"extattr"` // 扩展属性，第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
		QRCode           string `json:"qr_code"`           // 员工个人二维码，扫描可添加为外部联系人(注意返回的是一个url，可在浏览器上打开该url以展示二维码)；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
		ExternalPosition string `json:"external_position"` // 对外职务，如果设置了该值，则以此作为对外展示的职务，否则以position来展示。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
		ExternalProfile  struct {
			ExternalCorpName string `json:"external_corp_name"`
			ExternalAttr     []struct {
				Type int    `json:"type"`
				Name string `json:"name"`
				Text struct {
					Value string `json:"value"`
				} `json:"text"`
				Web struct {
					URL   string `json:"url"`
					Title string `json:"title"`
				} `json:"web"`
				MiniProgram struct {
					AppID    string `json:"appid"`
					PagePath string `json:"pagepath"`
					Title    string `json:"title"`
				} `json:"miniprogram"`
			} `json:"external_attr"`
		} `json:"external_profile"` // 成员对外属性，字段详情见对外属性；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	} `json:"userlist"`
}

// List 获取部门成员详情
// GET https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=ACCESS_TOKEN&department_id=DEPARTMENT_ID&fetch_child=FETCH_CHILD
func (t *cgiBinUser) List(departmentID int64, fetchChild int) (*cgiBinUserList, error) {
	rsp := new(cgiBinUserList)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/list", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithParam("department_id", strconv.FormatInt(departmentID, 10)),
		request.WithParam("fetch_child", strconv.Itoa(fetchChild)),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

// Create 创建成员
// POST https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=ACCESS_TOKEN
func (t *cgiBinUser) Create(body struct {
	UserID         string  `json:"userid"`                      // 必填 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节。只能由数字、字母和“_-@.”四种字符组成，且第一个字符必须是数字或字母。
	Name           string  `json:"name"`                        // 必填 成员名称。长度为1~64个utf8字符
	Department     []int64 `json:"department,omitempty"`        // 成员所属部门id列表,不超过100个
	Order          []int64 `json:"order,omitempty"`             // 部门内的排序值，默认为0，成员次序以创建时间从小到大排列。个数必须和参数department的个数一致，数值越大排序越前面。有效的值范围是[0, 2^32)
	Position       string  `json:"position,omitempty"`          // 职务信息。长度为0~128个字符
	Mobile         string  `json:"mobile"`                      // 必填 手机号码。企业内必须唯一，mobile/email二者不能同时为空
	Gender         string  `json:"gender,omitempty"`            // 性别。1表示男性，2表示女性
	Email          string  `json:"email,omitempty"`             // 邮箱。长度6~64个字节，且为有效的email格式。企业内必须唯一，mobile/email二者不能同时为空
	IsLeaderInDept []int64 `json:"is_leader_in_dept,omitempty"` // 个数必须和参数department的个数一致，表示在所在的部门内是否为上级。1表示为上级，0表示非上级。在审批等应用里可以用来标识上级审批人
	AvatarMediaID  string  `json:"avatar_mediaid,omitempty"`    // 成员头像的mediaid，通过素材管理接口上传图片获得的mediaid
	Telephone      string  `json:"telephone,omitempty"`         // 座机。32字节以内，由纯数字或’-‘号组成。
	Alias          string  `json:"alias,omitempty"`             // 成员别名。长度1~32个utf8字符
	Enable         int     `json:"enable,omitempty"`            // 启用/禁用成员。1表示启用成员，0表示禁用成员
	Address        string  `json:"address,omitempty"`           // 地址。长度最大128个字符
	ToInvite       bool    `json:"to_invite"`                   // 是否邀请该成员使用企业微信（将通过微信服务通知或短信或邮件下发邀请，每天自动下发一次，最多持续3个工作日），默认值为true.
	MainDepartment int64   `json:"main_department,omitempty"`   // 主部门
	Extattr        struct {
		Attrs []struct {
			Type int    `json:"type,omitempty"`
			Name string `json:"name,omitempty"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs,omitempty"`
	} `json:"extattr,omitempty"` // 自定义字段。自定义字段需要先在WEB管理端添加，见扩展属性添加方法，否则忽略未知属性的赋值。与对外属性一致，不过只支持type=0的文本和type=1的网页类型，详细描述查看对外属性
	QRCode           string `json:"qr_code,omitempty"`           //
	ExternalPosition string `json:"external_position,omitempty"` // 对外职务，如果设置了该值，则以此作为对外展示的职务，否则以position来展示。长度12个汉字内
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name,omitempty"`
		ExternalAttr     []struct {
			Type int    `json:"type,omitempty"`
			Name string `json:"name,omitempty"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			MiniProgram struct {
				AppID    string `json:"appid"`
				PagePath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr,omitempty"`
	} `json:"external_profile,omitempty"` // 成员对外属性，字段详情见对外属性
}) (*ErrorReturn, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(ErrorReturn)
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

// Update 更新成员
// POST https://qyapi.weixin.qq.com/cgi-bin/user/update?access_token=ACCESS_TOKEN
func (t *cgiBinUser) Update(body struct {
	UserID         string  `json:"userid"`                      // 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节
	Name           string  `json:"name,omitempty"`              // 成员名称。长度为1~64个utf8字符
	Department     []int64 `json:"department,omitempty"`        // 成员所属部门id列表，不超过100个
	Order          []int64 `json:"order,omitempty"`             // 部门内的排序值，默认为0。数量必须和department一致，数值越大排序越前面。有效的值范围是[0, 2^32)
	Position       string  `json:"position,omitempty"`          // 职务信息。长度为0~128个字符
	Mobile         string  `json:"mobile,omitempty"`            // 手机号码。企业内必须唯一。若成员已激活企业微信，则需成员自行修改（此情况下该参数被忽略，但不会报错）
	Gender         string  `json:"gender,omitempty"`            // 性别。1表示男性，2表示女性
	Email          string  `json:"email,omitempty"`             // 邮箱。长度不超过64个字节，且为有效的email格式。企业内必须唯一。若是绑定了腾讯企业邮箱的企业微信，则需要在腾讯企业邮箱中修改邮箱（此情况下该参数被忽略，但不会报错）
	IsLeaderInDept []int64 `json:"is_leader_in_dept,omitempty"` // 上级字段，个数必须和department一致，表示在所在的部门内是否为上级。
	AvatarMediaID  string  `json:"avatar_mediaid,omitempty"`    // 成员头像的mediaid，通过素材管理接口上传图片获得的mediaid
	Telephone      string  `json:"telephone,omitempty"`         // 座机。由1-32位的纯数字或’-‘号组成
	Alias          string  `json:"alias,omitempty"`             // 别名。长度为1-32个utf8字符
	Enable         int     `json:"enable,omitempty"`            // 启用/禁用成员。1表示启用成员，0表示禁用成员
	Address        string  `json:"address,omitempty"`           // 地址。长度最大128个字符
	MainDepartment int64   `json:"main_department,omitempty"`   // 主部门
	Extattr        struct {
		Attrs []struct {
			Type int    `json:"type,omitempty"`
			Name string `json:"name,omitempty"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs,omitempty"`
	} `json:"extattr,omitempty"` // 自定义字段。自定义字段需要先在WEB管理端添加，见扩展属性添加方法，否则忽略未知属性的赋值。与对外属性一致，不过只支持type=0的文本和type=1的网页类型，详细描述查看对外属性
	ExternalPosition string `json:"external_position,omitempty"` // 对外职务，如果设置了该值，则以此作为对外展示的职务，否则以position来展示。不超过12个汉字
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name,omitempty"`
		ExternalAttr     []struct {
			Type int    `json:"type,omitempty"`
			Name string `json:"name,omitempty"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			MiniProgram struct {
				AppID    string `json:"appid"`
				PagePath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr,omitempty"`
	} `json:"external_profile,omitempty"` // 成员对外属性，字段详情见对外属性
}) (*ErrorReturn, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(ErrorReturn)
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

type cgiBinUserGet struct {
	*ErrorReturn
	UserID         string  `json:"userid"`            // 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节
	Name           string  `json:"name"`              // 成员名称；第三方不可获取，调用时返回userid以代替name；对于非第三方创建的成员，第三方通讯录应用也不可获取；第三方页面需要通过通讯录展示组件来展示名字
	Department     []int64 `json:"department"`        // 成员所属部门id列表，仅返回该应用有查看权限的部门id
	Order          []int64 `json:"order"`             // 部门内的排序值，默认为0。数量必须和department一致，数值越大排序越前面。值范围是[0, 2^32)
	Position       string  `json:"position"`          // 职务信息；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	Mobile         string  `json:"mobile"`            // 手机号码，第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	Gender         string  `json:"gender"`            // 性别。0表示未定义，1表示男性，2表示女性
	Email          string  `json:"email"`             // 邮箱，第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	IsLeaderInDept []int64 `json:"is_leader_in_dept"` // 表示在所在的部门内是否为上级。；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	Avatar         string  `json:"avatar"`            // 头像url。 第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	ThumbAvatar    string  `json:"thumb_avatar"`      // 头像缩略图url。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	Telephone      string  `json:"telephone"`         // 座机。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	Alias          string  `json:"alias"`             // 别名；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	Status         int     `json:"status"`            // 激活状态: 1=已激活，2=已禁用，4=未激活，5=退出企业。已激活代表已激活企业微信或已关注微工作台（原企业号）。未激活代表既未激活企业微信又未关注微工作台（原企业号）。
	Address        string  `json:"address"`           // 地址。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	OpenUserID     string  `json:"open_userid"`       // 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取
	MainDepartment int64   `json:"main_department"`   // 主部门
	Extattr        struct {
		Attrs []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text"`
			Web struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"web"`
		} `json:"attrs"`
	} `json:"extattr"` // 扩展属性，第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	QRCode           string `json:"qr_code"`           // 员工个人二维码，扫描可添加为外部联系人(注意返回的是一个url，可在浏览器上打开该url以展示二维码)；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	ExternalPosition string `json:"external_position"` // 对外职务，如果设置了该值，则以此作为对外展示的职务，否则以position来展示。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name"`
		ExternalAttr     []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text"`
			Web struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"web"`
			MiniProgram struct {
				AppID    string `json:"appid"`
				PagePath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram"`
		} `json:"external_attr"`
	} `json:"external_profile"` // 成员对外属性，字段详情见对外属性；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
}

// Get 读取成员
// GET https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&userid=USERID
func (t *cgiBinUser) Get(userID string) (*cgiBinUserGet, error) {
	rsp := new(cgiBinUserGet)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/get", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithParam("userid", userID),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

// Delete 删除成员
// GET https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=ACCESS_TOKEN&userid=USERID
func (t *cgiBinUser) Delete(userID string) (*ErrorReturn, error) {
	rsp := new(ErrorReturn)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/delete", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithParam("userid", userID),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type cgiBinUserConvertToOpenID struct {
	*ErrorReturn
	OpenID string `json:"openid"`
}

// ConvertToOpenID userid转openid
// POST https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_openid?access_token=ACCESS_TOKEN
func (t *cgiBinUser) ConvertToOpenID(body struct {
	UserID string `json:"userid"`
}) (*cgiBinUserConvertToOpenID, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinUserConvertToOpenID)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/convert_to_openid", getBasePath()),
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

type cgiBinUserConvertToUserID struct {
	*ErrorReturn
	UserID string `json:"userid"`
}

// ConvertToUserID openid转userid
// POST https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_userid?access_token=ACCESS_TOKEN
func (t *cgiBinUser) ConvertToUserID(body struct {
	OpenID string `json:"openid"`
}) (*cgiBinUserConvertToUserID, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinUserConvertToUserID)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/convert_to_userid", getBasePath()),
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
