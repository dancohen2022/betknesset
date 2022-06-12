package models

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateDir(synagogue_name string) error {
	// Create main synagogue directory

	err := os.Mkdir(fmt.Sprint(SYNAGOGUESPATH, synagogue_name), 0755)
	if err != nil {
		return err
	}
	err = os.Mkdir(fmt.Sprint(SYNAGOGUESPATH, synagogue_name, CONFIGPATH), 0755)
	return err
}

func DirExist(synagogue_name string) bool {
	if _, err := os.Stat(fmt.Sprint(SYNAGOGUESPATH, synagogue_name)); !os.IsNotExist(err) {
		// path/to/whatever exists
		return true
	}
	return false
}

func DeleteDir(synagogue_name string) error {
	err := os.Remove(fmt.Sprint(SYNAGOGUESPATH, synagogue_name))
	return err

}

func WriteFile(synagogue_name, file_name string, content string) bool {
	myfile, err := os.Create(fmt.Sprint(SYNAGOGUESPATH, synagogue_name, "/", file_name))
	if err != nil {
		return false
	}
	myfile.WriteString(content)
	myfile.Close()
	return true

}

func ReadFile(synagogue_name, file_name string) *[]byte {
	dat, err := os.ReadFile(fmt.Sprint(SYNAGOGUESPATH, synagogue_name, "/", file_name))
	if err != nil {
		return &[]byte{}
	}
	return &dat
}

func FileExist(synagogue_name, file_name string) bool {
	if _, err := os.Stat(fmt.Sprint(SYNAGOGUESPATH, synagogue_name, "/", file_name)); !os.IsNotExist(err) {
		// path/to/whatever exists
		return true
	}
	return false
}

func DeleteFile(synagogue_name, file_name string) error {
	err := os.Remove(fmt.Sprint(SYNAGOGUESPATH, synagogue_name, "/", file_name))
	return err
}

func DeleteAllFiles(synagogue_name string) error {
	/*
		err := os.RemoveAll(fmt.Sprint(SYNAGOGUESPATH, synagogue_name, "/"))
	*/

	files, err := filepath.Glob(filepath.Join(SYNAGOGUESPATH, synagogue_name, "/", "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return err
}

func GetAllFilesInDir(synagogue_name string) *[]string {
	var files []string

	root := fmt.Sprint(SYNAGOGUESPATH, synagogue_name)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return &files
}
