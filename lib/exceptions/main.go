package exceptions

import (
	"log"
	"os"
	"strings"
	"time"
)

func HandleAnException(e string) {

	log.Println("Error log is: Unknown blockchain")

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
				os.Exit(1)
			}

			file, err = os.Create(filePath)
			if err != nil {
				log.Println("cannot create a file.")
				os.Exit(1)
			}
		}
	}

	file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println("cannot open a file.")
		os.Exit(1)
	}

	defer func() {
		file.Close()
		// os.Exit(1)
	}()

	str := time.Now().Format(time.RFC850) + " :> " + e + " \n"
	_, err = file.WriteString(str)
	if err != nil {
		log.Println("cannot write to the file.")
		os.Exit(1)
	}
}
