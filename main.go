package main

import (
	"bytes"
	_ "embed"
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

//go:embed Montserrat-ExtraBold.ttf
var titleFontFile []byte

//go:embed Montserrat-SemiBold.ttf
var contentFontFile []byte

//go:embed blue.png
var background []byte

var (
	width       = 1200
	height      = 630
	title       = flag.String("title", "Sample Post Title", "Title text")
	description = flag.String("desc", "Description of the post in about 25 words. This string could be your og description also!", "Description text")
	date        = flag.String("date", "30 April 2021", "Date text")
	readingtime = flag.String("readtime", "4 min read", "Reading time text")
	out         = flag.String("out", "out.png", "Name of the output file")
	dpi         = flag.Float64("dpi", 72, "screen resolution in dots per inch")
	tfont       = flag.String("tfont", "Montserrat-ExtraBold.ttf", "Font to use for the title")
	cfont       = flag.String("cfont", "Montserrat-SemiBold.ttf", "Font to use for the content (description and date)")
)

var (
	marginTop                = 130
	marginLeft               = 114
	titleSize        float64 = 42
	descriptionSize          = 22
	titleColor               = color.NRGBA{96, 234, 206, 255}
	descriptionColor         = color.NRGBA{232, 232, 232, 255}
)

func main() {
	flag.Parse()

	// create logger
	logger := log.New(os.Stdout, "og-image-generator: ", log.LstdFlags)

	tFont, err := truetype.Parse(titleFontFile)
	if err != nil {
		logger.Printf("Error parsing font: %s\n", *tfont)
		logger.Println(err)
		return
	}

	cFont, err := truetype.Parse(contentFontFile)
	if err != nil {
		logger.Printf("Error parsing font: %s", *cfont)
		logger.Println(err)
		return
	}

	bg, _, _ := image.Decode(bytes.NewReader(background))

	output := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.Draw(output, bg.Bounds(), bg, image.Point{0, 0}, draw.Over)

	outFile, err := os.Create(*out)
	defer outFile.Close()

	if err != nil {
		logger.Printf("Error creating output file : %s\n", *out)
		logger.Println(err)
	}

	tDrawer := &font.Drawer{
		Dst:  output,
		Face: truetype.NewFace(tFont, &truetype.Options{Size: titleSize, Hinting: font.HintingNone, DPI: *dpi}),
		Src:  image.NewUniform(titleColor),
	}

	tDrawer.Dot = fixed.Point26_6{
		X: fixed.I(marginLeft),
		Y: fixed.I(marginTop),
	}

	headline := wrapLines(*title, true)

	for _, r := range headline {
		if r != "" {
			tDrawer.DrawString(r)

			yCursor := int(titleSize) + 4 + marginTop

			// move the point to next line
			tDrawer.Dot = fixed.Point26_6{
				X: fixed.I(marginLeft),
				Y: fixed.I(yCursor), // 4 is space between lines
			}

			marginTop = yCursor // update the global cursor
		}
	}

	cDrawer := &font.Drawer{
		Dst:  output,
		Face: truetype.NewFace(cFont, &truetype.Options{Size: 20, Hinting: font.HintingNone, DPI: *dpi}),
		Src:  image.NewUniform(descriptionColor),
	}

	cDrawer.Dot = fixed.Point26_6{
		X: fixed.I(marginLeft),
		Y: fixed.I(marginTop), // 12 (8 + 4), 4 is from the last iteration of loop
	}

	descriptionText := wrapLines(*description, false)

	for _, r := range descriptionText {
		if r != "" {
			cDrawer.DrawString(r)

			yCursor := marginTop + int(descriptionSize) + 4

			// move the point to next line
			cDrawer.Dot = fixed.Point26_6{
				X: fixed.I(marginLeft),
				Y: fixed.I(yCursor), // 4 is space between lines
			}

			marginTop = yCursor // update the global cursor
		}
	}

	marginTop = 100

	cDrawer.Dot = fixed.Point26_6{
		X: fixed.I(marginLeft),
		Y: fixed.I(bg.Bounds().Size().Y - marginTop), // 12 (8 + 4), 4 is from the last iteration of loop
	}

	cDrawer.DrawString(*date)

	if *readingtime != "" {
		cDrawer.DrawString(" | ")
		cDrawer.DrawString(*readingtime)
	}

	png.Encode(outFile, output)

	logger.Printf("File created successfully: %s!", *out)
}

// wrapLines returns slice of string of max length 4
// There can be empty string is the length of arg text is
// small
func wrapLines(text string, isTitle bool) [4]string {
	limit := 14

	if isTitle {
		limit = 5
	}

	var out [4]string

	words := strings.Split(text, " ")

	for i, r := range words {
		out[i/limit] += r
		out[i/limit] += " "
	}

	return out
}
