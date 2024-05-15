package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/kolesa-team/go-webp/decoder"
	"github.com/kolesa-team/go-webp/webp"
	"golang.org/x/image/bmp"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"
)

var formatSupported = []string{"jpg", "jpeg", "png", "webp", "bmp"}

type DecodeFunc func(io.Reader) (image.Image, error)

func Decode(b []byte, decodeFunc DecodeFunc) image.Image {
	img, err := decodeFunc(bytes.NewReader(b))
	if err != nil {
		panic("cant decode image")
	}

	return img
}

func main() {
	formats := flag.String("to", "", "input many format image")
	filename := flag.String("file", "", "input name image")
	flag.Parse()

	if *formats == "" || *filename == "" {
		fmt.Println("\nUsage: amba --file <filename> --to <format>\n\nFormat supported:\n -png\n -jpg\n -jpeg\n -webp")
		return
	}

	splits := strings.Split(strings.ToLower(*formats), ",")

	file, err := os.Open(*filename)
	if err != nil {
		return
	}

	name := strings.Join(strings.Split(file.Name(), ".")[:len(strings.Split(file.Name(), "."))-1], "")

	byteFile, err := io.ReadAll(file)
	if err != nil {
		return
	}

	var decode image.Image
	switch http.DetectContentType(byteFile) {
	case "image/jpeg":
		decode = Decode(byteFile, jpeg.Decode)
	case "image/jpg":
		decode = Decode(byteFile, jpeg.Decode)
	case "image/png":
		decode = Decode(byteFile, png.Decode)
	case "image/webp":
		dec, err := webp.Decode(bytes.NewReader(byteFile), &decoder.Options{})
		if err != nil {
			panic(err.Error())
		}
		decode = dec
	case "image/bmp":
		decode = Decode(byteFile, bmp.Decode)
	default:
		panic("unknown image type")
	}

	for _, split := range splits {
		if slices.Contains(formatSupported, split) {
			buf := new(bytes.Buffer)
			switch split {
			case "png":
				err := png.Encode(buf, decode)
				if err != nil {
					panic("cant encode image")
				}
			case "jpg":
				err := jpeg.Encode(buf, decode, nil)
				if err != nil {
					panic("cant encode image")
				}
			case "jpeg":
				err := jpeg.Encode(buf, decode, nil)
				if err != nil {
					panic("cant encode image")
				}
			case "webp":
				err := webp.Encode(buf, decode, nil)
				if err != nil {
					panic("cant encode image")
				}
			case "bmp":
				err := bmp.Encode(buf, decode)
				if err != nil {
					panic("cant encode image")
				}
			}
			err := os.WriteFile(name+"."+split, buf.Bytes(), 0666)
			if err != nil {
				panic(err.Error())
			}
		}
	}
}
