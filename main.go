package main

import (
	// "encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func IndexHandler(c *gin.Context) {
	// name := c.Params.ByName("name")
	// c.JSON(200, gin.H{
	// 	"message":"hello "+ name,
	// })
	// c.XML(200, Person{FirstName: "Obaid Ullah",
	// 	LastName: "Khan"})
}

// type Person struct {
// 	XMLName   xml.Name `xml:"person"`
// 	FirstName string   `xml:"firstName,attr"`
// 	LastName  string   `xml:"lastName,attr"`
// }

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

var recipes []Recipe

func init() {
	fmt.Println("hi im printing init 1")
	recipes = make([]Recipe, 0)
	fmt.Println("hi im printing init 2")

}

func NewRecipeHandler(c *gin.Context) {
	fmt.Println("hi im printing")
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)

}

// main function
func main() {
	fmt.Println("before gin fun calll")
	router := gin.Default()
	fmt.Println("before fun calll")
	router.POST("/recipes", NewRecipeHandler)
	fmt.Println("after fun calll")

	router.Run()
}
