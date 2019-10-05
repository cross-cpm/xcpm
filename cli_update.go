package main

import (
	"path/filepath"

	"github.com/codeskyblue/go-sh"
)

func doCliUpdate() error {

	home, err := Home()
	if err != nil {
		return err
	}

	libs_path := filepath.Join(home, GLOBAL_PACKAGE_LIBS_PATH)
	libs_repo := GLOBAL_PACKAGE_LIBS_REPO

	if !FileExist(libs_path) {
		err = sh.Command("git", "clone", libs_repo, libs_path).Run()
		if err != nil {
			return err
		}
	}

	err = sh.NewSession().
		SetDir(libs_path).
		Command("git", "pull").
		Run()
	if err != nil {
		return err
	}

	return nil
}
