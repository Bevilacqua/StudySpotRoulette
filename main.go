package main

import (
  "encoding/json"
  "net/http"
  "os"
  "log"

  "github.com/gin-gonic/gin"
)

type Space struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	ShortDescription string `json:"short_description"`
	CurrentCheckins string `json:"current_checkins"`
	ImageLink string `json:"image_link"`
	PictureLink string `json:"picture_link"`
}

func LoadSpace(url string) Space {
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal("NewRequest: ", err)
  }

  req.Header.Add("Authorization", "Token " + os.Getenv("STUDYSPACE_KEY"))

  client := &http.Client{}

  resp, err := client.Do(req)
  if err != nil {
    log.Println("Do: ", err)
  }

  defer resp.Body.Close()

  var space Space

  if err := json.NewDecoder(resp.Body).Decode(&space); err != nil {
    log.Println(err)
  }
  return space
}

func main() {
  // Creates a gin router with default middleware:
  // logger and recovery (crash-free) middleware
  router := gin.Default()
  router.Use(gin.Logger())
  router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static") // For static assets

  router.GET("/", func(c *gin.Context) {
    space := LoadSpace("https://study.space/api/v1/spaces.json")
    log.Println(space.Name)
		c.HTML(http.StatusOK, "index.tmpl.html",gin.H{
            "name": space.Name,
            "description": space.ShortDescription,
            "current_checkins": space.CurrentCheckins,
            "picture_link": space.PictureLink,
            "id": space.ID,
    })
	})
  // By default it serves on :8080 unless a
  // PORT environment variable was defined.
  router.Run()
}
