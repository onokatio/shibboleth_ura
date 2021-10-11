package main

import (
	"encoding/json"
	"http"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		challenge := c.Query("login_challenge")
		c.JSON(200, gin.H{
			"status": "please login",
			"challenge": challenge,
		})
	})
	r.POST("/login", func(c *gin.Context) {
		email := c.PostForm("email")
		pass := c.PostForm("pass")
		challenge := c.Query("login_challenge")

		if (email == "foo" && pass == "bar"){
			resp, err := http.Get("http://127.0.0.1:4445/oauth2/auth/requests/login/accept?login_challenge=" + challenge)
			if err != nil {
				c.JSON(200, gin.H{
					"accept resuest": err,
				})
			}
			str, err := io.ReadAll(resp.Body)
			redirect_resp := map[string]string{}
			json.Unmarshal(str, &redirect_resp)
				c.JSON(200, gin.H{
					"accept request": redirect_resp,
				})
		} else {
			c.JSON(200, gin.H{
				"status": "error",
			})
		}

	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
