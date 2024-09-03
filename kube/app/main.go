package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Limit the size of the uploaded file (e.g., 10MB)
	r.ParseMultipartForm(10 << 20) // 10MB

	// Get the file from the form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "Error Retrieving the File: %v", err)
		return
	}
	defer file.Close()

	// Create a file on the server where the uploaded file will be saved
	dst, err := os.Create(filepath.Join("uploads", handler.Filename))
	if err != nil {
		fmt.Fprintf(w, "Error Creating the File: %v", err)
		return
	}
	defer dst.Close()

	// Copy the uploaded file data to the created file
	if _, err := io.Copy(dst, file); err != nil {
		fmt.Fprintf(w, "Error Saving the File: %v", err)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File: %s", handler.Filename)
}

func main() {
	// Ensure the uploads directory exists
	os.MkdirAll("uploads", os.ModePerm)

	// Read port from the environment variable, default to 8080 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/upload", uploadFileHandler)

	fmt.Printf("Starting server on :%s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
