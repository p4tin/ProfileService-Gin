package main

/*
 *
 * ProfileService.go
 *
 * This component will maintain a list of Profiles for a website.
 *
 * Postman collection for testing: https://www.getpostman.com/collections/bfff3fe6fa9a5846b3fb
 *
 */

import (
    "github.com/gin-gonic/gin"
	"net/http"
	"log"
)

type Profile struct {
	Name string			`json: "username"`
	Password string		`json: "password"`
	Age int				`json: "age"`
}
var Profiles []Profile

func init() {
	Profiles = make([]Profile, 3)
	Profiles[0].Name = "john@aol.com"
	Profiles[0].Password = "test1234"

	Profiles[1].Name = "harry@comcast.com"
	Profiles[1].Password = "testing1"

	Profiles[2].Name = "sally@microsoft.com"
	Profiles[2].Password = "pass4321"
}

func main() {
	//gin.SetMode(gin.ReleaseMode)  // NOTE: or "export GIN_MODE=release" and run
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/profile", profile_get)									//Get All profiles (version 1)
		v1.GET("/profile/:name", profile_get_one)						//Get a specific profile based in Name (version 1)
		v1.POST("/profile", profile_create)                				//Create a new profile (version 1)
		v1.PUT("/profile", profile_update)          					//Update some profile fields (version 1)
		v1.DELETE("/profile/:name", profile_delete)						//Delete a specific profile (version 1)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/profile", profile_get_v2)									//Get All profiles (version 2)
	}

	// Listen and server on 0.0.0.0:8080
	//router.Run(":8080")
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", router)
	if err != nil {
		log.Printf("Web Server cannot start (https): %v\n", err)
	}
}

/* ********************** v1 functions ********************** */

func profile_get(c *gin.Context) {
	c.JSON(http.StatusOK, Profiles)
}

func profile_get_one(c *gin.Context) {
	name := c.Param("name")
	for _, profile := range Profiles {
		if(name == profile.Name) {
			c.JSON(http.StatusOK, profile)
			return
		}
	}
	var json struct{}
	c.JSON(http.StatusNotFound, json)
}

func profile_create(c *gin.Context) {
	var json Profile
	if c.BindJSON(&json) == nil {
		for _, profile := range Profiles {
			if(json.Name == profile.Name) {
				c.JSON(http.StatusConflict, json)
				return;
			}
		}
		Profiles = append(Profiles, json)
		c.JSON(http.StatusCreated, json)
		return
	}
	c.JSON(http.StatusBadRequest, json)
}

func profile_update(c *gin.Context) {
	var json Profile
	if c.BindJSON(&json) == nil {
		for i, profile := range Profiles {
			if(json.Name == profile.Name) {
				Profiles[i].Password = json.Password
				Profiles[i].Age = json.Age
				c.JSON(http.StatusOK, json)
				return;
			}
		}
	}
	c.JSON(http.StatusBadRequest, json)
}


func profile_delete(c *gin.Context) {
	var json struct{}
	name := c.Param("name")
	for i, profile := range Profiles {
		if(name == profile.Name) {
			Profiles = append(Profiles[:i], Profiles[i+1:]...)
			c.JSON(http.StatusOK, json)
			return;
		}
	}
	c.JSON(http.StatusNotFound, json)
}

/* ********************** v2 functions ********************** */
func profile_get_v2(c *gin.Context) {
	c.JSON(http.StatusOK, Profiles)
}
