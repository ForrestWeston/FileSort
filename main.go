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

	img, str, err := image.Decode(fi)
	if err != nil {
		return err
	}

	imgBounds := img.Bounds().Size()
	err = makeImgDir(imgBounds.X, imgBounds.Y, path)
	if err != nil {
		return err
	}
	fmt.Printf("string: %s, error: %s\n", str, err)
	fmt.Printf("Image metrics W: %d H: %d\n", imgBounds.X, imgBounds.Y)

	return nil
}

func makeImgDir(w int, h int, filePath string) error {

	// directory named by image resolution (ir)
	irDir := fmt.Sprintf("%dx%d/", w, h)

	// full path to dir image will be moved to
	fullPath := os.Getenv("HOME") + "/Pictures/" + irDir

	// create the dir for the image
	err := os.MkdirAll(fullPath, 0777)
	if err != nil {
		return err
	}

	// seperate image name from path
	_, imageName := path.Split(filePath)

	// move image into new dir
	err = os.Rename(filePath, fullPath+imageName)
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
