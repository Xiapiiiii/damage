package types

type UserLoginReq struct {
	UserId   string `json:"user_id" binding:"required"`
	PassWord string `json:"pass_word" binding:"required"`
}

type UserRegisterInfoReq struct {
	UserName string `json:"user_name" binding:"required"`
	UserId   string `json:"user_id" binding:"required"`
	PassWord string `json:"pass_word" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
