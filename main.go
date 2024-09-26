package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, world!")
	rooms := make(map[int]*Hub)
	r := gin.Default()
	r.LoadHTMLFiles("home.html")

	r.GET("/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.GET("/ws/:id", func(c *gin.Context) {
		var hub *Hub
		log.Println(c.Param("id"))
		f := c.Param("id")
		id, err := strconv.Atoi(f)
		if err != nil {
			log.Println(err)
			return
		}

		if _, ok := rooms[id]; !ok {
			hub = NewHub(id)
			rooms[hub.id] = hub
			go hub.Run()
		} else {
			hub = rooms[id]
		}
		serveWs(hub, c.Writer, c.Request)
	})

	r.Run(":8000")
}
