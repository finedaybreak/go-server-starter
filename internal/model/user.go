package model

type User struct {
	Model
	UniCode     string     `gorm:"uniqueIndex;not null" json:"uniCode"`
	Email       string     `gorm:"index" json:"email"`
	Mobile      string     `gorm:"index:idx_mobile_country" json:"mobile"`
	CountryCode string     `gorm:"index:idx_mobile_country" json:"countryCode"`
	Desc        string     `json:"desc"`
	Password    string     `json:"-"`
	Salt        string     `json:"-"`
	Nickname    string     `gorm:"index" json:"nickname"`
	AvatarURL   string     `json:"avatarURL"`
	Roles       []UserRole `gorm:"many2many:user_role_refs;constraint:OnDelete:CASCADE" json:"roles"`
}

func (User) TableName() string {
	return "users"
}
