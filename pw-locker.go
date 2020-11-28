package main

import (
	"context"
	"fmt"
	"log"
	"os"
	. "pwlocker/model"
	"strconv"
)

const (
	AppName      string = "Password Locker"
	FilePath     string = "/Users/jamescathcart/workspace/golang/go-pass/pwlocker-test"
	ShowAll      string = "-s"
	CreateRecord string = "-c"
)


func main() {

	log.Printf("%s application started.\n", AppName)

	args := os.Args
	log.Printf("Received %d args\n", len(args))

	if len(args) < 2 || args[1] == "" {
		panic("Invalid input! Must provide the name of a service.")
	}

	switch args[1] {
	case ShowAll:
		log.Printf("Displaying all services\n")
		showAllServices()
		break
	case CreateRecord:
		log.Printf("Creating new record...\n")
		createServiceRecord()
		addCredentials()
		break
	default:
		log.Printf("No options detected, performing default lookup...\n")
		serviceQuery(args[1])
	}

}

func serviceQuery(serviceName string) {

	// get lines
	lines := readLines(FilePath)

	// deserialize to entities
	passwordRecords := deserializedRecords(lines)

	// get the appropriate record
	status, passwordRecord := findServiceRecords(serviceName, passwordRecords)

	// if no record was found, exit
	if !status {
		log.Printf("Did not find record!")
		os.Exit(0)
	}
	// check to see if there are multiple accounts for the service
	numberOfAccounts := len(passwordRecord.Credentials)

	if numberOfAccounts > 1 {

		log.Printf("Found %d accounts, presenting selection to user\n", numberOfAccounts)
		fmt.Println("\nSelect account: ")

		for i, v := range passwordRecord.Credentials {
			fmt.Printf("\t%d -> %s\n", i, v.Username)
		}

		var userInput string
		fmt.Scanln(&userInput)

		if userInput != "" {
			credentialsIndex, _ := strconv.Atoi(userInput)
			if credentialsIndex < len(passwordRecord.Credentials) {
				printAccountRecord(passwordRecord, credentialsIndex)
			} else {
				log.Fatal("Invalid selection!")
			}
		} else {
			panic("Invalid input!")
		}

	} else if numberOfAccounts == 1 {
		printAccountRecord(passwordRecord, 0)
	} else {
		panic("No records could be retrieved!")
	}

}

//func findServiceRecords(service string, passwordRecords []PasswordRecord) (bool, PasswordRecord) {
//
//	var passwordRecord PasswordRecord
//	foundRecord := false
//
//	log.Printf("Performing lookup for %s\n", service)
//	for k, v := range passwordRecords {
//
//		if v.Service == service {
//			passwordRecord = passwordRecords[k]
//			foundRecord = true
//			break
//		}
//
//	}
//
//	return foundRecord, passwordRecord
//}

//func deserializedRecords(rawLines []string) []PasswordRecord {
//
//	passwordRecordList := make([]PasswordRecord, len(rawLines))
//
//	for i, line := range rawLines {
//		recordPointer := new(PasswordRecord)
//		json.Unmarshal([]byte(line), &recordPointer)
//		passwordRecordList[i] = *recordPointer
//	}
//
//	return passwordRecordList
//}

//func readLines(filePath string) []string {
//
//	file, err := os.Open(filePath)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//
//	lines := make([]string, 1)
//
//	lineCount := 0
//	for scanner.Scan() {
//
//		lineCount++
//
//		if lineCount > (len(lines)) {
//			lines = append(
//				lines,
//				make([]string, lineCount*2)...,
//			)
//		}
//
//		lines[lineCount-1] = scanner.Text()
//
//	}
//
//	// get the final slice
//	return lines[0:lineCount]
//}

func printAccountRecord(record PasswordRecord, credentialsIndex int) {

	fmt.Printf(
		"\n%s -> %s -> %s\n",
		record.Service,
		record.Credentials[credentialsIndex].Username,
		record.Credentials[credentialsIndex].Password,
	)

}

func showAllServices() {

	// get all raw JSON lines
	// Unmarshall raw JSON lines to PasswordRecord type
	// collect PasswordRecord entities to list
	// iterate through list and display Service names

	lines := readLines(FilePath)

	records := deserializedRecords(lines)

	if len(records) < 1 {
		log.Printf("No records to display\n")
		return
	}

	for _, currentRecord := range records {
		fmt.Printf("%s\n", currentRecord.Service)
	}
}

func createServiceRecord() {
	context.TODO()
}

func addCredentials() {
	context.TODO()
}

func writeLines(filePath string, lines []string) {
	context.TODO()
}

func greeting() {
	fmt.Printf("%s\n| %-25s |\n%s\n", "+---------------------------+", AppName, "+---------------------------+")
}
