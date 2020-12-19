package datatransfers

import (
	"time"
)

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSignup struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Bio      string `json:"bio" binding:"-"`
}

type UserUpdate struct {
	Email string `json:"email" binding:"-"`
	Bio   string `json:"bio" binding:"-"`
}

type UserInfo struct {
	Username  string    `uri:"username" json:"username"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
}
