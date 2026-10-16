package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// next step : create a category folder, 

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

	// create cateories map:
	categories := map[string]string{
		".jpg": "Images",
		".png": "Images",
		".pdf": "Documents",
		".mp3": "Audio",
		".mp4": "Videos",
	}

	// Iterating/looping over 'files' (slice)
		for _, file := range files {

			// If there is a folder, skip it
			if file.IsDir() {
				continue
			}

			// Get the extension of the current file
			extension := filepath.Ext(file.Name())

			// Look up the extension in our categories map
			// category -> the category name
			// exists   -> whether the extension exists in the map
			category, exists := categories[extension] // comma-ok-idiom!

			// If the extension is not supported, skip this file
			if !exists {
				fmt.Println(" => Unsupported file:", file.Name())
				continue
			}

			categoryPath := filepath.Join(dir, category)

			info, err := os.Stat(categoryPath)

			if err == nil {
				fmt.Println("Path exists. Is it a directory?", info.IsDir())
			} else {
				fmt.Println("Path does not exist or cannot be accessed:", err)
				err = os.Mkdir(categoryPath, 0755);
					if err != nil {
        				fmt.Println("Error creating folder:", err)
    				} else {
        				fmt.Println("Category folder created:", categoryPath)
    				}
			}

			

			fmt.Println(
				" =>",
				"File Name:", file.Name(),
				" | Extension:", extension,
				" | Category:", category,
				" | Category folder path:", categoryPath,
			)
		}
}