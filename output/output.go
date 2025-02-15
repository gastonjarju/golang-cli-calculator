package output

import (
	"log"
	"os"
	"path/filepath"
)


func LogResult (result float64) float64 {
	dir, err := os.Getwd()
	
	if (err != nil) {
		log.Fatal("Error getting workig directory:", err)
	}

	ouputDir := filepath.Join(dir, "output")
	filepath := filepath.Join(ouputDir, "result.txt")

	file, err := os.OpenFile(filepath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if (err != nil){
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.SetFlags(0)
	log.Println("The Calculation Result is:" , result)

	return result

}