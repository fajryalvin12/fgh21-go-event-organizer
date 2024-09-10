package dtos

type ChangePassword struct {
	OldPassword     string `form:"oldPassword"`
	NewPassword     string `form:"newPassword"`
	ConfirmPassword string `form:"confirmPassword" binding:"eqfield=NewPassword"`
}