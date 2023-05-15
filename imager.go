// -*- coding: utf-8 -*-
// imager.go
// -----------------------------------------------------------------------------
//
// Started on <mar 09-05-2023 20:34:51.234689776 (1683657291)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Description
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
)

func plotFunc(bounds image.Rectangle, nbsteps int, y func(x int) int) gif.GIF {

	// first, create the necessary arrays for describing the GIF image
	ipaletted := make([]*image.Paletted, nbsteps)
	delay := make([]int, nbsteps)
	disposal := make([]byte, nbsteps)

	// create the palette which consists of 256 levels of yellow
	palette := make([]color.Color, 256)
	for i := 0; i < 256; i++ {
		palette[i] = color.RGBA{uint8(i), uint8(i), 0, 255}
	}

	// create one image at every step
	for istep := 0; istep < nbsteps; istep++ {

		// initialize the contents of this paletted image
		ipaletted[istep] = image.NewPaletted(bounds, palette)

		// and now enable the pixels which show the plot of the function at the
		// i-th step
		for x := bounds.Min.X; x < bounds.Max.X*istep/nbsteps; x++ {

			// apply the function and accept the result only if it falls within
			// bounds
			yvalue := y(x)
			if yvalue >= bounds.Min.Y && yvalue < bounds.Max.Y {
				ipaletted[istep].Pix[x+ipaletted[istep].Stride*(bounds.Max.Y-yvalue-1)] = 0xff
			}
		}

		// set the delay and the disposal method
		delay[istep] = 10
		disposal[istep] = 0
	}

	// create the GIF image and set its parameters
	return gif.GIF{Image: ipaletted,
		Delay:           delay,
		LoopCount:       0,
		Disposal:        disposal,
		Config:          image.Config{},
		BackgroundIndex: 0xff}
}

// save the given gif image into a file
func saveImage(image *gif.GIF, filename string) {

	// create a new file with the given name
	stream, err := os.Create(string(filename))
	if err != nil {
		fmt.Println(err)
		os.Exit(EXIT_FAILURE)
	}

	// make sure to close the file upon exit
	defer stream.Close()

	// create a new GIF image
	gif.EncodeAll(stream, image)
}

// Local Variables:
// mode:go
// fill-column:80
// End:
