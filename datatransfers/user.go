package datatransfers

type UserLogin struct {
	Username string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSignup struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Bio      string `json:"bio" binding:"-"`
}
