package models

type Profile struct {
	Id            int     `json:"id" db:"id"`
	Picture       *string `json:"picture" db:"picture"`
	FullName      string  `json:"fullName" db:"full_name"`
	BirthDate     *string `json:"birthDate" db:"birth_date"`
	Gender        *int    `json:"gender" db:"gender"`
	PhoneNumber   *string `json:"phoneNumber" db:"phone_number"`
	Profession    *string `json:"profession"`
	NationalityId *int    `json:"nationalityId" db:"nationality_id"`
	UserId        int     `json:"userId" db:"user_id"`
}
type JoinProfile struct {
	Id          int     `json:"id"`
	Picture     *string `json:"picture"`
	FullName    string  `json:"fullName" db:"full_name" form:"fullName"`
	Username    string  `json:"username,omitempty" db:"username" form:"userName"`
	Email       string  `json:"email,omitempty" form:"email" db:"email"`
	Gender      *int    `json:"gender,omitempty" form:"gender" db:"gender"`
	PhoneNumber *string `json:"phoneNumber,omitempty" db:"phone_number" form:"phoneNumber"`
	Profession  *string `json:"profession,omitempty" form:"profession" db:"profession"`
	Nationality *int    `json:"nationality,omitempty" form:"nationality" db:"nationality_id"`
	BirthDate   *string `json:"birthDate,omitempty" db:"birth_date" form:"birthDate"`
}