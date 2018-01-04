package services

import (
	"golanggeeks/common"
	"golanggeeks/model"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

// AddMenu is use to add the menu
func AddMenu(c *gin.Context) {

	var menu model.Menu
	err := c.BindJSON(&menu)
	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Request is invalid..",
		})
		log.Println(err)
		return
	}
	session, err := mgo.Dial("localhost")
	defer session.Close()

	db := session.DB(common.DB).C("menus")
	if db.Insert(menu) == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Issue occured..",
		})
		log.Println(err)
	}
}

// EditMenu is use to update the menu
func EditMenu(c *gin.Context) {
	session, err := mgo.Dial("localhost")
	var menu model.Menu
	err = c.BindJSON(&menu)
	if err != nil {
		log.Println("error")
	}
	menuCollection := session.DB(common.DB).C("menus")
	if menuCollection.Update(bson.M{"name": menu.Name}, menu) == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Issue occured..",
		})
		log.Println(err)
	}
}

// DeleteMenu is use to delete the menu
func DeleteMenu(c *gin.Context) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Println(err)
	}

	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "Identifier can not be blank...",
		})
		return
	}
	menuCollection := session.DB(common.DB).C("menus")
	if menuCollection.Remove(bson.M{"name": name}) == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Issue occured..",
		})
		log.Println(err)
	}
}

// GetMenu is use to get specific menu
func GetMenu(c *gin.Context) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Println(err)
	}

	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "Identifier can not be blank...",
		})
		return
	}
	var menu model.Menu
	err = session.DB(common.DB).C("menus").Find(bson.M{"name": name}).One(&menu)
	if err == nil {
		log.Println(menu)
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "Success",
			"result":  menu,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Issue occurred..",
		})
	}
}

// GetMenus is use to get all menus
func GetMenus(c *gin.Context) {
	session, _ := mgo.Dial("localhost")
	menusCollection := session.DB(common.DB).C("menus")
	var menus []model.Menu
	err := menusCollection.Find(bson.M{}).All(&menus)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"result":  menus,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "error occurred..",
		})
	}
}
