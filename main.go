package main

import (
	"flag"
	"fmt"
	"log"
	"main/task"
	"os"
	"path/filepath"
	"time"
)

var (
	srcFolder  string
	dstFolder  string
	filterName string
	taskName   string
	poolSize   int
)

func init() {
	flag.StringVar(&srcFolder, "src", "", "Source folder path")
	flag.StringVar(&dstFolder, "dst", "", "Destination folder path")
	flag.StringVar(&filterName, "filter", "", "Image filter to apply")
	flag.StringVar(&taskName, "task", "waitgrp", "Concurrent task repartition method")
	flag.IntVar(&poolSize, "poolsize", 4, "Size of channel pool")
	flag.Parse()
}

func main() {
	if srcFolder == "" || dstFolder == "" || filterName == "" {
		fmt.Println("Missing required parameters. Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	files, err := getFilesInFolder(srcFolder)
	if err != nil {
		log.Fatal("Error reading source folder:", err)
	}

	startTime := time.Now()

	switch taskName {
	case "waitgrp":
		task.ProcessImagesWaitGroup(files, filterName, dstFolder)
	case "channel":
		task.ProcessImagesChannel(files, filterName, dstFolder, poolSize)
	default:
		log.Fatal("Invalid task option. Available options: waitgrp, channel")
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Elapsed Time: %s\n", elapsedTime)
}

func getFilesInFolder(folderPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
