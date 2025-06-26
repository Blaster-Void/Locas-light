// Locas Light Begin of all -> (MUSB) <- <- <-
// MUSB Make Ur Self Better
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorCyan    = "\033[36m"
	colorWhite   = "\033[37m"
	colorReset   = "\033[0m"
)

func main() {
	fmt.Println(colorGreen, "Locas Light", colorReset)
	if len(os.Args) == 2 {
		op := os.Args[1]
		path, err := os.Getwd()

		if err != nil {
			fmt.Println(err)
		}
		path = fmt.Sprint(path, "/", op)
		files, err := os.ReadDir(path)
		for _, file := range files {
			if file.IsDir() {
				fmt.Println(file.Name(), colorMagenta, "<-DIR", colorReset)
			} else {
				continue
			}
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			} else {
				exes := filepath.Ext(file.Name())
				fmt.Println(file.Name(), colorYellow, "<-", filedetect(exes), colorReset)
			}
		}
	} else {
		path, err := os.Getwd()

		if err != nil {
			fmt.Println(err)
		}
		files, err := os.ReadDir(path)
		for _, file := range files {
			if file.IsDir() {
				fmt.Println(file.Name(), colorMagenta, "<-DIR", colorReset)
			} else {
				continue
			}
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			} else {
				exes := filepath.Ext(file.Name())
				fmt.Println(file.Name(), colorYellow, "<-", filedetect(exes), colorReset)
			}
		}
	}
}
func filedetect(exes string) string {
	if exes == ".go" {
		return "GOLANG"
	} else if exes == ".c" {
		return "CLANG"
	} else if exes == ".mp4" {
		return "Video"
	} else if exes == ".mp3" || exes == ".m4a" {
		return "Music"
	} else {
		return exes
	}
}
