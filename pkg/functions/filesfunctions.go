package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateDir(dirName string) error {
	// Create main synagogue directory

	err := os.Mkdir(fmt.Sprint("./", dirName), 0755)
	if err != nil {
		return err
	}
	return nil
}

func DirExist(dirName string) bool {
	if _, err := os.Stat(fmt.Sprint("./", dirName)); !os.IsNotExist(err) {
		return true
	}
	return false
}

func DeleteDir(dirName string) error {
	err := os.Remove(fmt.Sprint("./", dirName))
	return err

}

func WriteFile(dirName, file_name string, content string) bool {
	myfile, err := os.Create(fmt.Sprint("./", dirName, "/", file_name))
	if err != nil {
		return false
	}
	myfile.WriteString(content)
	myfile.Close()
	return true

}

func ReadFile(dirName, file_name string) *[]byte {
	dat, err := os.ReadFile(fmt.Sprint("./", dirName, "/", file_name))
	if err != nil {
		return &[]byte{}
	}
	return &dat
}

func FileExist(dirName, file_name string) bool {
	if _, err := os.Stat(fmt.Sprint("./", dirName, "/", file_name)); !os.IsNotExist(err) {
		// path/to/whatever exists
		return true
	}
	return false
}

func DeleteFile(dirName, file_name string) error {
	err := os.Remove(fmt.Sprint("./", dirName, "/", file_name))
	return err
}

func DeleteAllFiles(dirName string) error {
	/*
		err := os.RemoveAll(fmt.Sprint(SYNAGOGUESPATH, synagogue_name, "/"))
	*/

	files, err := filepath.Glob(fmt.Sprint("./", dirName, "/", "*"))
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

func GetAllFilesInDir(dirName string) *[]string {
	var files []string

	err := filepath.Walk(fmt.Sprint("./", dirName), func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return &files
}
