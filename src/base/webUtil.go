package base

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

func LoadConfigFile() []byte {
	dir, _ := os.Getwd()
	file := path.Join(dir, "/application.yaml")
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}
	return b
}
