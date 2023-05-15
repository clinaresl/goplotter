# Introduction

`goplotter` creates animated gifs with the plots of mathematical functions. This
repository is created only with educational purposes.

It has no additional dependencies

# Usage

`goplotter` acknowledges the following commands in the command-line interface:

- `width` and `height` are the width and height of the GIF image

- `nbsteps` is the number of images to generate in the GIF image

- `filename` is the name of the GIF filename

In addition, `--help` shows the help banner.

The specification of the function to plot is hardcoded as follows:

``` go
	gif := plotFunc(image.Rectangle{Min: image.Point{0, 0},
		Max: image.Point{width, height}},
		nbsteps,
		func(x int) int { return int(math.Pow(float64(x), 2.001)) - x*x })

```

This is, this package actually provides only two services:

- `plotFunc(bounds image.Rectangle, nbsteps int, y func(x int) int) gif.GIF`:
  return a GIF image with `nbsteps` partial plots of the given function which is
  constrained to the specified rectangle
  
- `func saveImage(image *gif.GIF, filename string)`: writes the contents of the
  given GIF image into the specified filename

# Examples

The following execution:

``` sh
$ ./goplotter --width 400 --height 1000 --nbsteps 100 --filename example.gif
```

generates the following GIF animation:

![Example 0](pics/example.gif)

# Author #

Carlos Linares Lopez <carlos.linares@uc3m.es>  
Computer Science and Engineering Department  
Universidad Carlos III de Madrid
