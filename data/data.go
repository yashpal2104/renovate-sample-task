package data

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"

	// "log"
	"net/http"
	"os"
	"strings"

)

type Data struct {
	Fruit string `json:"fruit"`
	Size  string `json:"size"`
	Color string `json:"color"`
}

var FileParsedData []Data

func readJSON(path string) ([]byte, error) {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		//URL, fetch the data
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return io.ReadAll(resp.Body)
	} else {
		//  It's a file path, read the file
		return os.ReadFile(path)

	}
}
func FileData(){
// For a local file path
	filePath := "json_file.json" //the file path
	fileData, err := readJSON(filePath)
	key := flag.String("key", "", "Key to filter from the JSON")
	flag.Parse()

	if err != nil {
		fmt.Println("Error reading a file: ", err)
		return
	}

	err = json.Unmarshal(fileData, &FileParsedData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	for i, d := range FileParsedData {
		fmt.Printf("Data from file: %d. %+v\n", i+1, d)
	}

	// Filter key if provided
	if *key != "" {
		for i, d := range FileParsedData {
			switch *key {
			case "fruit":
				fmt.Printf("%d. Fruit: %s\n", i+1, d.Fruit)
			case "size":
				fmt.Printf("%d. Size: %s\n", i+1, d.Size)
			case "color":
				fmt.Printf("%d. Color: %s\n", i+1, d.Color)
			default:
				fmt.Println("Unknown key:", *key)
				return
			}
		}
	} else {
		// Pretty-print all
		prettyJSON, err := json.MarshalIndent(FileParsedData, "", "  ")
		if err != nil {
			fmt.Println("Error formatting JSON:", err)
			return
		}
		fmt.Println(string(prettyJSON))
	}
}