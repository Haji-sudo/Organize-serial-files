package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var path string
	fmt.Print("Enter the path : ")
	fmt.Scanln(&path)
	files := GetSerialFiles(path)

	CreateFolders(files, path)
	MoveFiles(files, path)
}

func GetSerialFiles(dir string) []string {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if info.IsDir() == false {
			filename := filepath.Base(path)
			if GetSeason(filename) != "" {
				if GetResolution(filename) != "" {
					files = append(files, filename)
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return files
}

func GetSeason(filename string) string {
	getSeason := regexp.MustCompile(`([Ss](\d+))`)
	season := getSeason.FindStringSubmatch(filename)
	if len(season) > 0 {
		st, _ := strconv.Atoi(season[2])
		s := fmt.Sprintf("%02d", st)
		return "S" + s
	}
	return ""
}

func GetResolution(filename string) string {
	getResolution := regexp.MustCompile(`((480|1080|720|2160)[Pp])`)
	resolution := getResolution.FindStringSubmatch(filename)
	if len(resolution) > 0 {
		resolution[0] = strings.TrimSpace(resolution[0])
		return strings.ToLower(resolution[0])
	}
	return ""
}

func GetEncode(filename string) string {
	var text = ""
	getEncode1 := regexp.MustCompile(`[Xx]265`)
	getEncode2 := regexp.MustCompile(`(10.?[Bb][Ii][Tt])`)
	getEncode3 := regexp.MustCompile(`[Hh][Ee][Vv][Cc]`)
	x256 := getEncode1.FindStringSubmatch(filename)
	bit10 := getEncode2.FindStringSubmatch(filename)
	hevc := getEncode3.FindStringSubmatchIndex(filename)

	if len(x256) > 0 {
		text += x256[0]
	}
	if len(bit10) > 0 || len(hevc) > 0 {
		if text == "" {
			text += "10Bit"
		} else {
			text += " 10Bit"
		}
	}
	return text
}

func GetType(filename string) string {
	isDubbled := regexp.MustCompile(`[Dd][Uu][Bb][Bb]?[Ll]?[Ee][Dd]?`)
	isSoftSub := regexp.MustCompile(`[Ss][Oo|Uu][Ff|Bb][Tt]?[Bb]?[Ee]?[Dd]?`)
	if isDubbled.MatchString(filename) {
		return "Dubbled"
	}
	if isSoftSub.MatchString(filename) {
		return "SoftSub"
	}
	return ""
}

func GetNewPath(filename string, path string) string {
	newpath := path
	if season := GetSeason(filename); season != "" {
		newpath += "/" + season
	}
	if filetype := GetType(filename); filetype != "" {
		newpath += "/" + filetype
	}
	if res := GetResolution(filename); res != "" {
		newpath += "/" + res
	}
	if en := GetEncode(filename); en != "" {
		newpath += " " + en
	}

	return newpath
}

func CreateFolders(files []string, path string) {

	for _, file := range files {
		newpath := GetNewPath(file, path)
		os.MkdirAll(newpath, os.ModeDir)
	}
}

func MoveFiles(files []string, path string) bool {
	for _, file := range files {
		oldpath := path + "/" + file
		newpath := GetNewPath(file, path) + "/" + file
		err := os.Rename(oldpath, newpath)
		if err != nil {
			return false
		}
	}
	return true
}
