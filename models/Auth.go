package models

type Role int

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // here is just for example
	Role     Role   `json:"-"`
}

const (
	ADMIN        Role = 0x1
	FINANCE      Role = 0x1 << 1
	MOBILE       Role = 0x1 << 2
	SYSTEM_ADMIN Role = 0x1 << 2
)

func (r Role) IsAdmin() bool {
	return r&ADMIN != 0
}

func (r Role) IsFinance() bool {
	return r&FINANCE != 0
}

func (r Role) IsMobile() bool {
	return r&MOBILE != 0
}

func (r Role) IsSystemAdmin() bool {
	return r&SYSTEM_ADMIN != 0
}

var NaiveDatastore = map[string]User{
	"admin":        {"admin", "admin@admin.com", "admin", ADMIN},
	"finance":      {"finance", "f@f.com", "finance", FINANCE},
	"mobile":       {"mobile", "m@m.com", "mobile", MOBILE},
	"system_admin": {"system_admin", "sa@sa.com", "system_admin", SYSTEM_ADMIN},
}
