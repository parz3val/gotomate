package value_objects

import "fmt"

// var command = "osascript -e 'tell application "Finder" to set desktop picture to POSIX file "/Users/harriskunwar/Pictures/Att.jpg""
// "

// ChangeWallPaper is a function that changes the wallpaper of the Mac
func ChangeWallPaperCommand(path string) ( string, error ){
	command := fmt.Sprintf("osascript -e 'tell application \"Finder\" to set desktop picture to POSIX file \"%s\"'", path)
	return command, nil
}
