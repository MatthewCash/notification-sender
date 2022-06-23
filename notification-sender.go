package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Notification struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Ping bool `json:"ping"`
	ImageUrl string `json:"imageUrl"`
}

func sendNotification(notification Notification, path string) error {
	jsonData, _ := json.Marshal(notification)
	
	res, error := http.Post(path, "application/json", bytes.NewBuffer(jsonData))

	if error != nil {
		return error
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	bodyText := string(body)

	if res.StatusCode != 200 {
		return errors.New(bodyText)
	}

	fmt.Println(bodyText)

	return nil
}

func getNotificationFromFlags() Notification {
	titlePtr := flag.String("title", "", "Notification Title")
	descriptionPtr := flag.String("desc", "", "Notification Description")
	pingPtr := flag.Bool("ping", false, "Should notification ping")
	imageUrlPtr := flag.String("imageUrl", "", "Image URL")
	
	flag.Parse()

	var notification Notification

	notification.Title = *titlePtr
	notification.Description = *descriptionPtr
	notification.Ping = *pingPtr
	notification.ImageUrl = *imageUrlPtr

	return notification
}

func main() {
	notification := getNotificationFromFlags()

	path := os.Getenv("HTTP_PATH")

	if (path == "") {
		path = "http://127.0.0.1:8723/notification"
	}
	
    if error := sendNotification(notification, path); error != nil {
		fmt.Println("An error occured while sending notification:")
		fmt.Println(error)
        os.Exit(1)
    }
}