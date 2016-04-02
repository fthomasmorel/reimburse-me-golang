package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		log.Fatal(http.ListenAndServe(":9000", NewRouter()))
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		log.Fatal(http.ListenAndServe(":9090", http.FileServer(http.Dir("./img/"))))
		wg.Done()
	}()
	wg.Wait()

}

// UploadImage will manage the upload image from a POST request
func UploadImage(r *http.Request) string {

	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println("error 1")
		fmt.Println(err)
		return "error"
	}
	defer file.Close()

	fileName := RandomString(40)
	f, err := os.OpenFile("./img/"+fileName+".png", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error 2")
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
