package ascii

import (
	"fmt"
	"os"
)

// downloads the generated ascii art
func Download(s string, filename string) error {
	// Create the new file
	localFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating new file:", err)
		return err
	}
	defer localFile.Close()

	// Write the normalized content to the new local file
	_, err = localFile.WriteString(s)
	if err != nil {
		fmt.Println("Error writing content to new file:", err)
		return err
	}

	return nil
}
