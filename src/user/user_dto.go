package user

type UserDto struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Token    string `json:"token"`
}

type LoginDto struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
