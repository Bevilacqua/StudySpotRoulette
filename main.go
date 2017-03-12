package main

import (
  "encoding/json"
  "net/http"
  "os"
  "log"
  "math/rand"
  "time"

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

type Space struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	TotalCheckins int `json:"total_checkins"`
	ImageLink string `json:"image_link"`
	PictureLink string `json:"picture_link"`
}

func LoadSpaceList(url string) SpaceList {
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal("NewRequest: ", err)
    return nil
  }

  req.Header.Add("Authorization", "Token " + os.Getenv("STUDYSPACE_KEY"))

  client := &http.Client{}

  resp, err := client.Do(req)
  if err != nil {
    log.Println("Do: ", err)
    return nil
  }

  defer resp.Body.Close()

  var records SpaceList

  if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
    log.Println(err)
  }
  return records
}

func ChooseRandom(records SpaceList) Space {
  s := rand.NewSource(time.Now().Unix())
  r := rand.New(s) // initialize local pseudorandom generator
  return records[r.Intn(len(records))]
}

func main() {
  // Creates a gin router with default middleware:
  // logger and recovery (crash-free) middleware
  router := gin.Default()
  router.Use(gin.Logger())
  router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static") // For static assets

  router.GET("/", func(c *gin.Context) {
    space := ChooseRandom(LoadSpaceList("https://study.space/api/v1/spaces.json"))
    log.Println(space.Name)
		c.HTML(http.StatusOK, "index.tmpl.html", space.Name)
	})
  // By default it serves on :8080 unless a
  // PORT environment variable was defined.
  router.Run()
}
