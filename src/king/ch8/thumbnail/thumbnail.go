package thumbnail

import (
	"image"
	"github.com/gogo/protobuf/io"
	"image/jpeg"
	"os"
	"fmt"
	"path/filepath"
	"strings"
)

func Image(src image.Image)  {
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if ascpec := float64(xs) / float64(xy); ascpec < 1.0 {
		width = int(128 * ascpec)
	} else {
		height = int(123 * ascpec)
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0,0, width, height))

	for x := 0; x < width; x++  {
		for y := 0; y <height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}

	return dst
}

func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)

	return jpeg.Encode(w, dst, nil)
}

func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)

	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s : %s", infile, outfile, err)
	}
	return out.Close()
}

func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile)
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}