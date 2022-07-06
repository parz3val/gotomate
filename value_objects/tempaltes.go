package value_objects

import "fmt"

// var command = "osascript -e 'tell application "Finder" to set desktop picture to POSIX file "/Users/harriskunwar/Pictures/Att.jpg""
// "
/*
This kind of how the command is supposed to look like
tell application "Finder"
 set desktop picture to (POSIX file "/Users/harriskunwar/Pictures/1.jpg")
 end tell
*/
func buildFinderString(path string) string {
	return fmt.Sprintf(
		"tell application \"Finder\"\n" +
			"set desktop picture to (POSIX file \"%s\")\n" +
			"end tell",path,
	)
}
// ChangeWallPaper is a function that changes the wallpaper of the Mac
func ChangeWallPaperCommand(path string) ( Command, error ){
	cmd := Command{
		App: "osascript",
		Args:[]string {
			"-e",
			buildFinderString(path),		},
	}
	return cmd, nil
}


func Say(text string) ( Command, error ){
	script_ := fmt.Sprintf("say \"%s\"", text)
	cmd := Command{
		App: "osascript",
		Args:[]string {
			"-e",
			script_,
			// "say \"Hello from default!\"",
		},
	}
	return cmd, nil
}