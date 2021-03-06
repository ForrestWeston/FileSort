package main

import (
	"flag"
	"fmt"
	"github.com/pierrre/archivefile/zip"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path"
	"path/filepath"
)

func visit(filePath string, file os.FileInfo, err error) error {

	if file.IsDir() {
		return nil
	}
	fileExt := path.Ext(filePath)
	fmt.Printf("Visiting: %s\n", filePath)
	fmt.Printf("File Ext: %s\n", fileExt)

	switch fileExt {
	case ".png", ".jpg", ".jpeg", ".gif", ".svg":
		fmt.Println("Calling imageHandle")
		err := imageHandle(filePath)
		if err != nil {
			return err
		}
	case ".zip":
		fmt.Println("Calling zipHandle")
		err := zipHandle(filePath)
		if err != nil {
			return err
		}
	default:
		return nil
	}

	return nil
}

func zipHandle(filePath string) error {

	destPath := os.Getenv("HOME") + "/Downloads/"
	err := zip.UnarchiveFile(filePath, destPath, nil)
	if err != nil {
		return err
	}
	err = os.Remove(filePath)
	return nil
}

func imageHandle(filePath string) error {

	// open image file for processing
	imgFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	// decode the image
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}

	// for image resolution
	imgBound := img.Bounds().Size()
	// directory named by image resolution (ir)
	irDir := fmt.Sprintf("%dx%d/", imgBound.X, imgBound.Y)
	// full path to dir image will be moved to
	fullPath := os.Getenv("HOME") + "/Pictures/" + irDir

	// create the dir for the image
	err = os.MkdirAll(fullPath, 0777)
	if err != nil {
		return err
	}

	// seperate image name from path
	_, imageName := path.Split(imgFile.Name())

	// move image into new dir
	err = os.Rename(imgFile.Name(), fullPath+imageName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
