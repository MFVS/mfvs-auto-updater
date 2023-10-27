package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Define a struct to match the JSON structure
type Configuration struct {
	Repositories []string `json:"repositories"`
}

func main() {
	// Open and read the JSON file
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer configFile.Close()

	// Parse the JSON data into the Configuration struct
	var config Configuration
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	os.RemoveAll("repositories")
	os.Mkdir("repositories", 0777)

	// Iterate over repositories
	for _, url := range config.Repositories {
		log.Println("Cloning ", url)
		// Create a new Git command
		cmd := exec.Command("git", "clone", url, "--verbose", "--progress")
		cmd.Dir = "repositories"

		// Set the command's stdout to your program's stdout
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run the Git command
		err := cmd.Run()
		if err != nil {
			panic(fmt.Sprintln("[ERROR] GIT CLONE:", err))
		}
		// errMsg := fmt.Sprintln(err)
		// if errMsg == "Error running Git command: exit status 128" {
		// 	fmt.Println("already exists")
		// } else if err != nil {
		// 	fmt.Println("Error running Git command:", err)
		// }
	}
}
