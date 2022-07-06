package main

import (
	"fmt"
	"log"

	"github.com/poggybitz/gotomate/config"
	"github.com/poggybitz/gotomate/utils"
)


func main() {
	// ask user if they want to use existing wallpaper 
	// or download the new one

	config := config.NewConfig()

	// print the new config
	fmt.Println(config)
	// check if the vault is empty
	if !utils.DirectoryEmpty(config.VaultFolder){
		// Print the emoji of wallpaper
		fmt.Println("🖼🖼🖼🖼 GO WALLIE 🖼🖼🖼🖼🖼🖼🖼")
		
		// get list of directories from vault
		directories, _ := utils.GetDirectories(config.VaultFolder)

		// print the directories
		for _, directory := range directories {
			// rocket emoji
			fmt.Println("🚀🚀🚀🚀", directory, "🚀🚀🚀🚀")
		}


	}
	name := prompt()
	go utils.Mp4ToFrames(fmt.Sprintf("%s.mkv",name),name,20)
	// if OK != nil {
	// 	log.Fatal(OK)
	// }
	// wait for 10 seconds for the file lock to remove
	fmt.Printf("The name of the file is %s.mkv",name);
	utils.Wait(10)
	error := utils.WallPaperLoop(name, 5)
	if error != nil {
		log.Println(error)
	}
}

func prompt() string{
	fmt.Println("Enter the name of the wallpaper: ")
	var wallpaperName string
	// take in the name of the wallpaper
	fmt.Scanln(&wallpaperName)
	return wallpaperName
}

func onExit() {
	// clear the vault folder 
	// utils.ClearVaultFolder()
	fmt.Println("Exiting...")
}
