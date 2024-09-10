package models

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password" binding:"required,min=8"`
	Username string `json:"username,omitempty" form:"username" binding:"required"`
}

type ChangePassword struct {
	OldPassword     string `json:"oldPassword" form:"oldPassword"`
	NewPassword     string `json:"newPassword" form:"newPassword"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" binding:"eqfield=NewPassword"`
}

