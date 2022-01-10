package req

import "github.com/gin-gonic/gin"

//IamToken 用户token
type IamToken struct {
	ctx               *gin.Context
	SubAuthUser       *SubAuthUser     `json:"subAuthUser"`
	PrimaryAuthUser   *PrimaryAuthUser `json:"primaryAuthUser"`
	IsPrimaryAuthUser bool             `json:"isPrimaryAuthUser"`
	SubUserGroups     []Groups         `json:"subUserGroups"`
	Token             string           `header:"token"`
}

func (iamToken *IamToken) GetCtx() *gin.Context {
	return iamToken.ctx
}

//SetAccount 从用户请求中将token解析到iamToken中
//noinspection ALL
func (iamToken *IamToken) SetCtx(ctx *gin.Context) {
	iamToken.ctx = ctx
	tokenInterfaceValue, _ := ctx.Get("token")
	if tokenInterfaceValue != nil {
		token, ok := tokenInterfaceValue.(IamToken)
		if ok {
			iamToken.SubUserGroups = token.SubUserGroups
			iamToken.PrimaryAuthUser = token.PrimaryAuthUser
			iamToken.IsPrimaryAuthUser = token.IsPrimaryAuthUser
			iamToken.SubAuthUser = token.SubAuthUser
			iamToken.Token = token.Token
		}
	}
}

//GetHarborPwd 获取harbor密码
func (iamToken *IamToken) GetHarborPwd() string {
	if iamToken.PrimaryAuthUser != nil {
		return iamToken.PrimaryAuthUser.HarborPassword
	}
	return ""
}

//GetUserName 获取用户名称
func (iamToken *IamToken) GetUserName() string {
	if iamToken.IsPrimaryAuthUser {
		return iamToken.GetPrimaryAuthUserName()
	}
	return iamToken.GetSubAuthUserName()
}

//GetSubAuthUserName 获取子账号名称
func (iamToken *IamToken) GetSubAuthUserName() string {
	if iamToken.SubAuthUser != nil {
		return iamToken.SubAuthUser.KeystoneUserNameSub
	}
	return ""
}

//GetPrimaryAuthUserName 获取主账号名称
func (iamToken *IamToken) GetPrimaryAuthUserName() string {
	if iamToken.PrimaryAuthUser != nil {
		return iamToken.PrimaryAuthUser.KeystoneUserName
	}
	return ""
}

//GetUserID 获取用户ID
func (iamToken *IamToken) GetUserID() string {
	if iamToken.IsPrimaryAuthUser {
		return iamToken.GetPrimaryAuthUserID()
	}
	return iamToken.GetSubAuthUserKeystoneUserID()
}

//GetPrimaryAuthUserID 获取主账号ID
func (iamToken *IamToken) GetPrimaryAuthUserID() string {
	if iamToken.PrimaryAuthUser != nil {
		return iamToken.PrimaryAuthUser.KeystoneUserID
	}
	return ""
}

//GetSubAuthUserKeystoneUserID 获取子用户ID
func (iamToken *IamToken) GetSubAuthUserKeystoneUserID() string {
	if iamToken.SubAuthUser != nil {
		return iamToken.SubAuthUser.KeystoneUserIDSub
	}
	return ""
}

//GetSubAuthUserGroupIDs 获取子用户组列表
func (iamToken *IamToken) GetSubAuthUserGroupIDs() []string {
	var groupIds []string
	if iamToken.SubUserGroups != nil {
		for _, group := range iamToken.SubUserGroups {
			groupIds = append(groupIds, group.GroupID)
		}
	}
	return groupIds
}

//Groups 用户组
type Groups struct {
	GroupID   string `json:"groupId"`
	GroupName string `json:"groupName"`
}

//SubAuthUser 子用户信息
type SubAuthUser struct {
	KeystoneUserIDSub       string          `json:"keystoneUserIdSub"`
	KeystoneUserNameSub     string          `json:"keystoneUserNameSub"`
	ShowNameSub             string          `json:"showNameSub"`
	Email                   string          `json:"email"`
	Phone                   string          `json:"phone"`
	KeystoneUserIDPrimary   string          `json:"keystoneUserIdPrimary"`
	KeystoneUserNamePrimary string          `json:"keystoneUserNamePrimary"`
	DomainID                string          `json:"domainId"`
	ProjectID               string          `json:"projectId"`
	ProjectName             string          `json:"projectName"`
	HarborEmail             string          `json:"harborEmail"`
	HarborPassword          string          `json:"harborPassword"`
	UpdateTime              int             `json:"updateTime"`
	Enabled                 float64         `json:"enabled"`
	WhetherResetPassword    float64         `json:"whetherResetPassword"`
	Role                    SubAuthUserRole `json:"role"`
}

//SubAuthUserRole 子用户角色
type SubAuthUserRole struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Policies []string `json:"policies"`
}

//PrimaryAuthUser 主用户信息
type PrimaryAuthUser struct {
	KeystoneUserID   string  `json:"keystoneUserId"`
	KeystoneUserName string  `json:"keystoneUserName"`
	DomainID         string  `json:"domainId"`
	Phone            string  `json:"phone"`
	Email            string  `json:"email"`
	HarborEmail      string  `json:"harborEmail"`
	HarborPassword   string  `json:"harborPassword"`
	EmailStatus      float64 `json:"emailStatus"`
	EmailCode        string  `json:"emailCode"`
	CreatedTime      int     `json:"createdTime"`
	UpdateTime       int     `json:"updateTime"`
}
