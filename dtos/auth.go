package dtos

type FormRegister struct {
	FullName		string `form:"fullName"`
	Email			string `form:"email"`
	Password		string `form:"password"`
	ConfirmPassword	string `form:"confirmPassword" binding:"eqfield=Password"`
	Username 		string `form:"username"`
}

type FormLogin struct {
	Email			string `form:"email"`
	Password		string `form:"password"`
}