package main

import (
	"log"
	"os"
	"strings"

	"github.com/kardianos/osext"
)

func getCurrentDirName() string {
	dir, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	arr := strings.Split(dir, string(os.PathSeparator))
	return arr[len(arr)-1]
}

// CheckErrFatal Prints Error In StdErr
func CheckErrFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
