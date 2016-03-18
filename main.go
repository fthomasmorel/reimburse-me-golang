package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	log.Fatal(http.ListenAndServe(":9000", NewRouter()))
}

// UploadImage will manage the upload image from a POST request
func UploadImage(r *http.Request) string {

	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("image")
	if err != nil {
		return "error"
	}
	defer file.Close()

	fileName := RandomString(40)
	f, err := os.OpenFile("./img/"+fileName+".png", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "error"
	}

	defer f.Close()
	io.Copy(f, file)

	return fileName
}

// RandomString generates a randomString (y)
func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
