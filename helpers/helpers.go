package helpers

import (
	"log"
	"os"
)

// WriteToLogfile : Write log to specified file
func WriteToLogfile(filename string, msq ...interface{}) {
	//create your file with desired read/write permissions
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	log.SetOutput(f)

	log.Println(msq)
}
