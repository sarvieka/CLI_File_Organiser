package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// fmt.Println("CLI File Organiser")

	// Check if the user provided a directory name
	if len(os.Args) == 1 {
		fmt.Println("Please provide directory name")
		return
	}

	// Getting the directory/folder's name
	dir := os.Args[1]

	files, err := os.ReadDir(dir)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(files)

	categories := map[string]string{
		".jpg": "Images",
		".png": "Images",
		".pdf": "Documents",
		".mp3": "Audio",
		".mp4": "Videos",
	}

	// Iterating/looping over 'files' (slice)
		for i, file := range files {

			// If there is a folder, skip it
			if file.IsDir() {
				continue
			}

			// Get the extension of the current file
			extension := filepath.Ext(file.Name())

			// Look up the extension in our categories map
			// category -> the category name
			// exists   -> whether the extension exists in the map
			category, exists := categories[extension]

			// If the extension is not supported, skip this file
			if !exists {
				fmt.Println("Unsupported file:", file.Name())
				continue
			}

			fmt.Println(
				i+1,
				file.Name(),
				"extension:", extension,
				"category:", category,
			)
		}
}