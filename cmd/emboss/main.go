package main

import (
	"context"
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sfomuseum/go-image-emboss"
)

func main() {

	embosser_uri := flag.String("embosser-uri", "grpc://localhost:1234", "A valid sfomuseum/go-image-emboss.Embosser URI.")
	combined := flag.Bool("combined", false, "Return a single image combining all the subjects that were derived from an image.")

	flag.Parse()

	ctx := context.Background()

	em, err := emboss.NewEmbosser(ctx, *embosser_uri)

	if err != nil {
		log.Fatalf("Failed to create new embosser, %v", err)
	}

	defer em.Close(ctx)

	for _, path := range flag.Args() {

		rsp, err := em.EmbossImage(ctx, path, *combined)

		if err != nil {
			log.Fatalf("Failed to extract image from %s, %v", path, err)
		}

		root := filepath.Dir(path)
		fname := filepath.Base(path)
		ext := filepath.Ext(fname)

		fname = strings.Replace(fname, ext, "", 1)

		for idx, im := range rsp {

			im_fname := fmt.Sprintf("%s-emboss-%03d.png", fname, idx+1)

			if *combined {
				im_fname = fmt.Sprintf("%s-emboss-combined-%03d.png", fname, idx+1)
			}

			im_path := filepath.Join(root, im_fname)

			im_wr, err := os.OpenFile(im_path, os.O_RDWR|os.O_CREATE, 0644)

			if err != nil {
				log.Fatalf("Failed to open file for writing %s, %v", im_path, err)
			}

			err = png.Encode(im_wr, im)

			if err != nil {
				log.Fatalf("Failed to write PNG data for %s, %v", im_path, err)
			}

			err = im_wr.Close()

			if err != nil {
				log.Fatalf("Failed to close file for %s, %v", im_path, err)
			}

			log.Println(im_path)
		}

	}
}
