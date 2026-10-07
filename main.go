package main 

import (

	"fmt"
	"os"
)

func main(){
	// fmt.Println("CLI File Organiser");

dir := os.Args[1]

// getting the directory/folder's name => 
	if(len(os.Args) == 1){
		fmt.Println("Please provide directory name")
		return
	}
	

	files, err := os.ReadDir(dir)

	if err != nil {
		fmt.Println("Error", err)
		return 
	}

	fmt.Println(files)

	// iterating/looping over 'files' (slice)
	for i, file := range files{

		// if there is a folder, skip it
		if(file.IsDir()){
			continue // skip item, keep loop running
		}

		// if there is a hidden folder/file, skip it (like .DS_Store)
		
		fmt.Println(i+1, file.Name());
	}
	
}
