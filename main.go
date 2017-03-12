package main

import (
  "encoding/json"
  "net/http"
  "os"
  "log"

  "github.com/gin-gonic/gin"
)

type SpaceList []struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	TotalCheckins int `json:"total_checkins"`
	ImageLink string `json:"image_link"`
	PictureLink string `json:"picture_link"`
}

func LoadSpaceList(url string) {
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal("NewRequest: ", err)
    return
  }

  req.Header.Add("Authorization", "Token " + os.Getenv("STUDYSPACE_KEY"))

  client := &http.Client{}

  resp, err := client.Do(req)
  if err != nil {
    log.Println("Do: ", err)
    return
  }

  defer resp.Body.Close()

  var records SpaceList

  if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
    log.Println(err)
  }

  log.Println("Look: " + records[0].Name + "!")
}

func main() {
  // Creates a gin router with default middleware:
  // logger and recovery (crash-free) middleware
  router := gin.Default()
  router.Use(gin.Logger())
  router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static") // For static assets

  router.GET("/", func(c *gin.Context) {
    LoadSpaceList("https://study.space/api/v1/spaces.json")
		c.HTML(http.StatusOK, "index.tmpl.html", ", world")
	})
  // By default it serves on :8080 unless a
  // PORT environment variable was defined.
  router.Run()
}
