package main

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func LoadYaml(filename string, content interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("open file failed!", filename, err)
		return err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(content)
	if err != nil {
		log.Println("decode file failed!", filename, err)
		return err
	}

	return nil
}

func FileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}
