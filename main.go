package main

import (
	logging "GoWhich/pkg"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		log.Fatal("Fatal: No args")
		return
	}
	fileName := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, fileName)
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			if mode.IsRegular() {
				if mode&0111 != 0 {
					logging.Logging(fmt.Sprintf("Success find: %s | %s", fileName, fullPath), "goFinder")
					fmt.Println(fullPath)
					return
				}

			}
		}
	}

}
