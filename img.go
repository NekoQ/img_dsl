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

type Image struct {
	g *gift.GIFT

	src []image.Image

	fileNames []string
}

type ImgListener struct {
	*parser.BaseImgListener

	images       map[string]Image
	currentImage string
	g            *gift.GIFT
}

func main() {

	args := os.Args[1:]
	f, err := ioutil.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	var l ImgListener
	l.images = make(map[string]Image)
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
}

func (s *ImgListener) VisitErrorNode(node antlr.ErrorNode) {
	os.Exit(1)
}

// Open and Export file ---------------

func (l *ImgListener) ExitOpenFile(c *parser.OpenFileContext) {
	src := loadImage(strings.Trim(c.FileName().GetText(), "\""))
	fmt.Println("Opening file: " + strings.Trim(c.FileName().GetText(), "\"") + " ✓")

	image := Image{g: gift.New(), src: []image.Image{src}}
	l.images[c.ID().GetText()] = image
}

func (l *ImgListener) ExitOpenFolder(c *parser.OpenFolderContext) {
	folderName := strings.Trim(c.FolderName().GetText(), "\"")
	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		log.Fatal(err)
	}
	src := []image.Image{}
	fileNames := []string{}
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
		src = append(src, loadImage(folderName+"/"+f.Name()))
	}

	image := Image{g: gift.New(), src: src, fileNames: fileNames}
	l.images[c.ID().GetText()] = image
}

func (l *ImgListener) ExitExport(c *parser.ExportContext) {
	current := c.ID().GetText()
	if _, ok := l.images[current]; !ok {
		fmt.Println("line " + strconv.Itoa(c.GetStart().GetLine()) + " wrong variable name: " + current)
		os.Exit(3)
	}
	img := l.images[current]
	var dst []*image.RGBA
	for _, src := range img.src {
		d := image.NewRGBA(img.g.Bounds(src.Bounds()))
		img.g.Draw(d, src)
		dst = append(dst, d)
	}
	fmt.Println("Applying filters on: " + current + " ✓")
	if len(img.src) > 1 {
		name := strings.Trim(c.FolderName().GetText(), "\"")
		os.Mkdir(name, os.ModePerm)
		for i, d := range dst {
			saveImage(name+"/"+img.fileNames[i], d)
		}
		fmt.Println("Creating new folder: " + name + " ✓")
	} else {
		name := strings.Trim(c.FileName().GetText(), "\"")
		saveImage(name, dst[0])
		fmt.Println("Creating new file: " + name + " ✓")
	}

}

// Actions ---------------------

func (l *ImgListener) EnterAction_(c *parser.Action_Context) {
	l.currentImage = c.ID().GetText()
	if _, ok := l.images[l.currentImage]; !ok {
		fmt.Println("line " + strconv.Itoa(c.GetStart().GetLine()) + " wrong variable name: " + l.currentImage)
		os.Exit(3)
	}
}

func (l *ImgListener) ExitRotate(c *parser.RotateContext) {
	if c.NUMBER() == nil {
		os.Exit(1)
	}
	params := c.NUMBER().GetText()
	angle, err := strconv.ParseFloat(params, 32)
	if err != nil {
		log.Fatalf("ParseFloat error: %v", err)
	}
	l.images[l.currentImage].g.Add(gift.Rotate(float32(angle), color.Transparent, gift.CubicInterpolation))
}

func (l *ImgListener) ExitFlipY(c *parser.FlipYContext) {
	l.images[l.currentImage].g.Add(gift.FlipVertical())
}

func (l *ImgListener) ExitFlipX(c *parser.FlipXContext) {
	l.images[l.currentImage].g.Add(gift.FlipHorizontal())
}

func (l *ImgListener) ExitCrop(c *parser.CropContext) {
	params := c.AllNUMBER()
	if params == nil || len(params) != 4 {
		os.Exit(1)
	}
	x0, _ := strconv.Atoi(params[0].GetText())
	y0, _ := strconv.Atoi(params[1].GetText())
	x1, _ := strconv.Atoi(params[2].GetText())
	y1, _ := strconv.Atoi(params[3].GetText())

	l.images[l.currentImage].g.Add(gift.Crop(image.Rect(x0, y0, x1, y1)))
}

func (l *ImgListener) ExitResize(c *parser.ResizeContext) {
	params := c.AllNUMBER()
	if params == nil || len(params) != 2 {
		os.Exit(1)
	}
	width, _ := strconv.Atoi(params[0].GetText())
	height, _ := strconv.Atoi(params[1].GetText())
	l.images[l.currentImage].g.Add(gift.Resize(width, height, gift.LanczosResampling))
}

func (l *ImgListener) ExitBrightness(c *parser.BrightnessContext) {
	if c.NUMBER() == nil {
		os.Exit(1)
	}
	number, _ := strconv.ParseFloat(c.NUMBER().GetText(), 32)
	l.images[l.currentImage].g.Add(gift.Brightness(float32(number)))
}

func (l *ImgListener) ExitContrast(c *parser.ContrastContext) {
	if c.NUMBER() == nil {
		os.Exit(1)
	}
	number, _ := strconv.ParseFloat(c.NUMBER().GetText(), 32)
	l.images[l.currentImage].g.Add(gift.Contrast(float32(number)))
}

func (l *ImgListener) ExitSaturation(c *parser.SaturationContext) {
	if c.NUMBER() == nil {
		os.Exit(1)
	}
	number, _ := strconv.ParseFloat(c.NUMBER().GetText(), 32)
	l.images[l.currentImage].g.Add(gift.Saturation(float32(number)))
}

func (l *ImgListener) ExitPixelate(c *parser.PixelateContext) {
	if c.NUMBER() == nil {
		os.Exit(1)
	}
	number, _ := strconv.Atoi(c.NUMBER().GetText())
	l.images[l.currentImage].g.Add(gift.Pixelate(number))
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
