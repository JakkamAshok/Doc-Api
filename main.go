package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// _ "github.com/JakkamAshok/Doc-Api"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

type message struct {
	ID       string `json:"id"`
	Header   string `json: "header"`
	Provider string `json: "provider"`
	Quantity string `json: "quantity"`
}

var messages = []message{
	{ID: "1", Header: "Hello", Provider: "Airtel1", Quantity: "2"},
	{ID: "2", Header: "Hello1", Provider: "Airtel2", Quantity: "6"},
	{ID: "3", Header: "Hello2", Provider: "Airtel3", Quantity: "8"},
}

func getMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, messages)
}

func addMessages(c *gin.Context) {
	var newMessage message
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}

	messages = append(messages, newMessage)
	c.IndentedJSON(http.StatusCreated, newMessage)
}

func main() {
	router := gin.Default()
	router.GET("/getMessages", getMessages)
	router.POST("/addMessage", addMessages)
	// user := router.Group("/api/v1/users")
	// router.GET(relativePath:"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("localhost:8080")
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

// --------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	_ "your-module/docs"

// 	"github.com/gorilla/mux"
// )

// // @title Your API Title
// // @version 1.0
// // @description This is a sample API documentation
// // @termsOfService https://example.com/terms/
// // @contact.name API Support
// // @contact.url https://www.example.com/support
// // @license.name MIT
// // @license.url https://opensource.org/licenses/MIT
// func main() {
// 	router := mux.NewRouter()

// 	// @Summary Say hello
// 	// @Description Get a friendly greeting
// 	// @ID say-hello
// 	// @Produce plain
// 	// @Success 200 {string} string "Hello, World!"
// 	// @Router /hello [get]
// 	router.HandleFunc("/hello", helloHandler).Methods("GET")

// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!")
// }
