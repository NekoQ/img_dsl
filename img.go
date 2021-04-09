package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"img/parser"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/disintegration/gift"
)

type ImgListener struct {
	*parser.BaseImgListener

	g *gift.GIFT

	src string
	dst string
}

func main() {

	args := os.Args[1:]
	f, err := ioutil.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	var l ImgListener
	l.g = gift.New()

	// Setup the input
	is := antlr.NewInputStream(string(f))

	// Create the Lexer
	lexer := parser.NewImgLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewImgParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&l, p.Start())
	executeActions(l)
}

func executeActions(l ImgListener) {
	fmt.Println("Start executing actions")

	src := loadImage(l.src)
	fmt.Println("Opening file: " + l.src + " ✓")

	dst := image.NewRGBA(l.g.Bounds(src.Bounds()))
	l.g.Draw(dst, src)
	fmt.Println("Applying filters ✓")

	saveImage(l.dst, dst)
	fmt.Println("Creating new file: " + l.dst + " ✓")
}

func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
	}
	return img
}

func saveImage(filename string, img image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		log.Fatalf("png.Encode failed: %v", err)
	}
}

// Open and Export file ---------------

func (l *ImgListener) ExitOpenFile(c *parser.OpenFileContext) {
	l.src = strings.Trim(c.FileName().GetText(), "\"")
}

func (l *ImgListener) ExitExport(c *parser.ExportContext) {
	l.dst = strings.Trim(c.FileName().GetText(), "\"")
}

// Actions ---------------------

func (l *ImgListener) ExitRotate(c *parser.RotateContext) {
	params := c.NUMBER().GetText()
	angle, _ := strconv.ParseFloat(params, 32)
	// Check for error
	l.g.Add(gift.Rotate(float32(angle), color.Transparent, gift.CubicInterpolation))
}

func (l *ImgListener) ExitFlipY(c *parser.FlipYContext) {
	l.g.Add(gift.FlipVertical())
}

func (l *ImgListener) ExitFlipX(c *parser.FlipXContext) {
	l.g.Add(gift.FlipHorizontal())
}

func (l *ImgListener) ExitCrop(c *parser.CropContext) {
	params := c.AllNUMBER()
	x0, _ := strconv.Atoi(params[0].GetText())
	y0, _ := strconv.Atoi(params[1].GetText())
	x1, _ := strconv.Atoi(params[2].GetText())
	y1, _ := strconv.Atoi(params[3].GetText())

	l.g.Add(gift.Crop(image.Rect(x0, y0, x1, y1)))

}
