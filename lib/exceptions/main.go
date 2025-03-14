package exceptions

import (
	"log"
	"os"
	"strings"
	"time"
)

func HandleAnException(e string) {

	log.Println("Error log is: ", e)

	var file *os.File
	var err error
	fileName := "log.txt"

	curTime := time.Now().Format("2006-01-02")
	dirName := strings.Join([]string{curTime, "_ERRORlog"}, "")
	filePath := dirName + "/" + fileName

	if _, err = os.Stat(dirName); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			err = os.Mkdir(dirName, 0700)
			if err != nil {
				log.Println("cannot create a dir.")
				// os.Exit(1)
			}

			file, err = os.Create(filePath)
			if err != nil {
				log.Println("cannot create a file.")
				// os.Exit(1)
			}
		}
	}

	file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println("Cannot open a file.")
		// os.Exit(1)
	}

	// close a file and terminate the process
	defer func() {
		file.Close()
		// os.Exit(1)
	}()

	str := time.Now().Format(time.RFC850) + " :> " + e + " \n"
	_, err = file.WriteString(str)
	if err != nil {
		log.Println("Cannot write to the file.")
		// os.Exit(1)
	}
}

func HandleAnHttpExceprion() {
	status, message := definteHttpError()
	log.Printf("http exception: {message: %s, status %d} \n", message, status)
}

func definteHttpError() (int, string) {
	return 400, "Bad Request"
}
