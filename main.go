package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path"
	"path/filepath"
)

func visit(path string, file os.FileInfo, err error) error {

	if file.IsDir() {
		return nil
	}

	fmt.Printf("Visiting: %s\n", path)

	fi, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fi.Close()

	err = imageHandle(fi)

	return nil
}

func imageHandle(imgFile *os.File) error {

	//open the image for processing
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
