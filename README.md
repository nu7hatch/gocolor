# Go has colors!

This package is a wrapper for standard `fmt` and provides neat colorization
to printing functions.

## Installation

Install with `go get`:

    go get github.com/nu7hatch/gocolor

## Usage

Add to your imports:

    import color "github.com/nu7hatch/gocolor"

And colorize!

    color.Println(color.Red + color.Bright, "I'm bright and red!")
    text := color.Sprintf(color.Blue, "I'm blue %s!", "dabu di dabu da")
    color.Fprint(color.Yellow + color.Underline, os.Stderr, "Yellow warning!")

Enjoy!

## Copyright

Copyright (c) 2012 by Chris Kowalik a.k.a nu7hatch

Released under MIT License. See LICENSE for details.
