package filter

import (
	"fmt"
	"image"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

func ApplyFilter(filename string, filter string, dstFolder string) error {
	srcImg, err := imaging.Open(filename)
	if err != nil {
		return fmt.Errorf("Error opening image %s: %s", filename, err)
	}

	var dstImg *image.NRGBA

	switch strings.ToLower(filter) {
	case "grayscale":
		dstImg = imaging.Grayscale(srcImg)
	case "blur":
		dstImg = imaging.Blur(srcImg, 3.5)
	default:
		return fmt.Errorf("Invalid filter option: %s", filter)
	}

	outputPath := getOutputPath(filename, dstFolder)
	err = imaging.Save(dstImg, outputPath)
	if err != nil {
		return fmt.Errorf("Error processing image %s: %s", filename, err)
	}

	return nil
}

func getOutputPath(filename string, dstFolder string) string {
	base := filepath.Base(filename)
	return filepath.Join(dstFolder, base)
}
