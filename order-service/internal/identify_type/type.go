package identify_type

import "time"

type UserKeyToken struct {
	ID           string    `json:"id"`
	UserId       string    `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	PrivateKey   string    `json:"private_key"`
	PublicKey    string    `json:"public_key"`
	ExpiredAt    time.Time `json:"expireAt"`
}
type UserBase struct {
	UserID      string         `json:"user_id"`
	UserAccount string         `json:"user_account"`
	UserSalt    string         `json:"user_salt"`
	UserRoles   []UserRoleType `json:"user_roles"`
}
type UserRoleType struct {
	RoleID          string         `json:"role_id"`
	RoleName        string         `json:"role_name"`
	RoleDescription string         `json:"role_description"`
	RoleMenus       []UserMenuType `json:"role_menus"`
}
type UserMenuType struct {
	MenuID     string `json:"menu_id"`
	MenuName   string `json:"menu_name"`
	MenuUrl    string `json:"menu_url"`
	MenuPrefix string `json:"menu_prefix"`
	MenuPid    string `json:"menu_pid"`
}
