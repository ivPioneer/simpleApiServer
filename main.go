package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

var jsonUsers Users

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUsers(context *gin.Context) {

	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &jsonUsers)

	context.IndentedJSON(http.StatusOK, jsonUsers)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("127.0.0.1:3000")

}
