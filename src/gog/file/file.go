package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"gog/conf"
)

func writeFile(path string, data []byte) bool {
	path = conf.FILE_PATH + "/" + path + ".go"
	if _, err := ioutil.ReadDir(path); err != nil {
		if err = os.Mkdir(conf.FILE_PATH, os.ModePerm); err == nil {
			log.Fatal(err)
			return false
		}
	}

	if ioutil.WriteFile(path, data, 0644) == nil {
		fmt.Println("'success")
		return true
	}
	fmt.Println("'error")
	return false
}
