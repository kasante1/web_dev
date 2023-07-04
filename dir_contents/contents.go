package dir_contents

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"regexp"
)


//html file
// file contents

func HtmlFileContents() string {
	var mainFileContents string = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<link rel="stylesheet" href="loopstudio.css" />
		<title>Document</title>
	</head>
	<body>
		<nav>
		</nav>

		<footer>
	
		</footer>
	</body>
	</html> 
	
`
	return mainFileContents
}


func CssFileContents() string{
	cssFileContents := ` 
	*{
		margin: 0px;
		padding: 0px;
		box-sizing: border-box;
	}
	
	`
	return cssFileContents

}

func JsFileContents() string{
	jsFileContents := ``
	return jsFileContents

}


// create project files
func CreateProjectFiles(SubDirectories, fileName, file_contents string) {
	// file directory
	file_path := filepath.Join(SubDirectories, fileName)

	// check if file exist or otherwise
	_, err := os.Stat(file_path)

	// if file does not exits, create file.
	if errors.Is(err, os.ErrNotExist) {
		// write content to files
		WriteProjectFiles(file_path, file_contents)
		fmt.Println("[OK]", fileName, "created succesfully")
	} else {
		fmt.Println("[X]", fileName, "failed. already exits!")
	}
}

func WriteProjectFiles(fileName, file_contents string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	if _, err := file.WriteString(file_contents); err != nil {
		fmt.Println(err)
	}
}

// create file, folder, tests and test directories