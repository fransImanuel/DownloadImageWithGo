package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "https://apigw.vnetcloud.com/minio_dev/sdh-order/5UE1hiJHCq.png"
	filepath := "image.jpg"

	err := downloadImage(url, filepath)
	if err != nil {
		log.Fatalf("Error downloading image: %v", err)
	}

	log.Println("Image downloaded successfully!")
}

func downloadImage(url, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
