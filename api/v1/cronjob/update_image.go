package cronjob

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"time"
	"crypto/md5"
	"io"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"regexp"
	"log"
	"encoding/hex"
)


func Update_S3_image_url() {
  db, err := sql.Open("postgres", "password=password host=localhost dbname=online_test_dev sslmode=disable")
  if err != nil {
    panic(err)
  }
  get_images, err := db.Query("SELECT id,image from questions where image != 'nil'")
  if err != nil || get_images == nil {
    panic(err)
  }

  for get_images.Next() {
    var id int
    var image string

    err = get_images.Scan(&id, &image)

    if image != "" {
      auth := aws.Auth{
        AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
        SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
      }
      euwest := aws.USWest2
      connection := s3.New(auth, euwest)
      thumbnails := connection.Bucket("q-auth")
      public_thumbnails, err := thumbnails.List(image, "", "", 1000)
      if err != nil {
        panic(err)
      }
      for _, v := range public_thumbnails.Contents {
      // Creates a URL to access thumbnail
        image_url := connection.Bucket("q-auth").SignedURL(v.Key, time.Now().Add(5*time.Hour))
        update_image_url, err := db.Query("UPDATE questions SET image_url = $1 where id = $2", image_url, id)
        if err != nil || update_image_url == nil {
          panic(err)
        }
        defer update_image_url.Close()
      }
    }
  }
  fmt.Println("All task removed")
  db.Close()
}