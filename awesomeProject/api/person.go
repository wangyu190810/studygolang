package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"fmt"
	. "awesomeProject/models"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetPersonsApi(c *gin.Context) {
	var p Person
	//var id = c.Param("id")     // Notice: string
	var id = c.Query("id")     // Notice: string
	p.Id, _ = strconv.Atoi(id) // Notice: int(id) not work!
	person, err := p.GetPersons()
	if err != nil {
		//log.Println("Error:", res.Desc)
		c.JSON(http.StatusOK, gin.H{"data":"msg",})
	} else {
		c.JSON(http.StatusOK, gin.H{"data":person,})
	}
}
