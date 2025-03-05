package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CreateLogFile(file_name, initialContent string) error {
	createFile, err := os.Create(file_name)
	if err != nil {
		LogError(fmt.Sprintf("Error in creating file %v", err))
		return err
	}
	defer createFile.Close()
	content := "Service started successfully."
	_, err = createFile.WriteString(content)
	if err != nil {
		LogError(fmt.Sprintf("Error in Writing to file %v", err))

		return err
	}
	return nil
}

func ReadLogFile(filename string) (string, error) {

	content, err := os.ReadFile(filename)
	if err != nil {
		LogError(fmt.Sprintf("Error in Reading file %v", err))
		return "", err
	}
	return string(content), nil
}
func AppendToLogFile(fileName string, contentToAdd string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		LogError(fmt.Sprintf("Error in open file %v", err))
		return err
	}

	_, err = file.WriteString(contentToAdd)
	if err != nil {
		LogError(fmt.Sprintf("Error in Writing file %v", err))
		return err
	}
	defer file.Close()

	log.Println("New content appended to ", fileName)
	content, err := os.ReadFile(file.Name())
	if err != nil {
		LogError(fmt.Sprintf("Error in Reading file %v", err))

		return err
	}
	fmt.Println(string(content))
	return nil
}
func ReportFileStats(fileName string) error {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		LogError(fmt.Sprintf("Error in Stat file %v", err))

		return err
	}
	file, err := os.Open(fileName)
	if err != nil {
		LogError(fmt.Sprintf("Error in Opening file %v", err))

		return err
	}
	scaneer := bufio.NewScanner(file)
	no := 0
	for scaneer.Scan() {
		no++
	}
	log.Printf("File Name : %s\n", fileInfo.Name())
	log.Printf("Number of lines : %v", no)
	log.Printf("Size in bytes : %d", fileInfo.Size())
	log.Printf("Modified time of file : %s", fileInfo.ModTime())
	return nil
}
func LogError(errorMessage string) {
	errorFile := "error_log.txt"
	file, err := os.OpenFile(errorFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to log error ", err)
		return
	}
	defer file.Close()
	logmessage := fmt.Sprintf("Error %s", errorMessage)
	file.WriteString(logmessage)
}
func ReadFromOffset(fileName string, offset int64, length int) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		LogError(fmt.Sprintf("Failed to open file for reading at offset: %v", err))
		return "", err
	}
	defer file.Close()

	_, err = file.Seek(offset, 0)

	if err != nil {
		LogError(fmt.Sprintf("Failed to seek to offset: %v", err))
		return "", err
	}
	// Read specified length
	buf := make([]byte, length)
	n, err := file.Read(buf)
	if err != nil {
		LogError(fmt.Sprintf("Failed to read from file: %v", err))
		return "", err
	}
	return string(buf[:n]), nil
}

func DeleteLogFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		LogError(fmt.Sprintf("Error in removing file : %v", err))
		return err
	}
	fmt.Println("File deleted successfully")
	return nil
}
func main() {
	filename := "service_log.txt"
	err := CreateLogFile(filename, "Service started successfully.")
	if err != nil {
		log.Fatal("Error ", err)
	}

	content, err := ReadLogFile(filename)
	if err != nil {
		log.Fatal("Error in Reading ", err)
	}
	fmt.Println("content : ", content)
	AppendToLogFile(filename, "\nAppending new content ")

	err = ReportFileStats(filename)

	if err != nil {
		log.Fatal("Error in report ", err)
	}
	data, _ := ReadFromOffset(filename, 8, 15)
	fmt.Println("data ", data)
	err = DeleteLogFile("test.txt")
	if err != nil {
		log.Fatal("Error in Delete ", err)
	}
}
