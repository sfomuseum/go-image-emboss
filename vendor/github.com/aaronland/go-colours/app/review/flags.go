package review

import (
	"flag"
	"fmt"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var extruder_uris multi.MultiString
var palette_uris multi.MultiString

var allow_remote bool
var root string
var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("review")

	fs.Var(&extruder_uris, "extruder-uri", "Zero or more aaronland/go-colours/extruder.Extruder URIs. Default is to use all registered extruder schemes.")
	fs.Var(&palette_uris, "palette-uri", "Zero or more aaronland/go-colours/palette.Palette URIs. Default is to use all registered palette schemes.")
	fs.StringVar(&root, "root", "", "The path to a directory where images and HTML files associated with the review should be stored. If empty a new temporary directory will be created (and deleted when the application exits).")
	fs.BoolVar(&allow_remote, "allow-remote", true, "Allow fetching remote images (HTTP(S)).")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Command line tool to perform colour extraction and \"snap-to-grid\" matching with one or more colour palettes for images, emitting the results as an HTML page.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options] uri(N) uri(N)\n", os.Args[0])
		fs.PrintDefaults()
	}

	return fs
}
