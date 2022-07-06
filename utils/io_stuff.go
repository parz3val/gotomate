package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func FolderFileCount(path string) (int, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 0, err
	}
	return len(files), nil
}

func DirectoryEmpty(path string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false
	}
	return len(files) == 0
}

// func RelativeToAbsolutePath(path string) string {
// 	return fmt.Sprintf("%s/%s", os.Getenv("HOME"), path)
// }

func RelativeDirToAbsDir(dir_ string) string {
	// get the absolute path from the dir_
	abs_dir_path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/%s", abs_dir_path, dir_)
}

func Wait(seconds int) {
	fmt.Println("Waiting for ", seconds, " seconds")
	for i := 0; i < seconds; i++ {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	fmt.Println("")
}

func GetDirectories(path string) ([]string, error) {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var dirs_ []string
	for _, dir := range dirs {
		if dir.IsDir() {
			dirs_ = append(dirs_, dir.Name())
		}
	}
	return dirs_, nil
}