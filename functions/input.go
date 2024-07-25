package ascii

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// handles command line inputs checking if correctly formatted
// checks if the banner file is modified and is not as it should be
// it downloads the banner file if different from original
// checks flags usage from CLI if has correct format
func Input(text string, banner string) (string, error) {
	if text == "" {
		return "", fmt.Errorf("empty strings not accepted")
	}
	// preset the checksum values of the files
	standardCheckSum := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	shadowCheckSum := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoyCheckSum := "092d0cde973bfbb02522f18e00e8612e269f53bac358bb06f060a44abd0dbc52"
	acCheckSum := "cb217d33c89b7320ebd39f18573ab7ab90c3cb9042a7702f5dd21833524e73b2"

	banner = strings.ToLower(banner)

	
	switch banner {
	case "standard":
		banner = "functions/resources/standard.txt"
	case "thinkertoy":
		banner = "functions/resources/thinkertoy.txt"
	case "shadow":
		banner = "functions/resources/shadow.txt"
	case "ac":
		banner = "functions/resources/ac.txt"
	default:
		banner = "functions/resources/standard.txt"
	}

	File, err := os.Open(banner)
	if err != nil {

		return "", fmt.Errorf("file not found")
	}
	defer File.Close()

	bannerTemp := sha256.New()
	if _, err := io.Copy(bannerTemp, File); err != nil {
		log.Fatal(err)
	}
	checkSum := string(fmt.Sprintf("%x", bannerTemp.Sum(nil)))

	if checkSum != standardCheckSum && checkSum != thinkertoyCheckSum && checkSum != shadowCheckSum && checkSum != acCheckSum {
		err:=Checkfiles(banner)
		if err!=nil{
			return "",fmt.Errorf("failed to connect successflully")
			}else{
				
				bannerFile, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	mapped := AsciiArtMap(string(bannerFile))
	result := Tab(text)
	asciiArt, err := AsciiCombine(result, mapped)
	if err != nil {
		return "", errors.New("invalid banner type")
	}
	return strings.Join(asciiArt, ""), nil
		}
		
	} 
	
		bannerFile, err := os.ReadFile(banner)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	
		mapped := AsciiArtMap(string(bannerFile))
		result := Tab(text)
		asciiArt, err := AsciiCombine(result, mapped)
		if err != nil {
			return "", errors.New("invalid banner type")
		}
		return strings.Join(asciiArt, ""), nil



}
