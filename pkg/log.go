package logging

import (
	"fmt"
	"log"
	"os"
	"path"
)

func Logging(message, prefix string) {
	currentDir, _ := os.Getwd()
	LOGFILE := path.Join(fmt.Sprintf("%s/logs", currentDir), "logs.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	lstFlags := log.Ldate | log.Lshortfile | log.Ltime
	if err != nil {
		log.Fatal(err)
		return
	}
	iLog := log.New(f, fmt.Sprintf("%s ", prefix), lstFlags)
	iLog.Println(message)
	defer f.Close()
}
