package dtos

type JoinProfile struct {
	FullName    string  `form:"fullName"`
	Username    string  `form:"userName"`
	Email       string  `form:"email"`
	Gender      *int    `form:"gender"`
	PhoneNumber *string `form:"phoneNumber"`
	Profession  *string `form:"profession"`
	Nationality *int    `form:"nationality"`
	BirthDate   *string `form:"birthDate"`
}
