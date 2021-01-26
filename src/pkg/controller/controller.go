package controller

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"

)

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

var userSlice = []User{
	User{
		Username:  "Welcome",
		FirstName: "Web-Clan",
		LastName:  "Okay",
		Email:     "adefemi171@gmail.com",
	},
	User{
		Username:  "Welcome",
		FirstName: "Mugun-Clan",
		LastName:  "Okay",
		Email:     "adefemi171@gmail.com",
	},
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}

func UserHandler(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	enc := json.NewEncoder(w)
	enc.SetIndent(" ", "\t")
	enc.Encode(userSlice)
}


func GetHeader(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
    //Iterate over all header fields
    for k, v := range r.Header {
        fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
    }

    fmt.Fprintf(w, "Host = %q\n", r.Host)
    fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
    //Get value for a specified token
    fmt.Fprintf(w, "\n\nFinding value of \"Accept\" %q", r.Header["Accept"])
}