package models

import "github.com/actiontech/dtle/driver/common"

type UserListReq struct {
	FilterUsername string `query:"filter_username"`
	FilterTenant   string `query:"filter_tenant"`
}

type UserListResp struct {
	UserList []*common.User `json:"user_list"`
	BaseResp
}

type TenantListResp struct {
	TenantList []string `json:"tenant_list"`
	BaseResp
}

type CreateUserReqV2 struct {
	Username string `json:"username" validate:"required"`
	Tenant   string `json:"tenant" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Remark   string `json:"remark"`
	PassWord string `json:"pass_word" validate:"required"`
}

type CreateUserRespV2 struct {
	BaseResp
}

type UpdateUserReqV2 struct {
	Username string `json:"username" validate:"required"`
	Tenant   string `json:"tenant" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Remark   string `json:"remark"`
}

type UpdateUserRespV2 struct {
	BaseResp
}

type ResetPasswordReqV2 struct {
	Username            string `json:"username" validate:"required"`
	Tenant              string `json:"tenant" validate:"required"`
	Password            string `json:"password" validate:"required"`
	CurrentUserPassword string `json:"current_user_password" validate:"required"`
}

type ResetPasswordRespV2 struct {
	BaseResp
}

type DeleteUserReqV2 struct {
	Username string `form:"username" validate:"required"`
	Tenant   string `form:"tenant" validate:"required"`
}

type DeleteUserRespV2 struct {
	BaseResp
}

type CurrentUserResp struct {
	CurrentUser *common.User `json:"current_user"`
	BaseResp
}

type ListActionRespV2 struct {
	Authority []MenuItem `json:"authority"`
	BaseResp
}
