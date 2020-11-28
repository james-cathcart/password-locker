package service

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	. "pwlocker/model"
)

const (
	FilePath string = "/Users/jamescathcart/workspace/golang/go-pass/pwlocker-test"
)

type FileService struct {

}

func (fs *FileService) GetAll() []PasswordRecord {

	lines := readLines(FilePath)

	records := deserializedRecords(lines)

	return records
}

func (fs *FileService) LookupPassword(serviceName string) PasswordRecord {



}

func (fs *FileService) GetByServiceName(serviceName string) (bool, PasswordRecord) {

	records := fs.GetAll()

	var recordToReturn PasswordRecord
	foundRecord := false

	log.Printf("Performing lookup for %s\n", serviceName)
	for _, record := range records {

		if record.Service == serviceName {
			recordToReturn = record
			foundRecord = true
			break
		}

	}

	return foundRecord, recordToReturn
}

func readLines(filePath string) []string {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 1)

	// TODO remove manual slice doubling code, append() is smart enough to do this:
	//lineCount := 0
	for scanner.Scan() {

		lines = append(lines, scanner.Text())

		//lineCount++
		//
		//if lineCount > (len(lines)) {
		//	lines = append(
		//		lines,
		//		make([]string, lineCount*2)...,
		//	)
		//}
		//lines[lineCount-1] = scanner.Text()

	}

	// get the final slice
	return lines
}

func deserializedRecords(rawLines []string) []PasswordRecord {

	passwordRecordList := make([]PasswordRecord, len(rawLines))

	for i, line := range rawLines {
		recordPointer := new(PasswordRecord)
		json.Unmarshal([]byte(line), &recordPointer)
		passwordRecordList[i] = *recordPointer
	}

	return passwordRecordList
}
