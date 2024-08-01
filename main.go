package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
) 
type Data struct {
	Id 			int `json:"id"`
	Name 		string `json:"name" form:"name"`
	Email 		string `json:"email" form:"email"`
	Password 	string `json:"-" form:"password"`
}

type Users struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Results interface{} `json:"results"`
}

func main() {
	response := []Data {}

	r := gin.Default()

	r.Use(corsMiddleware())

	r.GET("/users", func (c *gin.Context){
		c.JSON(http.StatusOK, Users{
			Success: true,
			Message: "OK",
			Results : response,
		})
	})
	r.POST("/users", func(c *gin.Context) {
		user := Data{}
		c.Bind(&user)
		user.Id = len(response) + 1
		response = append(response, user)
		c.JSON(http.StatusOK, Users{
			Success: true,
			Message: "Create user success",
			Results: user,
		})
	})
	r.POST("/auth/login", func(c *gin.Context) {
		user := Data{}
		c.Bind(&user)
		email := user.Email
		fmt.Println(email)
		searchData := true 
		if searchData {
			for searchData {
				for i:=0; i < len(response); i++ {
					dataEmail := response[i].Email
					if email == dataEmail {
						c.JSON(http.StatusOK, Users{
							Success: true,
							Message: "Login success",
						})
					}
				}
				searchData = false
			}
		} else {
			c.JSON(http.StatusUnauthorized, Users{
				Success: false,
				Message: "Wrong Email or Password",
				Results: response,
			})
		}
	})
	r.PATCH("users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		
		selected := -1
		
		for index, item := range response {
			if item.Id == id {
				selected = index
			}
		}
		
		if selected != -1 {
			form := Data{}
			c.Bind(&form)
			response[selected].Name = form.Name
			response[selected].Email = form.Email
			response[selected].Password = form.Password
			c.JSON(http.StatusOK, Users{
				Success: true,
				Message: "Deleted data success",
			})
			} else {
				c.JSON(http.StatusNotFound, Users{
					Success: false,
					Message: "Data Not Found",
				})
			}
		})
	r.DELETE("/users/:id", func(c *gin.Context){
		id, _ := strconv.Atoi(c.Param("id"))
			
		selected := -1
		
		fmt.Println(selected)
		for index, item := range response {
			if item.Id == id {
				selected = index
			}
		}

		if selected != -1 {
			responseChecked := response[selected]
			sliceLeft := response[0:selected]
			sliceRight := response[selected+1:]
			response = sliceLeft
			response = append(response, sliceRight...)
			c.JSON(http.StatusOK, Users{
				Success: true,
				Message: "Your Account has been deleted",
				Results: responseChecked,
			})
		} else {
			c.JSON(http.StatusNotFound, Users{
				Success: false,
				Message: "Data Not Found",
			})
		}
		})
		
	r.Run("localhost:8888")
}

func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
