package main

import (
	"github.com/jacobgc/AoC-2022/day-7/part-1/internal"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read File
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file: " + err.Error())
	}
	dataString := string(dat)
	dataStringArray := strings.Split(dataString, "\r\n")

	rootFs := internal.FileSystem{}
	rootFs.RootFolder = &internal.Folder{
		Name:      "/",
		TotalSize: 0,
		Files:     []*internal.File{},
		Folders:   make(map[string]*internal.Folder),
	}

	var cwd []string

	for _, line := range dataStringArray {
		if strings.HasPrefix(line, "$") { // Command
			if strings.HasPrefix(line, "$ cd") { // Change directory
				if line == "$ cd /" { // Ignore cding into /
					continue
				}
				folderToMoveTo := strings.Split(line, "cd ")[1]
				if folderToMoveTo == ".." {
					cwd = cwd[:len(cwd)-1]
				} else {
					rootFs.AddFolder(cwd, folderToMoveTo)
					cwd = append(cwd, folderToMoveTo)
				}
			} // Ignore other commands
		} else {
			if strings.HasPrefix(line, "dir ") {
				continue // Ignore dir commands as we make folders when cding into them
			} else {
				fileContents := strings.Split(line, " ")
				fileSize, _ := strconv.Atoi(fileContents[0])
				fileName := fileContents[1]
				fileToAdd := &internal.File{
					Name: fileName,
					Size: fileSize,
				}
				rootFs.AddFile(cwd, fileToAdd)
			}
		}
	}

	totalSize := 0

	addFolders(&totalSize, rootFs.RootFolder.Folders)

	println(totalSize)

	println("Done")
}

func addFolders(totalSize *int, folders map[string]*internal.Folder) {
	for _, folder := range folders {
		if folder.TotalSize <= 100000 {
			*totalSize += folder.TotalSize
		}
		if len(folder.Folders) != 0 {
			addFolders(totalSize, folder.Folders)
		}
	}
}
