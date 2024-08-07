package models

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password" binding:"required,min=8"`
}

var Response = []Users{
	{
		Id:       1,
		Name:     "Admin",
		Email:    "admin@mail.com",
		Password: "1234",
	},
}

func FindAllUsers() []Users {
	return Response
}
func FindUserId(id int) Users {

	users := Response
	selectedUser := Users{}

	for _, item := range users {
		if id == item.Id {
			selectedUser = item
		}
	}

	return selectedUser
}
func CreateNewUser(data Users) Users {
	id := 0

	for _, v := range Response {
		id = v.Id
	}
	data.Id = id + 1
	Response = append(Response, data)
	return data
}
func EditTheUser(data Users, id int) Users {
	num := -1

	for index, v := range Response {
		if id == v.Id {
			num = index
		}
	}

	if num != 0 {
		Response[num].Name = data.Name
		Response[num].Email = data.Email
		Response[num].Password = data.Password
		Response[num].Id = data.Id
	}

	return data
}
func RemoveUser(id int) Users {
	index := -1
	userRemoved := Users{}

	for idx, v := range Response {
		if v.Id == id {
			index = idx
			userRemoved = v
		}
	}
	sliceLeft := Response[:index]
	sliceRight := Response[index+1:]
	if userRemoved.Id != 0 {
		Response = append(sliceLeft, sliceRight...)
	}

	return userRemoved
}