package routes

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/api")
}

func Index(c *gin.Context) {
	pageContents, err := ioutil.ReadFile("pages/index.html")
	if err != nil {
		pageContents = []byte("Sorry, an error has occurred while trying to load this page!")
		fmt.Print("Error: ", err)
	}
	c.Data(200, "text/html; charset=utf-8", pageContents)
}

func Docs(c *gin.Context) {
	pageContents, err := ioutil.ReadFile("pages/docs.html")
	if err != nil {
		pageContents = []byte("Sorry, an error has occurred while trying to load this page!")
		fmt.Println("Error: ", err)
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", pageContents)
}

type request_body_template struct {
	Data      string `json:"data"`
	Algorithm string `json:"algorithm"`
}

type err_response_body_template struct {
	Error int `json:"error"`
}

type response_body_template struct {
	HashedData string `json:"hashedData"`
}

func API(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	request_body := request_body_template{}
	binding_err := c.BindJSON(&request_body)
	if binding_err != nil || request_body.Data == "" {
		response_body := err_response_body_template{1}
		c.JSON(http.StatusBadRequest, response_body)
	} else {
		if request_body.Algorithm == "MD5" {
			hashed_data := md5.Sum([]byte(request_body.Data))
			response_body := response_body_template{hex.EncodeToString(hashed_data[:])}
			c.JSON(http.StatusOK, response_body)
		} else if request_body.Algorithm == "SHA1" {
			hashed_data := sha1.Sum([]byte(request_body.Data))
			response_body := response_body_template{hex.EncodeToString(hashed_data[:])}
			c.JSON(http.StatusOK, response_body)
		} else if request_body.Algorithm == "SHA224" {
			hashed_data := sha256.Sum224([]byte(request_body.Data))
			response_body := response_body_template{hex.EncodeToString(hashed_data[:])}
			c.JSON(http.StatusOK, response_body)
		} else if request_body.Algorithm == "SHA256" {
			hashed_data := sha256.Sum256([]byte(request_body.Data))
			response_body := response_body_template{hex.EncodeToString(hashed_data[:])}
			c.JSON(http.StatusOK, response_body)
		} else if request_body.Algorithm == "SHA384" {
			hashed_data := sha512.Sum384([]byte(request_body.Data))
			response_body := response_body_template{hex.EncodeToString(hashed_data[:])}
			c.JSON(http.StatusOK, response_body)
		} else if request_body.Algorithm == "SHA512" {
			hashed_data := sha512.Sum512([]byte(request_body.Data))
			response_body := response_body_template{hex.EncodeToString(hashed_data[:])}
			c.JSON(http.StatusOK, response_body)
		} else {
			response_body := err_response_body_template{2}
			c.JSON(http.StatusBadRequest, response_body)
		}
	}
}
