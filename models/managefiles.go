package models

import (
	"fmt"
	"os"
	"path/filepath"
)

var synagogues_path string

func InitSynagoguesPath() {
	synagogues_path = "./files/synagogues/"
}

func CreateDir(synagogue_name string) bool {
	// Create main synagogue directory

	err := os.Mkdir(fmt.Sprint(synagogues_path, synagogue_name), 0755)
	return err == nil
}

func DirExist(synagogue_name string) bool {
	if _, err := os.Stat(fmt.Sprint(synagogues_path, synagogue_name)); !os.IsNotExist(err) {
		// path/to/whatever exists
		return true
	}
	return false
}

func DeleteDir(synagogue_name string) bool {
	err := os.Remove(fmt.Sprint(synagogues_path, synagogue_name))
	return err == nil

}

func WriteFile(synagogue_name, file_name string, content string) bool {
	myfile, err := os.Create(fmt.Sprint(synagogues_path, synagogue_name, "/", file_name))
	if err != nil {
		return false
	}
	myfile.WriteString(content)
	myfile.Close()
	return true

}

func ReadFile(synagogue_name, file_name string, content string) *[]byte {
	dat, err := os.ReadFile(fmt.Sprint(synagogues_path, synagogue_name, "/", file_name))
	if err != nil {
		return &[]byte{}
	}
	return &dat
}

func FileExist(synagogue_name, file_name string) bool {
	if _, err := os.Stat(fmt.Sprint(synagogues_path, synagogue_name, "/", file_name)); !os.IsNotExist(err) {
		// path/to/whatever exists
		return true
	}
	return false
}

func DeleteFile(synagogue_name, file_name string) bool {
	err := os.Remove(fmt.Sprint(synagogues_path, synagogue_name, "/", file_name))
	return err == nil
}

func GetAllFilesInDir(synagogue_name string) *[]string {
	var files []string

	root := fmt.Sprint(synagogues_path, synagogue_name)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return &files
}
