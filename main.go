package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

func visit(path string, file os.FileInfo, err error) error {
	fmt.Printf("Visiting: %s\n", path)

	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	_, str, err := image.Decode(fi)
	//	imgBounds := img.Bounds().Size()
	fmt.Printf("string: %s, error: %s\n", str, err)
	//	fmt.Printf("Image metrics W: %d H: %d", imgBounds.X, imgBounds.Y)
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
