package config

import (
	"io/ioutil"
	"os"
)

// config struct
type Config struct {
	// app name
	AppName string `json:"app_name"`
	// app version
	AppVersion string `json:"app_version"`
	// app author
	AppAuthor string `json:"app_author"`
	// app author email
	AppAuthorEmail string `json:"app_author_email"`
	// default vault folder
	VaultFolder string `json:"vault_folder"`
}

// NewConfig returns a new config
func NewConfig() *Config {
	// check vault folder
	checkVaultFolder("./vault")
	return &Config{
		AppName:        "gotomate",
		AppVersion:     "0.0.1",
		AppAuthor:      "Poggybitz",
		AppAuthorEmail: "poggybitz@poggybitz.com",
		VaultFolder:    "vault",
	}
}


func checkVaultFolder(vaultFolder string) error {
	// check if the vault exists
	_, err := ioutil.ReadDir(vaultFolder)
	if err != nil {
		// create new directory if it doesn't exist
		err = os.MkdirAll(vaultFolder, 0755)
	}
	return err
}