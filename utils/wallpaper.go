package utils

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"

	"github.com/qeesung/image2ascii/convert"
)


func ChangeWallpaperFromFile(file string) error {
	return exec.Command("osascript", "-e", `tell application "System Events" to tell every desktop to set picture to `+strconv.Quote(file)).Run()
}

func WallPaperLoop(dir string, fps int) error {
	// inifinite loop
	dir = RelativeDirToAbsDir(fmt.Sprintf("vault/%s",dir))
	fmt.Printf("Starting wallpaper loop in %s\n", dir)
	for {
		wallpaperLoop(dir, fps)
	}
}

func wallpaperLoop(dir string, fpt int) {
	number_of_files, _ := FolderFileCount(dir)
	loopCount := number_of_files
	if (number_of_files > 500) {
		loopCount = 500
	}
	for i := 1; i < loopCount; i++ {
		// file name 001, 002, 010
		file_name := fmt.Sprintf("%03d", i)
		path_ := dir + "/" + file_name + ".jpg"
		// set wallpaper
		error := ChangeWallpaperFromFile(path_)
//		if i % 100 == 0{
//			colAscii(path_)
//		}
		// sleep 1/60 of a second
		time.Sleep(time.Second / time.Duration(fpt))
		if error != nil {
			fmt.Println(error)
		} 

	}

}


// ImgToAscii: does image to ascii command with jp2a os command
func ImgToAscii(img_path string) error {
	// save the output of the 

	// return exec.Command("jp2a", img_path).Run()
	// save the output and print
	// clear terminal
	fmt.Print("\033[2J\033[1;1H")
	out, err := exec.Command("jp2a", img_path).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func colAscii(imageFilename string) error {
	// Create convert options
	convertOptions := convert.DefaultOptions
	convertOptions.Colored = true

	// Create the image converter
	converter := convert.NewImageConverter()
	fmt.Print(converter.ImageFile2ASCIIString(imageFilename, &convertOptions))
	return nil
}
