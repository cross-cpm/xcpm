package main

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"

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

func SaveYaml(filename string, content interface{}) error {
	f, err := os.Create(filename)
	if err != nil {
		log.Println("create file failed!", filename, err)
		return err
	}
	defer f.Close()

	err = yaml.NewEncoder(f).Encode(content)
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

// Home returns the home directory for the executing user.
//
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
