// -*- coding: utf-8 -*-
// goplotter.go
// -----------------------------------------------------------------------------
//
// Started on <mar 09-05-2023 19:28:24.654583739 (1683653304)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Description
package main

import (
	"flag"
	"fmt"
	"image"
	"math"
	"os"
)

// global variables
// ----------------------------------------------------------------------------
const VERSION string = "0.1.0" // current version
const EXIT_SUCCESS int = 0     // exit with success
const EXIT_FAILURE int = 1     // exit with failure

// command-line interface
var width, height int // width and height of the image
var nbsteps int       // number of steps in the animation
var filename string   // name of the file where the GIF is saved

var verbose bool // has verbose output been requested?
var version bool // has version info been requested?

// functions
// ----------------------------------------------------------------------------

// initialize the command line parser
func init() {

	// mandatory arguments follow
	flag.IntVar(&width, "width", 1000, "width of the image. 1000 by default")
	flag.IntVar(&height, "height", 1000, "height of the image. 1000 by default")
	flag.IntVar(&nbsteps, "nbsteps", 100, "number of steps in the animation. 100 by default")
	flag.StringVar(&filename, "filename", "image.gif", "name of the file where the GIF is saved. 'image.gif' by default")

	// other optional parameters are verbose and version
	flag.BoolVar(&verbose, "verbose", false, "provides verbose output")
	flag.BoolVar(&version, "version", false, "shows version info and exists")
}

// show current version
func showVersion() {

	// show version
	fmt.Printf("goplotter version %s\n", VERSION)
	os.Exit(EXIT_SUCCESS)
}

// verify the user commands given in the command line interface
func verify() {

	// in case the user requested version info, show and exit
	if version {
		showVersion()
	}

	// Check the width and height are legal values
	if width <= 0 || height <= 0 {
		fmt.Println("Error: width and height must be positive integers")
		os.Exit(EXIT_FAILURE)
	}

	// and also that the number of steps is a positive number
	if nbsteps <= 0 {
		fmt.Println("Error: number of steps must be a positive integer")
		os.Exit(EXIT_FAILURE)
	}
}

// main function
func main() {

	// parse the command line
	flag.Parse()

	// verify the arguments and act accordinly in case it is necessary
	verify()

	// compute all the different plots at every step
	gif := plotFunc(image.Rectangle{Min: image.Point{0, 0},
		Max: image.Point{width, height}},
		nbsteps,
		func(x int) int { return int(math.Pow(float64(x), 2.001)) - x*x })

	// save the image in a gif file
	saveImage(&gif, filename)
}

// Local Variables:
// mode:go
// fill-column:80
// End:
