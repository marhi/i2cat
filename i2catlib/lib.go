package i2catlib

import (
	"io"
	"image"
	"bytes"
	"image/png"
	_ "image/jpeg"
	_ "image/gif"
	"os"
	"fmt"
	"encoding/base64"
)

func ReadIn(in io.Reader) ([]byte, error) {
	img, _, err := image.Decode(in)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err = png.Encode(buf, img); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func isScreen() bool {
	return os.Getenv("TERM") == "screen"
}

func encode(src []byte, dest io.Writer) error {
	enc := base64.NewEncoder(base64.StdEncoding, dest)
	_, err := enc.Write(src)
	return err
}

func stdout(output string) error {
	_, err := fmt.Fprint(os.Stdout, output)
	return err
}

func PrintImg(img io.Reader) error {
	buf, err := ReadIn(img)
	if err != nil {
		return err
	}

	isScr := isScreen()
	if isScr {
		stdout( "\033Ptmux;\033")
	}

	stdout("\033]1337;File=;inline=1:")
	if err = encode(buf, os.Stdout); err != nil {
		return err
	}

	stdout("\a")
	if isScr {
		stdout("\033\\")
	}
	stdout("\n")

	return nil
}
