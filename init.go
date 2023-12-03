package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a day number")
		return
	}

	// Get directory number from args
	arg := os.Args[1]
	num, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Invalid number:", err)
		return
	}
	formattedNum := fmt.Sprintf("%02d", num)

	// Create day directory
	dayDirName := fmt.Sprintf("days/%s", formattedNum)
	createDirectory(dayDirName, 0755)

	// Create input directory
	inputDirName := fmt.Sprintf("%s/input", dayDirName)
	createDirectory(inputDirName, 0755)

	// Create input.txt file
	filePath := fmt.Sprintf("%s/input.txt", inputDirName)
	createFile(filePath, nil)

	// Template replacements
	replacements := map[string]string{
		"Template": "main",
		"01":       formattedNum,
	}

	// Create day{arg}.go file
	fileName := fmt.Sprintf("%s/day%s.go", dayDirName, formattedNum)
	copyTemplateFile("template.go", fileName, replacements)
	// Create day{arg}_test.go file
	testFileName := fmt.Sprintf("%s/day%s_test.go", dayDirName, formattedNum)
	copyTemplateFile("template_test.go", testFileName, replacements)

	// Success
	fmt.Printf("Files and directories created for Day %s. Good luck!", formattedNum)
}

func createDirectory(dirName string, perm os.FileMode) {
	err := os.MkdirAll(dirName, perm)

	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
}

func createFile(path string, data []byte) {
	file, err := os.Create(path)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// Close file after creating
	defer file.Close()

	// If data is passed in, write it to file
	if len(data) > 0 {
		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

}

func copyTemplateFile(templateFileName, newFileName string, replacements map[string]string) {
	templateContent, err := os.ReadFile(templateFileName)

	if err != nil {
		fmt.Println("Error reading template file:", err)
		return
	}

	modifiedContent := string(templateContent)
	for placeholder, replacement := range replacements {
		modifiedContent = strings.ReplaceAll(modifiedContent, placeholder, replacement)
	}

	createFile(newFileName, []byte(modifiedContent))
}
