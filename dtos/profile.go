package dtos

type JoinProfile struct {
	FullName    string  `form:"fullName"`
	Username    string  `form:"userName"`
	Email       string  `form:"email"`
	PhoneNumber string `form:"phoneNumber"`
	Gender      int    `form:"gender"`
	Profession  string `form:"profession"`
	Nationality int    `form:"nationality"`
	BirthDate   string `form:"birthDate"`
}
