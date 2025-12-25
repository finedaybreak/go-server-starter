package dto

// 更新用户信息请求DTO
type UserUpdateInfoReqDto struct {
	Nickname  *string `json:"nickname" form:"nickname" binding:"required,min=2,max=20"`
	AvatarURL *string `json:"avatarURL" form:"avatarURL" binding:"omitempty,url"`
	Desc      *string `json:"desc" form:"desc" binding:"omitempty,max=200"`
}

type UserInfoResDto struct {
	UniCode     string   `json:"uniCode"`
	Email       string   `json:"email"`
	Mobile      string   `json:"mobile"`
	CountryCode string   `json:"countryCode"`
	Desc        string   `json:"desc"`
	Nickname    string   `json:"nickname"`
	AvatarURL   string   `json:"avatarURL"`
	Roles       []string `json:"roles"`
}

type UserTableQueryReqDto struct {
	PaginationReqDto
	Nickname    *string `json:"nickname" form:"nickname"`
	Email       *string `json:"email" form:"email"`
	Mobile      *string `json:"mobile" form:"mobile"`
	CountryCode *string `json:"countryCode" form:"countryCode"`
}

type UserListItemResDto struct {
	ID          uint64   `json:"id"`
	CreatedAt   string   `json:"createdAt"`
	UniCode     string   `json:"uniCode"`
	Email       string   `json:"email"`
	Mobile      string   `json:"mobile"`
	CountryCode string   `json:"countryCode"`
	Desc        string   `json:"desc"`
	Nickname    string   `json:"nickname"`
	AvatarURL   string   `json:"avatarURL"`
	Roles       []string `json:"roles"`
}
