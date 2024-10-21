package exceptions

import (
	"log"
	"os"
	"strings"
	"time"
)

func HandleAnException(e string) {

	log.Println("Error log is: Unknown blockchain")

	curTime := time.Now().Format("2006-01-02")
	dirName := strings.Join([]string{curTime, "_errorLog"}, "")
	filePath := dirName + "/log.txt"

	err := os.Mkdir(dirName, 0700)
	if err != nil {
		log.Println("cannot create a dir.")
		os.Exit(1)
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Println("cannot create a file.")
		os.Exit(1)
	}
	defer func() {
		file.Close()
		os.Exit(1)
	}()

	_, err = file.WriteString(e)
	if err != nil {
		log.Println("cannot write to the file.")
		os.Exit(1)
	}
}
